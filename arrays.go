package utils

func ConvertArray[srcType any, dstType any](srcArr []srcType, converter func(srcType) (dstType, error)) ([]dstType, error) {
	var err error
	dstArr := make([]dstType, len(srcArr))
	for i := range srcArr {
		dstArr[i], err = converter(srcArr[i])
		if err != nil {
			return nil, err
		}
	}
	return dstArr, nil
}
