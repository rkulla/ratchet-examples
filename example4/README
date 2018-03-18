# Example 4

More involved example of a Branching Pipeline.

## USAGE

	mysql> source schema.sql   (make sure you read through schema.sql first)
	$ go build
	$ ./example4

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

		user_id  |  name      | city
		---------|------------------
		123      |  Alex      | Austin
		456      |  John      | Los Angeles
		789      |  Jane      | San Diego
