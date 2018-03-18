# Example 3

Simple example of how to start a Branching Pipeline.
It Currently does the same thing as example2's non-branching pipeline
but you can optionally add more DataProcessors to each stage. 

See example4 for a more serious usage.

## USAGE

	mysql> source schema.sql   (make sure you read through schema.sql first)
	$ go build
	$ ./example3

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
