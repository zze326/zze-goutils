package zze_goutils

//
//  MergeMaps
//  @Description: 将一个 map[interface{}]interface{} 补丁合并到另一个 map[interface{}]interface{}
//  @param dest 合并到的目标 map
//  @param src 合并源 map
//  @return map[interface{}]interface{} 合并后的 map
//
func MergeMaps(dest, src map[interface{}]interface{}) map[interface{}]interface{} {
	out := make(map[interface{}]interface{}, len(dest))
	for k, v := range dest {
		out[k] = v
	}
	for k, v := range src {
		value := v
		if av, ok := out[k]; ok {
			if v, ok := v.(map[interface{}]interface{}); ok {
				if av, ok := av.(map[interface{}]interface{}); ok {
					out[k] = MergeMaps(av, v)
				} else {
					out[k] = v
				}
			} else {
				out[k] = value
			}
		} else {
			out[k] = v
		}
	}
	return out
}

