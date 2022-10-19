package libs

type Params map[string]interface{}
type Result map[string]interface{}

func MergeParams(dst Params, src ...Params) Params {
	for _, params := range src {
		for key, value := range params {
			dst[key] = value
		}
	}

	return dst
}
