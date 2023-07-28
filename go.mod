module github.com/jmwample/nested_mods

go 1.20

replace github.com/jmwample/nested_mods/mods/client => ./mods/client

require github.com/jmwample/nested_mods/mods/client v0.0.0-00010101000000-000000000000
