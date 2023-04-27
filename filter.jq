# Convert int64 API fields to use integer rather than string. The OpenAPI uses
# strings to represent int64 values in order to better support Javascript. But
# this is awkward in Go, so generate these fields with int64 types instead. Use
# mustache to add a ",string" modifier to the json attribute on these fields so
# that the JSON unmarshaler will convert the API string to an int64 value.
(.. | objects) |=
  if has("type") and .type == "string" and has("format") and .format == "int64" then
    .type = "integer"
  else
    .
  end
