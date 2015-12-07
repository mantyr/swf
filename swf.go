package swf

import (
    "io"
    "os"
    "errors"
    "compress/zlib"
)

type SWF struct {
    magic string
    compressed bool
    version int
    size int
    width int
    height int
    Error error
}

func Open(address string) (s *SWF, err error) {
    s = new(SWF)
    file, err := os.Open(address)
    if err != nil {
        s.Error = err
        return
    }
    defer file.Close()
    s, err = Decode(file)
    if err != nil {
        s.Error = err
        return
    }
    return
}

func (s *SWF) Width() int {
    return s.width
}

func (s *SWF) Height() int {
    return s.height
}

func Decode(r io.Reader) (s *SWF, err error) {
    s = new(SWF)
    buffer := make([]byte, 8)

    buffer_read, err := r.Read(buffer)
    if err != nil {
        return
    }
    if buffer_read < 8 {
        err = errors.New("Bad header SWF file")
        return
    }
    s.magic = string(buffer[0:3])
    if s.magic != "FWS" && s.magic != "CWS" {
        err = errors.New("No SWF file")
        return
    }
    if s.magic[0:1] == "C" {
        s.compressed = true
    }
    s.version = int(buffer[3])
    for i := 0; i < 4; i++ {
        t := int(buffer[i+4])
        s.size += t<<uint(8*i)
    }
    buffer_body := make([]byte, s.size)
    if s.compressed {
        var buffer_io io.ReadCloser
        buffer_io, err = zlib.NewReader(r)
        if err != nil {
            return s, err
        }
        buffer_read, err = buffer_io.Read(buffer_body)
        buffer_io.Close()
    } else {
        buffer_read, err = r.Read(buffer_body)
    }

    cbyte := int(buffer_body[0])
    bits  := cbyte>>3
    cbyte &= 7
    cbyte <<= 5
    cval := ""
    cbit := 2
    var bitcount int
    var buffer_i int = 1
    for i := 0; i < 4; i++ {
        cval = ""
        bitcount = 0
        for bitcount < bits {
            if (cbyte & 128) > 0 {
                cval += "1"
            } else {
                cval += "0"
            }
            cbyte <<=1
            cbyte &= 255
            cbit--
            bitcount++
            if cbit < 0 {
                cbyte = int(buffer_body[buffer_i])
                buffer_i++
                cbit = 7
            }
        }
        c := 1
        val := 0
        for n := len(cval); n > 0; n-- {
            if string(cval[n-1]) == "1" {
                val +=c
            }
            c *=2
        }
        val /= 20
        switch i {
            case 0:
                s.width = val
            case 1:
                s.width = val - s.width
            case 2:
                s.height = val
            case 3:
                s.height = val - s.height
        }
    }
    return
}