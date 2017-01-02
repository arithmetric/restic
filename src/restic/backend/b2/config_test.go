package b2

import "testing"

var configTests = []struct {
	s   string
	cfg Config
}{
	{"b2:bucketname", Config{
		Bucket: "bucketname",
		Prefix: "restic",
	}},
	{"b2:bucketname/", Config{
		Bucket: "bucketname",
		Prefix: "restic",
	}},
	{"b2:bucketname/prefix/directory", Config{
		Bucket: "bucketname",
		Prefix: "prefix/directory",
	}},
	{"b2:bucketname/prefix/directory/", Config{
		Bucket: "bucketname",
		Prefix: "prefix/directory",
	}},
	{"b2:foobar", Config{
		Bucket: "foobar",
		Prefix: "restic",
	}},
	{"b2:foobar/", Config{
		Bucket: "foobar",
		Prefix: "restic",
	}},
}

func TestParseConfig(t *testing.T) {
	for i, test := range configTests {
		cfg, err := ParseConfig(test.s)
		if err != nil {
			t.Errorf("test %d:%s failed: %v", i, test.s, err)
			continue
		}

		if cfg != test.cfg {
			t.Errorf("test %d:\ninput:\n  %s\n wrong config, want:\n  %v\ngot:\n  %v",
				i, test.s, test.cfg, cfg)
			continue
		}
	}
}

var invalidConfigTests = []struct {
	s   string
	cfg interface{}
}{
	{"a2:foobar/", nil},
	{"b2:", nil},
	{"b2:/", nil},
}

func TestParseInvalidConfig(t *testing.T) {
	for i, test := range invalidConfigTests {
		cfg, err := ParseConfig(test.s)
		if err == nil || cfg != test.cfg {
			t.Errorf("test %d:%s should fail, but returned: %v", i, test.s, cfg)
			continue
		}
	}
}
