# Example 2
This builds on example1 but adds a custom DataProcessor to transform the data.

## USAGE

	mysql> source schema.sql   (make sure you read through schema.sql first)
	$ go build
	$ ./example2

This should have populated dstDB's `users2` table using transformed data:

	source (srcDB `users`):

		id    |  name
		------|-------
		123   |  Alex
		456   |  John

	destination (dstDB `users2`):

		user_id  |  some_new_field
		---------|----------------
		123      |  whatever
		456      |  whatever

Our transformer changed the name `id` to `user_id` for the destination table.
It also omitted the `name` field and used a new field `some_new_field` that 
wasn't present in the source table.
