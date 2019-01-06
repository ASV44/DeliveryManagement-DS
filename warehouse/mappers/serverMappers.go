package mappers

func UrlQueryToMap(urlQueries map[string][]string, queryMapper func(query string) string) map[string]string {
	queryMap := make(map[string]string)
	for query, value := range urlQueries {
		mappedQuery := query
		if queryMapper != nil {
			mappedQuery = queryMapper(query)
		}
		queryMap[mappedQuery] = ""
		if len(value) != 0 {
			queryMap[mappedQuery] = value[0]
		}
	}

	return queryMap
}
