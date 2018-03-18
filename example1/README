# Example 1

This example is the most basic Ratchet program.

It consists of just an SQLReader and an SQLWriter, no transformer.

It demonstrates that a transformer can be omitted from a pipeline as long as
the destination table matches the source table.  While this can be useful for
moving data around, typically you will use transformers (see example2).

## USAGE

Run the following commands, in order:

	mysql> source schema.sql   (make sure you read through schema.sql first)
	$ go build
	$ ./example1

This should have populated dstDB's `users2` table and should match the srcDB's
`users` table exactly:

    source (srcDB `users`):

		id    |  name
		------|-------
		123   |  Alex
		456   |  John

	destination (dstDB `users2`):

		id    |  name
		------|-------
		123   |  Alex
		456   |  John
