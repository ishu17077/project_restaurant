// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package avx2

var _text_b64encode = []byte{
	// .p2align 5, 0x00
	// LCPI0_0
	0x47, // .byte 71
	0xfc, //0x00000001 .byte 252
	0xfc, //0x00000002 .byte 252
	0xfc, //0x00000003 .byte 252
	0xfc, //0x00000004 .byte 252
	0xfc, //0x00000005 .byte 252
	0xfc, //0x00000006 .byte 252
	0xfc, //0x00000007 .byte 252
	0xfc, //0x00000008 .byte 252
	0xfc, //0x00000009 .byte 252
	0xfc, //0x0000000a .byte 252
	0xed, //0x0000000b .byte 237
	0xf0, //0x0000000c .byte 240
	0x41, //0x0000000d .byte 65
	0x00, //0x0000000e .byte 0
	0x00, //0x0000000f .byte 0
	0x47, //0x00000010 .byte 71
	0xfc, //0x00000011 .byte 252
	0xfc, //0x00000012 .byte 252
	0xfc, //0x00000013 .byte 252
	0xfc, //0x00000014 .byte 252
	0xfc, //0x00000015 .byte 252
	0xfc, //0x00000016 .byte 252
	0xfc, //0x00000017 .byte 252
	0xfc, //0x00000018 .byte 252
	0xfc, //0x00000019 .byte 252
	0xfc, //0x0000001a .byte 252
	0xed, //0x0000001b .byte 237
	0xf0, //0x0000001c .byte 240
	0x41, //0x0000001d .byte 65
	0x00, //0x0000001e .byte 0
	0x00, //0x0000001f .byte 0
	//0x00000020 LCPI0_1
	0x47, //0x00000020 .byte 71
	0xfc, //0x00000021 .byte 252
	0xfc, //0x00000022 .byte 252
	0xfc, //0x00000023 .byte 252
	0xfc, //0x00000024 .byte 252
	0xfc, //0x00000025 .byte 252
	0xfc, //0x00000026 .byte 252
	0xfc, //0x00000027 .byte 252
	0xfc, //0x00000028 .byte 252
	0xfc, //0x00000029 .byte 252
	0xfc, //0x0000002a .byte 252
	0xef, //0x0000002b .byte 239
	0x20, //0x0000002c .byte 32
	0x41, //0x0000002d .byte 65
	0x00, //0x0000002e .byte 0
	0x00, //0x0000002f .byte 0
	0x47, //0x00000030 .byte 71
	0xfc, //0x00000031 .byte 252
	0xfc, //0x00000032 .byte 252
	0xfc, //0x00000033 .byte 252
	0xfc, //0x00000034 .byte 252
	0xfc, //0x00000035 .byte 252
	0xfc, //0x00000036 .byte 252
	0xfc, //0x00000037 .byte 252
	0xfc, //0x00000038 .byte 252
	0xfc, //0x00000039 .byte 252
	0xfc, //0x0000003a .byte 252
	0xef, //0x0000003b .byte 239
	0x20, //0x0000003c .byte 32
	0x41, //0x0000003d .byte 65
	0x00, //0x0000003e .byte 0
	0x00, //0x0000003f .byte 0
	//0x00000040 LCPI0_2
	0x01, //0x00000040 .byte 1
	0x00, //0x00000041 .byte 0
	0x02, //0x00000042 .byte 2
	0x01, //0x00000043 .byte 1
	0x04, //0x00000044 .byte 4
	0x03, //0x00000045 .byte 3
	0x05, //0x00000046 .byte 5
	0x04, //0x00000047 .byte 4
	0x07, //0x00000048 .byte 7
	0x06, //0x00000049 .byte 6
	0x08, //0x0000004a .byte 8
	0x07, //0x0000004b .byte 7
	0x0a, //0x0000004c .byte 10
	0x09, //0x0000004d .byte 9
	0x0b, //0x0000004e .byte 11
	0x0a, //0x0000004f .byte 10
	0x01, //0x00000050 .byte 1
	0x00, //0x00000051 .byte 0
	0x02, //0x00000052 .byte 2
	0x01, //0x00000053 .byte 1
	0x04, //0x00000054 .byte 4
	0x03, //0x00000055 .byte 3
	0x05, //0x00000056 .byte 5
	0x04, //0x00000057 .byte 4
	0x07, //0x00000058 .byte 7
	0x06, //0x00000059 .byte 6
	0x08, //0x0000005a .byte 8
	0x07, //0x0000005b .byte 7
	0x0a, //0x0000005c .byte 10
	0x09, //0x0000005d .byte 9
	0x0b, //0x0000005e .byte 11
	0x0a, //0x0000005f .byte 10
	//0x00000060 LCPI0_3
	0x00, 0xfc, //0x00000060 .word 64512
	0xc0, 0x0f, //0x00000062 .word 4032
	0x00, 0xfc, //0x00000064 .word 64512
	0xc0, 0x0f, //0x00000066 .word 4032
	0x00, 0xfc, //0x00000068 .word 64512
	0xc0, 0x0f, //0x0000006a .word 4032
	0x00, 0xfc, //0x0000006c .word 64512
	0xc0, 0x0f, //0x0000006e .word 4032
	0x00, 0xfc, //0x00000070 .word 64512
	0xc0, 0x0f, //0x00000072 .word 4032
	0x00, 0xfc, //0x00000074 .word 64512
	0xc0, 0x0f, //0x00000076 .word 4032
	0x00, 0xfc, //0x00000078 .word 64512
	0xc0, 0x0f, //0x0000007a .word 4032
	0x00, 0xfc, //0x0000007c .word 64512
	0xc0, 0x0f, //0x0000007e .word 4032
	//0x00000080 LCPI0_4
	0x40, 0x00, //0x00000080 .word 64
	0x00, 0x04, //0x00000082 .word 1024
	0x40, 0x00, //0x00000084 .word 64
	0x00, 0x04, //0x00000086 .word 1024
	0x40, 0x00, //0x00000088 .word 64
	0x00, 0x04, //0x0000008a .word 1024
	0x40, 0x00, //0x0000008c .word 64
	0x00, 0x04, //0x0000008e .word 1024
	0x40, 0x00, //0x00000090 .word 64
	0x00, 0x04, //0x00000092 .word 1024
	0x40, 0x00, //0x00000094 .word 64
	0x00, 0x04, //0x00000096 .word 1024
	0x40, 0x00, //0x00000098 .word 64
	0x00, 0x04, //0x0000009a .word 1024
	0x40, 0x00, //0x0000009c .word 64
	0x00, 0x04, //0x0000009e .word 1024
	//0x000000a0 LCPI0_5
	0xf0, 0x03, //0x000000a0 .word 1008
	0x3f, 0x00, //0x000000a2 .word 63
	0xf0, 0x03, //0x000000a4 .word 1008
	0x3f, 0x00, //0x000000a6 .word 63
	0xf0, 0x03, //0x000000a8 .word 1008
	0x3f, 0x00, //0x000000aa .word 63
	0xf0, 0x03, //0x000000ac .word 1008
	0x3f, 0x00, //0x000000ae .word 63
	0xf0, 0x03, //0x000000b0 .word 1008
	0x3f, 0x00, //0x000000b2 .word 63
	0xf0, 0x03, //0x000000b4 .word 1008
	0x3f, 0x00, //0x000000b6 .word 63
	0xf0, 0x03, //0x000000b8 .word 1008
	0x3f, 0x00, //0x000000ba .word 63
	0xf0, 0x03, //0x000000bc .word 1008
	0x3f, 0x00, //0x000000be .word 63
	//0x000000c0 LCPI0_6
	0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, //0x000000c0 QUAD $0x1a1a1a1a1a1a1a1a; QUAD $0x1a1a1a1a1a1a1a1a  // .space 16, '\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a'
	0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, 0x1a, //0x000000d0 QUAD $0x1a1a1a1a1a1a1a1a; QUAD $0x1a1a1a1a1a1a1a1a  // .space 16, '\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a\x1a'
	//0x000000e0 LCPI0_7
	0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, //0x000000e0 QUAD $0x3333333333333333; QUAD $0x3333333333333333  // .space 16, '3333333333333333'
	0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, //0x000000f0 QUAD $0x3333333333333333; QUAD $0x3333333333333333  // .space 16, '3333333333333333'
	//0x00000100 LCPI0_8
	0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, //0x00000100 QUAD $0x0d0d0d0d0d0d0d0d; QUAD $0x0d0d0d0d0d0d0d0d  // .space 16, '\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r'
	0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, 0x0d, //0x00000110 QUAD $0x0d0d0d0d0d0d0d0d; QUAD $0x0d0d0d0d0d0d0d0d  // .space 16, '\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r'
	//0x00000120 .p2align 4, 0x90
	//0x00000120 _b64encode
	0x55, //0x00000120 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000121 movq         %rsp, %rbp
	0x41, 0x57, //0x00000124 pushq        %r15
	0x41, 0x56, //0x00000126 pushq        %r14
	0x41, 0x54, //0x00000128 pushq        %r12
	0x53, //0x0000012a pushq        %rbx
	0x4c, 0x8b, 0x4e, 0x08, //0x0000012b movq         $8(%rsi), %r9
	0x4d, 0x85, 0xc9, //0x0000012f testq        %r9, %r9
	0x0f, 0x84, 0x11, 0x03, 0x00, 0x00, //0x00000132 je           LBB0_26
	0x4c, 0x8b, 0x26, //0x00000138 movq         (%rsi), %r12
	0x4c, 0x8b, 0x07, //0x0000013b movq         (%rdi), %r8
	0x4d, 0x01, 0xe1, //0x0000013e addq         %r12, %r9
	0x49, 0x8d, 0x71, 0xe4, //0x00000141 leaq         $-28(%r9), %rsi
	0x48, 0x8d, 0x0d, 0x14, 0x03, 0x00, 0x00, //0x00000145 leaq         $788(%rip), %rcx  /* _TabEncodeCharsetStd+0(%rip) */
	0x4c, 0x8d, 0x15, 0x4d, 0x03, 0x00, 0x00, //0x0000014c leaq         $845(%rip), %r10  /* _TabEncodeCharsetURL+0(%rip) */
	0xf6, 0xc2, 0x01, //0x00000153 testb        $1, %dl
	0x4c, 0x0f, 0x44, 0xd1, //0x00000156 cmoveq       %rcx, %r10
	0x0f, 0x84, 0x30, 0x02, 0x00, 0x00, //0x0000015a je           LBB0_2
	0xc5, 0xfe, 0x6f, 0x05, 0xb8, 0xfe, 0xff, 0xff, //0x00000160 vmovdqu      $-328(%rip), %ymm0  /* LCPI0_1+0(%rip) */
	0x4c, 0x03, 0x47, 0x08, //0x00000168 addq         $8(%rdi), %r8
	0x4c, 0x39, 0xe6, //0x0000016c cmpq         %r12, %rsi
	0x0f, 0x82, 0x30, 0x02, 0x00, 0x00, //0x0000016f jb           LBB0_5
	//0x00000175 LBB0_6
	0xc5, 0xfe, 0x6f, 0x0d, 0xc3, 0xfe, 0xff, 0xff, //0x00000175 vmovdqu      $-317(%rip), %ymm1  /* LCPI0_2+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x15, 0xdb, 0xfe, 0xff, 0xff, //0x0000017d vmovdqu      $-293(%rip), %ymm2  /* LCPI0_3+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x1d, 0xf3, 0xfe, 0xff, 0xff, //0x00000185 vmovdqu      $-269(%rip), %ymm3  /* LCPI0_4+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x25, 0x0b, 0xff, 0xff, 0xff, //0x0000018d vmovdqu      $-245(%rip), %ymm4  /* LCPI0_5+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x2d, 0x23, 0xff, 0xff, 0xff, //0x00000195 vmovdqu      $-221(%rip), %ymm5  /* LCPI0_6+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x35, 0x3b, 0xff, 0xff, 0xff, //0x0000019d vmovdqu      $-197(%rip), %ymm6  /* LCPI0_7+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x3d, 0x53, 0xff, 0xff, 0xff, //0x000001a5 vmovdqu      $-173(%rip), %ymm7  /* LCPI0_8+0(%rip) */
	0x4d, 0x89, 0xc6, //0x000001ad movq         %r8, %r14
	//0x000001b0 .p2align 4, 0x90
	//0x000001b0 LBB0_7
	0xc4, 0x41, 0x7a, 0x6f, 0x04, 0x24, //0x000001b0 vmovdqu      (%r12), %xmm8
	0xc4, 0x43, 0x3d, 0x38, 0x44, 0x24, 0x0c, 0x01, //0x000001b6 vinserti128  $1, $12(%r12), %ymm8, %ymm8
	0xc4, 0x62, 0x3d, 0x00, 0xc1, //0x000001be vpshufb      %ymm1, %ymm8, %ymm8
	0xc5, 0x3d, 0xdb, 0xca, //0x000001c3 vpand        %ymm2, %ymm8, %ymm9
	0xc5, 0x35, 0xe4, 0xcb, //0x000001c7 vpmulhuw     %ymm3, %ymm9, %ymm9
	0xc5, 0x3d, 0xdb, 0xc4, //0x000001cb vpand        %ymm4, %ymm8, %ymm8
	0xc4, 0xc1, 0x2d, 0x71, 0xf0, 0x08, //0x000001cf vpsllw       $8, %ymm8, %ymm10
	0xc4, 0xc1, 0x3d, 0x71, 0xf0, 0x04, //0x000001d5 vpsllw       $4, %ymm8, %ymm8
	0xc4, 0x43, 0x3d, 0x0e, 0xc2, 0xaa, //0x000001db vpblendw     $170, %ymm10, %ymm8, %ymm8
	0xc4, 0x41, 0x3d, 0xeb, 0xc1, //0x000001e1 vpor         %ymm9, %ymm8, %ymm8
	0xc4, 0x41, 0x55, 0x64, 0xc8, //0x000001e6 vpcmpgtb     %ymm8, %ymm5, %ymm9
	0xc5, 0x35, 0xdb, 0xcf, //0x000001eb vpand        %ymm7, %ymm9, %ymm9
	0xc5, 0x3d, 0xd8, 0xd6, //0x000001ef vpsubusb     %ymm6, %ymm8, %ymm10
	0xc4, 0x41, 0x35, 0xeb, 0xca, //0x000001f3 vpor         %ymm10, %ymm9, %ymm9
	0xc4, 0x42, 0x7d, 0x00, 0xc9, //0x000001f8 vpshufb      %ymm9, %ymm0, %ymm9
	0xc4, 0x41, 0x35, 0xfc, 0xc0, //0x000001fd vpaddb       %ymm8, %ymm9, %ymm8
	0xc4, 0x41, 0x7e, 0x7f, 0x06, //0x00000202 vmovdqu      %ymm8, (%r14)
	0x49, 0x83, 0xc6, 0x20, //0x00000207 addq         $32, %r14
	0x49, 0x83, 0xc4, 0x18, //0x0000020b addq         $24, %r12
	0x49, 0x39, 0xf4, //0x0000020f cmpq         %rsi, %r12
	0x0f, 0x86, 0x98, 0xff, 0xff, 0xff, //0x00000212 jbe          LBB0_7
	0x49, 0x8d, 0x71, 0xe8, //0x00000218 leaq         $-24(%r9), %rsi
	0x49, 0x39, 0xf4, //0x0000021c cmpq         %rsi, %r12
	0x0f, 0x87, 0x83, 0x00, 0x00, 0x00, //0x0000021f ja           LBB0_10
	//0x00000225 LBB0_9
	0xc4, 0xc1, 0x7a, 0x6f, 0x0c, 0x24, //0x00000225 vmovdqu      (%r12), %xmm1
	0xc4, 0xc1, 0x7a, 0x6f, 0x54, 0x24, 0x08, //0x0000022b vmovdqu      $8(%r12), %xmm2
	0xc5, 0xe9, 0x73, 0xda, 0x04, //0x00000232 vpsrldq      $4, %xmm2, %xmm2
	0xc4, 0xe3, 0x75, 0x38, 0xca, 0x01, //0x00000237 vinserti128  $1, %xmm2, %ymm1, %ymm1
	0xc4, 0xe2, 0x75, 0x00, 0x0d, 0xfa, 0xfd, 0xff, 0xff, //0x0000023d vpshufb      $-518(%rip), %ymm1, %ymm1  /* LCPI0_2+0(%rip) */
	0xc5, 0xf5, 0xdb, 0x15, 0x12, 0xfe, 0xff, 0xff, //0x00000246 vpand        $-494(%rip), %ymm1, %ymm2  /* LCPI0_3+0(%rip) */
	0xc5, 0xed, 0xe4, 0x15, 0x2a, 0xfe, 0xff, 0xff, //0x0000024e vpmulhuw     $-470(%rip), %ymm2, %ymm2  /* LCPI0_4+0(%rip) */
	0xc5, 0xf5, 0xdb, 0x0d, 0x42, 0xfe, 0xff, 0xff, //0x00000256 vpand        $-446(%rip), %ymm1, %ymm1  /* LCPI0_5+0(%rip) */
	0xc5, 0xe5, 0x71, 0xf1, 0x08, //0x0000025e vpsllw       $8, %ymm1, %ymm3
	0xc5, 0xf5, 0x71, 0xf1, 0x04, //0x00000263 vpsllw       $4, %ymm1, %ymm1
	0xc4, 0xe3, 0x75, 0x0e, 0xcb, 0xaa, //0x00000268 vpblendw     $170, %ymm3, %ymm1, %ymm1
	0xc5, 0xf5, 0xeb, 0xca, //0x0000026e vpor         %ymm2, %ymm1, %ymm1
	0xc5, 0xfe, 0x6f, 0x15, 0x46, 0xfe, 0xff, 0xff, //0x00000272 vmovdqu      $-442(%rip), %ymm2  /* LCPI0_6+0(%rip) */
	0xc5, 0xed, 0x64, 0xd1, //0x0000027a vpcmpgtb     %ymm1, %ymm2, %ymm2
	0xc5, 0xf5, 0xd8, 0x1d, 0x5a, 0xfe, 0xff, 0xff, //0x0000027e vpsubusb     $-422(%rip), %ymm1, %ymm3  /* LCPI0_7+0(%rip) */
	0xc5, 0xed, 0xdb, 0x15, 0x72, 0xfe, 0xff, 0xff, //0x00000286 vpand        $-398(%rip), %ymm2, %ymm2  /* LCPI0_8+0(%rip) */
	0xc5, 0xed, 0xeb, 0xd3, //0x0000028e vpor         %ymm3, %ymm2, %ymm2
	0xc4, 0xe2, 0x7d, 0x00, 0xc2, //0x00000292 vpshufb      %ymm2, %ymm0, %ymm0
	0xc5, 0xfd, 0xfc, 0xc1, //0x00000297 vpaddb       %ymm1, %ymm0, %ymm0
	0xc4, 0xc1, 0x7e, 0x7f, 0x06, //0x0000029b vmovdqu      %ymm0, (%r14)
	0x49, 0x83, 0xc6, 0x20, //0x000002a0 addq         $32, %r14
	0x49, 0x83, 0xc4, 0x18, //0x000002a4 addq         $24, %r12
	//0x000002a8 LBB0_10
	0x4d, 0x39, 0xcc, //0x000002a8 cmpq         %r9, %r12
	0x0f, 0x84, 0x91, 0x01, 0x00, 0x00, //0x000002ab je           LBB0_25
	0x4d, 0x8d, 0x59, 0xfc, //0x000002b1 leaq         $-4(%r9), %r11
	0x4d, 0x39, 0xdc, //0x000002b5 cmpq         %r11, %r12
	0x0f, 0x87, 0x59, 0x00, 0x00, 0x00, //0x000002b8 ja           LBB0_14
	0x90, 0x90, //0x000002be .p2align 4, 0x90
	//0x000002c0 LBB0_12
	0x41, 0x8b, 0x34, 0x24, //0x000002c0 movl         (%r12), %esi
	0x0f, 0xce, //0x000002c4 bswapl       %esi
	0x49, 0x89, 0xf7, //0x000002c6 movq         %rsi, %r15
	0x49, 0xc1, 0xef, 0x1a, //0x000002c9 shrq         $26, %r15
	0x89, 0xf1, //0x000002cd movl         %esi, %ecx
	0xc1, 0xe9, 0x14, //0x000002cf shrl         $20, %ecx
	0x83, 0xe1, 0x3f, //0x000002d2 andl         $63, %ecx
	0x89, 0xf3, //0x000002d5 movl         %esi, %ebx
	0xc1, 0xeb, 0x0e, //0x000002d7 shrl         $14, %ebx
	0x83, 0xe3, 0x3f, //0x000002da andl         $63, %ebx
	0xc1, 0xee, 0x08, //0x000002dd shrl         $8, %esi
	0x83, 0xe6, 0x3f, //0x000002e0 andl         $63, %esi
	0x49, 0x83, 0xc4, 0x03, //0x000002e3 addq         $3, %r12
	0x43, 0x0f, 0xb6, 0x04, 0x3a, //0x000002e7 movzbl       (%r10,%r15), %eax
	0x41, 0x88, 0x06, //0x000002ec movb         %al, (%r14)
	0x41, 0x0f, 0xb6, 0x04, 0x0a, //0x000002ef movzbl       (%r10,%rcx), %eax
	0x41, 0x88, 0x46, 0x01, //0x000002f4 movb         %al, $1(%r14)
	0x41, 0x0f, 0xb6, 0x04, 0x1a, //0x000002f8 movzbl       (%r10,%rbx), %eax
	0x41, 0x88, 0x46, 0x02, //0x000002fd movb         %al, $2(%r14)
	0x41, 0x0f, 0xb6, 0x04, 0x32, //0x00000301 movzbl       (%r10,%rsi), %eax
	0x41, 0x88, 0x46, 0x03, //0x00000306 movb         %al, $3(%r14)
	0x49, 0x83, 0xc6, 0x04, //0x0000030a addq         $4, %r14
	0x4d, 0x39, 0xdc, //0x0000030e cmpq         %r11, %r12
	0x0f, 0x86, 0xa9, 0xff, 0xff, 0xff, //0x00000311 jbe          LBB0_12
	//0x00000317 LBB0_14
	0x4d, 0x29, 0xe1, //0x00000317 subq         %r12, %r9
	0x45, 0x0f, 0xb6, 0x1c, 0x24, //0x0000031a movzbl       (%r12), %r11d
	0x49, 0x83, 0xf9, 0x01, //0x0000031f cmpq         $1, %r9
	0x0f, 0x84, 0xd5, 0x00, 0x00, 0x00, //0x00000323 je           LBB0_21
	0x4c, 0x89, 0xde, //0x00000329 movq         %r11, %rsi
	0x48, 0xc1, 0xe6, 0x10, //0x0000032c shlq         $16, %rsi
	0x49, 0x83, 0xf9, 0x02, //0x00000330 cmpq         $2, %r9
	0x0f, 0x84, 0x80, 0x00, 0x00, 0x00, //0x00000334 je           LBB0_18
	0x49, 0x83, 0xf9, 0x03, //0x0000033a cmpq         $3, %r9
	0x0f, 0x85, 0xfe, 0x00, 0x00, 0x00, //0x0000033e jne          LBB0_25
	0x41, 0x0f, 0xb6, 0x54, 0x24, 0x02, //0x00000344 movzbl       $2(%r12), %edx
	0x09, 0xd6, //0x0000034a orl          %edx, %esi
	0x41, 0x0f, 0xb6, 0x44, 0x24, 0x01, //0x0000034c movzbl       $1(%r12), %eax
	0xc1, 0xe0, 0x08, //0x00000352 shll         $8, %eax
	0x09, 0xf0, //0x00000355 orl          %esi, %eax
	0x49, 0xc1, 0xeb, 0x02, //0x00000357 shrq         $2, %r11
	0x43, 0x8a, 0x0c, 0x1a, //0x0000035b movb         (%r10,%r11), %cl
	0x41, 0x88, 0x0e, //0x0000035f movb         %cl, (%r14)
	0x89, 0xc1, //0x00000362 movl         %eax, %ecx
	0xc1, 0xe9, 0x0c, //0x00000364 shrl         $12, %ecx
	0x83, 0xe1, 0x3f, //0x00000367 andl         $63, %ecx
	0x41, 0x8a, 0x0c, 0x0a, //0x0000036a movb         (%r10,%rcx), %cl
	0x41, 0x88, 0x4e, 0x01, //0x0000036e movb         %cl, $1(%r14)
	0xc1, 0xe8, 0x06, //0x00000372 shrl         $6, %eax
	0x83, 0xe0, 0x3f, //0x00000375 andl         $63, %eax
	0x41, 0x8a, 0x04, 0x02, //0x00000378 movb         (%r10,%rax), %al
	0x41, 0x88, 0x46, 0x02, //0x0000037c movb         %al, $2(%r14)
	0x83, 0xe2, 0x3f, //0x00000380 andl         $63, %edx
	0x41, 0x8a, 0x04, 0x12, //0x00000383 movb         (%r10,%rdx), %al
	0x41, 0x88, 0x46, 0x03, //0x00000387 movb         %al, $3(%r14)
	0xe9, 0x9c, 0x00, 0x00, 0x00, //0x0000038b jmp          LBB0_24
	//0x00000390 LBB0_2
	0xc5, 0xfe, 0x6f, 0x05, 0x68, 0xfc, 0xff, 0xff, //0x00000390 vmovdqu      $-920(%rip), %ymm0  /* LCPI0_0+0(%rip) */
	0x4c, 0x03, 0x47, 0x08, //0x00000398 addq         $8(%rdi), %r8
	0x4c, 0x39, 0xe6, //0x0000039c cmpq         %r12, %rsi
	0x0f, 0x83, 0xd0, 0xfd, 0xff, 0xff, //0x0000039f jae          LBB0_6
	//0x000003a5 LBB0_5
	0x4d, 0x89, 0xc6, //0x000003a5 movq         %r8, %r14
	0x49, 0x8d, 0x71, 0xe8, //0x000003a8 leaq         $-24(%r9), %rsi
	0x49, 0x39, 0xf4, //0x000003ac cmpq         %rsi, %r12
	0x0f, 0x86, 0x70, 0xfe, 0xff, 0xff, //0x000003af jbe          LBB0_9
	0xe9, 0xee, 0xfe, 0xff, 0xff, //0x000003b5 jmp          LBB0_10
	//0x000003ba LBB0_18
	0x41, 0x0f, 0xb6, 0x44, 0x24, 0x01, //0x000003ba movzbl       $1(%r12), %eax
	0x89, 0xc1, //0x000003c0 movl         %eax, %ecx
	0xc1, 0xe1, 0x08, //0x000003c2 shll         $8, %ecx
	0x09, 0xf1, //0x000003c5 orl          %esi, %ecx
	0x49, 0xc1, 0xeb, 0x02, //0x000003c7 shrq         $2, %r11
	0x43, 0x8a, 0x1c, 0x1a, //0x000003cb movb         (%r10,%r11), %bl
	0x41, 0x88, 0x1e, //0x000003cf movb         %bl, (%r14)
	0xc1, 0xe9, 0x0c, //0x000003d2 shrl         $12, %ecx
	0x83, 0xe1, 0x3f, //0x000003d5 andl         $63, %ecx
	0x41, 0x8a, 0x0c, 0x0a, //0x000003d8 movb         (%r10,%rcx), %cl
	0x41, 0x88, 0x4e, 0x01, //0x000003dc movb         %cl, $1(%r14)
	0x83, 0xe0, 0x0f, //0x000003e0 andl         $15, %eax
	0x41, 0x8a, 0x04, 0x82, //0x000003e3 movb         (%r10,%rax,4), %al
	0x41, 0x88, 0x46, 0x02, //0x000003e7 movb         %al, $2(%r14)
	0xf6, 0xc2, 0x02, //0x000003eb testb        $2, %dl
	0x0f, 0x85, 0x41, 0x00, 0x00, 0x00, //0x000003ee jne          LBB0_19
	0x41, 0xc6, 0x46, 0x03, 0x3d, //0x000003f4 movb         $61, $3(%r14)
	0xe9, 0x2e, 0x00, 0x00, 0x00, //0x000003f9 jmp          LBB0_24
	//0x000003fe LBB0_21
	0x4c, 0x89, 0xd8, //0x000003fe movq         %r11, %rax
	0x48, 0xc1, 0xe8, 0x02, //0x00000401 shrq         $2, %rax
	0x41, 0x8a, 0x04, 0x02, //0x00000405 movb         (%r10,%rax), %al
	0x41, 0x88, 0x06, //0x00000409 movb         %al, (%r14)
	0x41, 0xc1, 0xe3, 0x04, //0x0000040c shll         $4, %r11d
	0x41, 0x83, 0xe3, 0x30, //0x00000410 andl         $48, %r11d
	0x43, 0x8a, 0x04, 0x1a, //0x00000414 movb         (%r10,%r11), %al
	0x41, 0x88, 0x46, 0x01, //0x00000418 movb         %al, $1(%r14)
	0xf6, 0xc2, 0x02, //0x0000041c testb        $2, %dl
	0x0f, 0x85, 0x19, 0x00, 0x00, 0x00, //0x0000041f jne          LBB0_22
	0x66, 0x41, 0xc7, 0x46, 0x02, 0x3d, 0x3d, //0x00000425 movw         $15677, $2(%r14)
	//0x0000042c LBB0_24
	0x49, 0x83, 0xc6, 0x04, //0x0000042c addq         $4, %r14
	0xe9, 0x0d, 0x00, 0x00, 0x00, //0x00000430 jmp          LBB0_25
	//0x00000435 LBB0_19
	0x49, 0x83, 0xc6, 0x03, //0x00000435 addq         $3, %r14
	0xe9, 0x04, 0x00, 0x00, 0x00, //0x00000439 jmp          LBB0_25
	//0x0000043e LBB0_22
	0x49, 0x83, 0xc6, 0x02, //0x0000043e addq         $2, %r14
	//0x00000442 LBB0_25
	0x4d, 0x29, 0xc6, //0x00000442 subq         %r8, %r14
	0x4c, 0x01, 0x77, 0x08, //0x00000445 addq         %r14, $8(%rdi)
	//0x00000449 LBB0_26
	0x5b, //0x00000449 popq         %rbx
	0x41, 0x5c, //0x0000044a popq         %r12
	0x41, 0x5e, //0x0000044c popq         %r14
	0x41, 0x5f, //0x0000044e popq         %r15
	0x5d, //0x00000450 popq         %rbp
	0xc5, 0xf8, 0x77, //0x00000451 vzeroupper   
	0xc3, //0x00000454 retq         
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x00000455 .p2align 4, 0x00
	//0x00000460 _TabEncodeCharsetStd
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, //0x00000460 QUAD $0x4847464544434241; QUAD $0x504f4e4d4c4b4a49  // .ascii 16, 'ABCDEFGHIJKLMNOP'
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, //0x00000470 QUAD $0x5857565554535251; QUAD $0x6665646362615a59  // .ascii 16, 'QRSTUVWXYZabcdef'
	0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, //0x00000480 QUAD $0x6e6d6c6b6a696867; QUAD $0x767574737271706f  // .ascii 16, 'ghijklmnopqrstuv'
	0x77, 0x78, 0x79, 0x7a, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x2b, 0x2f, //0x00000490 QUAD $0x333231307a797877; QUAD $0x2f2b393837363534  // .ascii 16, 'wxyz0123456789+/'
	//0x000004a0 .p2align 4, 0x00
	//0x000004a0 _TabEncodeCharsetURL
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, //0x000004a0 QUAD $0x4847464544434241; QUAD $0x504f4e4d4c4b4a49  // .ascii 16, 'ABCDEFGHIJKLMNOP'
	0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, //0x000004b0 QUAD $0x5857565554535251; QUAD $0x6665646362615a59  // .ascii 16, 'QRSTUVWXYZabcdef'
	0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, //0x000004c0 QUAD $0x6e6d6c6b6a696867; QUAD $0x767574737271706f  // .ascii 16, 'ghijklmnopqrstuv'
	0x77, 0x78, 0x79, 0x7a, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x2d, 0x5f, //0x000004d0 QUAD $0x333231307a797877; QUAD $0x5f2d393837363534  // .ascii 16, 'wxyz0123456789-_'
}
 