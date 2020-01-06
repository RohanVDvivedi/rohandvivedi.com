package data

func Find(parameterizedQuery string, parameters interface{}) interface{} {
	// convert parameterized query to nativeQuery
	// serialize parameters
	// search the results in memcache against given parameters
	// if the solution is found return
	// else go search in mysql
	// if found cache the result, and return result
	// else return nil
}