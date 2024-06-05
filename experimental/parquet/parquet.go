package parquet

import (
	"fmt"
	"os"
	"path"

	"github.com/apache/arrow/go/v15/parquet"
	"github.com/apache/arrow/go/v15/parquet/pqarrow"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

var logger = log.DefaultLogger

func ToParquet(frames []*data.Frame, chunk int) (map[string]string, error) {
	dirs := map[string]string{}
	frameIndex := framesByRef(frames)

	// TODO - appending lables to fields for now
	// need to return multiple frames instead
	// for _, f := range frames {
	// 	for _, fld := range f.Fields {
	// 		if len(fld.Labels) > 0 {
	// 			lbls := fld.Labels.String()
	// 			fld.Name = fmt.Sprintf("%s %s", fld.Name, lbls)
	// 		}
	// 	}
	// }

	for _, frameList := range frameIndex {
		labelsToFields(frameList)

		dir, err := os.MkdirTemp("", "duck")
		if err != nil {
			logger.Error("failed to create temp dir", "error", err)
			return nil, err
		}

		SIZELEN := int64(1024 * 1024)
		props := parquet.NewWriterProperties(parquet.WithDictionaryDefault(false))

		mergeFrames(frameList)
		for i, frame := range frameList {
			dirs[frame.RefID] = dir

			tbl, err := data.FrameToArrowTable(frame)
			if err != nil {
				logger.Error("failed to marshal arrow schema", "error", err)
				return nil, err
			}

			name := fmt.Sprintf("%s%d", frame.RefID, i)
			filename := path.Join(dir, name+".parquet")
			output, err := os.Create(filename)
			if err != nil {
				logger.Error("failed to create parquet file", "file", filename, "error", err)
				return nil, err
			}
			defer output.Close()

			err = pqarrow.WriteTable(tbl, output, SIZELEN, props, pqarrow.DefaultWriterProps())
			if err != nil {
				return nil, err
			}
		}
	}
	return dirs, nil
}

func framesByRef(frames []*data.Frame) map[string][]*data.Frame {
	byRef := map[string][]*data.Frame{}
	for _, f := range frames {
		fr := byRef[f.RefID]
		if fr == nil {
			refFrames := []*data.Frame{}
			byRef[f.RefID] = refFrames
		}
		byRef[f.RefID] = append(byRef[f.RefID], clone(f))
	}
	return byRef
}

func clone(f *data.Frame) *data.Frame {
	copy := data.NewFrame(f.Name, f.Fields...)
	copy.RefID = f.RefID
	copy.Meta = f.Meta
	return copy
}

func mergeFrames(frames []*data.Frame) {
	fields := map[string]*data.Field{}
	for _, f := range frames {
		for _, fld := range f.Fields {
			fields[fld.Name] = fld
		}
	}
	for _, fld := range fields {
		for _, f := range frames {
			found := false
			for _, fld2 := range f.Fields {
				if fld2.Name == fld.Name {
					found = true
					break
				}
			}
			if !found {
				makeArray := maker[fld.Type()]
				arr := makeArray(f.Rows())
				nullField := data.NewField(fld.Name, fld.Labels, arr)
				f.Fields = append(f.Fields, nullField)
			}
		}
	}
}

func labelsToFields(frames []*data.Frame) {
	for _, f := range frames {
		fields := []*data.Field{}
		for _, fld := range f.Fields {
			if fld.Labels != nil {
				for lbl, val := range fld.Labels {
					newFld := newField(lbl, val, f.Rows())
					fields = append(fields, newFld)
				}
			}
		}
		f.Fields = append(f.Fields, fields...)
	}
}

func newField(name string, val string, size int) *data.Field {
	values := make([]string, size)
	newField := data.NewField(name, nil, values)
	for i := 0; i < size; i++ {
		newField.Set(i, val)
	}
	return newField
}
