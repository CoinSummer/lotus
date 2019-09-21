package rlepluslazy

var goldenRLE = []byte{0x40, 0x31, 0x40, 0x84, 0x5c, 0xc0, 0xc6, 0x38, 0xc0, 0xc1, 0x5e, 0x20, 0x92, 0x6, 0x12, 0xbd, 0x9, 0x3a, 0xa3, 0x1, 0x7c, 0x60, 0x97, 0xec, 0x3, 0x8f, 0x4f, 0x1, 0x95, 0x19, 0x26, 0x30, 0x66, 0x80, 0x7, 0x4b, 0x1, 0x7, 0x5f, 0x8, 0xc7, 0x94, 0x28, 0x1f, 0xd0, 0xa9, 0x47, 0x88, 0x7f, 0x80, 0x82, 0x71, 0x82, 0x8, 0x64, 0x82, 0x83, 0x91, 0x46, 0x93, 0x7, 0xb5, 0x40, 0x5f, 0xe4, 0x3b, 0x20, 0x1b, 0x1a, 0x71, 0x44, 0x70, 0x6e, 0x21, 0x81, 0x2e, 0x40, 0x6, 0xfc, 0x7, 0x16, 0xa9, 0x2, 0x16, 0xc8, 0x2, 0x12, 0x34, 0x48, 0xe0, 0x3e, 0xe0, 0x43, 0x2d, 0xa0, 0x52, 0xff, 0xe0, 0xf1, 0xc, 0x1, 0x4, 0xb5, 0x2, 0x44, 0x98, 0x12, 0x7a, 0x26, 0x36, 0xa, 0xa5, 0x2, 0xa9, 0x24, 0x80, 0x2, 0x9, 0x8b, 0x57, 0x1, 0x5, 0xd8, 0x4, 0x11, 0x53, 0x2, 0x95, 0xdc, 0x3, 0x95, 0xc5, 0x0, 0x91, 0xcb, 0x1, 0x36, 0xb8, 0x71, 0xd8, 0xe, 0x50, 0x28, 0xf4, 0x70, 0x18, 0xe, 0xd0, 0x89, 0x64, 0xb0, 0xc0, 0xf, 0x70, 0x18, 0xb3, 0x18, 0xd, 0x48, 0x5, 0x10, 0x84, 0x6, 0x20, 0x41, 0xff, 0x81, 0x8, 0xed, 0x81, 0xc3, 0x6a, 0x82, 0xf, 0xee, 0x80, 0xc2, 0xe4, 0x0, 0x20, 0x34, 0x3f, 0x98, 0x4c, 0x78, 0xd0, 0x1b, 0xf0, 0xc1, 0x34, 0x83, 0x8, 0x78, 0x81, 0x88, 0xb4, 0x80, 0x42, 0xbb, 0x1, 0x10, 0x30, 0x89, 0xbc, 0xf, 0x24, 0x33, 0xc1, 0x82, 0xf4, 0x40, 0xa2, 0x74, 0xc0, 0x4, 0x5f, 0xc1, 0x21, 0x52, 0x40, 0xa1, 0x73, 0xc0, 0x26, 0xfb, 0xc0, 0xe4, 0x4, 0x9, 0xda, 0x9, 0x3c, 0xca, 0x9, 0xa0, 0x20, 0x7a, 0xc0, 0x21, 0x51, 0x0, 0x14, 0xf2, 0xf, 0x34, 0x44, 0x9, 0x58, 0xa1, 0xbe, 0x80, 0xe, 0xe2, 0xd, 0x54, 0x60, 0x7, 0x34, 0x20, 0x44, 0xd2, 0x19, 0x14, 0xbe, 0x1c, 0xfa, 0xd, 0x6c, 0x54, 0x25, 0x88, 0xc0, 0xd5, 0x40, 0x7}