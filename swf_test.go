package swf

import (
    "testing"
)

func TestWidthHeight(t *testing.T) {
    test_group_1("Width/Height", "./testdata/f0673319418b054b4f108c18736a042f5dcec4e2123401c48a727dfaab7354ef.swf", 640, 480, t)
//    test_group_1("Width/Height", "./testdata/f3c6de85c2ca1260ad2117a0c536a82458c6d01a38f0edbc1196097af1f609a4.swf", 6409, 480, t) // No SWF file
    test_group_1("Width/Height", "./testdata/1955837c326a7b842763d65e68d62968b6b42898d1644b5ea90b3a656dc65b6e.swf", 650, 490, t)
    test_group_1("Width/Height", "./testdata/1a6d3ae1a4a7756d38138d24e3c41d3b6c7085e115b0d91ad78d3c9b7e336185.swf", 650, 550, t)
    test_group_1("Width/Height", "./testdata/1ddcf6de92c5534805a634a3bb3649f4e8fc346cc7c81f1f55c74737978ad27a.swf", 660, 700, t)

    filename := "./testdata/f3c6de85c2ca1260ad2117a0c536a82458c6d01a38f0edbc1196097af1f609a4.swf"
    name := "No SWF file"
    s, err := Open(filename)
    if err == nil || s.Error == nil || err.Error() != "No SWF file" || s.Error.Error() != "No SWF file" {
        t.Errorf("Error %q %q %q %q", name, filename, err.Error(), s.Error.Error())
    }
}

func test_group_1(name string, filename string, width int, height int, t *testing.T) {
    s, err := Open(filename)
    if err != nil {
        t.Errorf("Error %q %q %q", name, filename, err.Error())
    }
    if s.Error != nil {
        t.Errorf("Error %q %q %q", name, filename, s.Error.Error())
    }
    if s.Width() != width || s.Height() != height {
        t.Errorf("Error %q %q size width/height,  %dx%d => error size %dx%d", name, filename, width, height, s.Width(), s.Height())
    }
}