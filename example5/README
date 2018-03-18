# Example 5

Example of a DynamicSQLReader

## USAGE

	mysql> source schema.sql   (make sure you read through schema.sql first)
	$ go build
	$ ./example5

This should do:

	source1 (srcDB `users`):

		id    |  name
		------|-------
		123   |  Alex
		456   |  John
		789   |  Jane

	source2 (srcDB `addresses`):

		id    |  name
		------|-------
		123   |  Austin
		456   |  Los Angeles
		789   |  San Diego

	destination (dstDB `users2`):

		user_id  |  city
		---------|----------------
		456      |  Los Angeles
		789      |  San Diego
