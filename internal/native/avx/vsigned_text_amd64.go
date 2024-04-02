// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package avx

var _text_vsigned = []byte{
	// .p2align 4, 0x90
	// _vsigned
	0x55, // pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000001 movq         %rsp, %rbp
	0x53, //0x00000004 pushq        %rbx
	0x48, 0x8b, 0x1e, //0x00000005 movq         (%rsi), %rbx
	0x4c, 0x8b, 0x07, //0x00000008 movq         (%rdi), %r8
	0x4c, 0x8b, 0x57, 0x08, //0x0000000b movq         $8(%rdi), %r10
	0x48, 0xc7, 0x02, 0x09, 0x00, 0x00, 0x00, //0x0000000f movq         $9, (%rdx)
	0x48, 0xc7, 0x42, 0x08, 0x00, 0x00, 0x00, 0x00, //0x00000016 movq         $0, $8(%rdx)
	0x48, 0xc7, 0x42, 0x10, 0x00, 0x00, 0x00, 0x00, //0x0000001e movq         $0, $16(%rdx)
	0x48, 0x8b, 0x0e, //0x00000026 movq         (%rsi), %rcx
	0x48, 0x89, 0x4a, 0x18, //0x00000029 movq         %rcx, $24(%rdx)
	0x4c, 0x39, 0xd3, //0x0000002d cmpq         %r10, %rbx
	0x0f, 0x83, 0x44, 0x00, 0x00, 0x00, //0x00000030 jae          LBB0_1
	0x41, 0x8a, 0x0c, 0x18, //0x00000036 movb         (%r8,%rbx), %cl
	0x41, 0xb9, 0x01, 0x00, 0x00, 0x00, //0x0000003a movl         $1, %r9d
	0x80, 0xf9, 0x2d, //0x00000040 cmpb         $45, %cl
	0x0f, 0x85, 0x17, 0x00, 0x00, 0x00, //0x00000043 jne          LBB0_5
	0x48, 0xff, 0xc3, //0x00000049 incq         %rbx
	0x4c, 0x39, 0xd3, //0x0000004c cmpq         %r10, %rbx
	0x0f, 0x83, 0x25, 0x00, 0x00, 0x00, //0x0000004f jae          LBB0_1
	0x41, 0x8a, 0x0c, 0x18, //0x00000055 movb         (%r8,%rbx), %cl
	0x49, 0xc7, 0xc1, 0xff, 0xff, 0xff, 0xff, //0x00000059 movq         $-1, %r9
	//0x00000060 LBB0_5
	0x8d, 0x79, 0xd0, //0x00000060 leal         $-48(%rcx), %edi
	0x40, 0x80, 0xff, 0x0a, //0x00000063 cmpb         $10, %dil
	0x0f, 0x82, 0x1a, 0x00, 0x00, 0x00, //0x00000067 jb           LBB0_7
	0x48, 0x89, 0x1e, //0x0000006d movq         %rbx, (%rsi)
	0x48, 0xc7, 0x02, 0xfe, 0xff, 0xff, 0xff, //0x00000070 movq         $-2, (%rdx)
	0x5b, //0x00000077 popq         %rbx
	0x5d, //0x00000078 popq         %rbp
	0xc3, //0x00000079 retq         
	//0x0000007a LBB0_1
	0x4c, 0x89, 0x16, //0x0000007a movq         %r10, (%rsi)
	0x48, 0xc7, 0x02, 0xff, 0xff, 0xff, 0xff, //0x0000007d movq         $-1, (%rdx)
	0x5b, //0x00000084 popq         %rbx
	0x5d, //0x00000085 popq         %rbp
	0xc3, //0x00000086 retq         
	//0x00000087 LBB0_7
	0x80, 0xf9, 0x30, //0x00000087 cmpb         $48, %cl
	0x0f, 0x85, 0x35, 0x00, 0x00, 0x00, //0x0000008a jne          LBB0_8
	0x48, 0x8d, 0x7b, 0x01, //0x00000090 leaq         $1(%rbx), %rdi
	0x4c, 0x39, 0xd3, //0x00000094 cmpq         %r10, %rbx
	0x0f, 0x83, 0x7a, 0x00, 0x00, 0x00, //0x00000097 jae          LBB0_17
	0x41, 0x8a, 0x0c, 0x38, //0x0000009d movb         (%r8,%rdi), %cl
	0x80, 0xc1, 0xd2, //0x000000a1 addb         $-46, %cl
	0x80, 0xf9, 0x37, //0x000000a4 cmpb         $55, %cl
	0x0f, 0x87, 0x6a, 0x00, 0x00, 0x00, //0x000000a7 ja           LBB0_17
	0x44, 0x0f, 0xb6, 0xd9, //0x000000ad movzbl       %cl, %r11d
	0x48, 0xb9, 0x01, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x00, //0x000000b1 movabsq      $36028797027352577, %rcx
	0x4c, 0x0f, 0xa3, 0xd9, //0x000000bb btq          %r11, %rcx
	0x0f, 0x83, 0x52, 0x00, 0x00, 0x00, //0x000000bf jae          LBB0_17
	//0x000000c5 LBB0_8
	0x31, 0xff, //0x000000c5 xorl         %edi, %edi
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000000c7 .p2align 4, 0x90
	//0x000000d0 LBB0_9
	0x4c, 0x39, 0xd3, //0x000000d0 cmpq         %r10, %rbx
	0x0f, 0x83, 0x6c, 0x00, 0x00, 0x00, //0x000000d3 jae          LBB0_22
	0x49, 0x0f, 0xbe, 0x0c, 0x18, //0x000000d9 movsbq       (%r8,%rbx), %rcx
	0x8d, 0x41, 0xd0, //0x000000de leal         $-48(%rcx), %eax
	0x3c, 0x09, //0x000000e1 cmpb         $9, %al
	0x0f, 0x87, 0x34, 0x00, 0x00, 0x00, //0x000000e3 ja           LBB0_18
	0x48, 0x6b, 0xff, 0x0a, //0x000000e9 imulq        $10, %rdi, %rdi
	0x0f, 0x80, 0x14, 0x00, 0x00, 0x00, //0x000000ed jo           LBB0_13
	0x48, 0xff, 0xc3, //0x000000f3 incq         %rbx
	0x48, 0x83, 0xc1, 0xd0, //0x000000f6 addq         $-48, %rcx
	0x49, 0x0f, 0xaf, 0xc9, //0x000000fa imulq        %r9, %rcx
	0x48, 0x01, 0xcf, //0x000000fe addq         %rcx, %rdi
	0x0f, 0x81, 0xc9, 0xff, 0xff, 0xff, //0x00000101 jno          LBB0_9
	//0x00000107 LBB0_13
	0x48, 0xff, 0xcb, //0x00000107 decq         %rbx
	0x48, 0x89, 0x1e, //0x0000010a movq         %rbx, (%rsi)
	0x48, 0xc7, 0x02, 0xfb, 0xff, 0xff, 0xff, //0x0000010d movq         $-5, (%rdx)
	0x5b, //0x00000114 popq         %rbx
	0x5d, //0x00000115 popq         %rbp
	0xc3, //0x00000116 retq         
	//0x00000117 LBB0_17
	0x48, 0x89, 0x3e, //0x00000117 movq         %rdi, (%rsi)
	0x5b, //0x0000011a popq         %rbx
	0x5d, //0x0000011b popq         %rbp
	0xc3, //0x0000011c retq         
	//0x0000011d LBB0_18
	0x80, 0xf9, 0x65, //0x0000011d cmpb         $101, %cl
	0x0f, 0x84, 0x12, 0x00, 0x00, 0x00, //0x00000120 je           LBB0_21
	0x80, 0xf9, 0x45, //0x00000126 cmpb         $69, %cl
	0x0f, 0x84, 0x09, 0x00, 0x00, 0x00, //0x00000129 je           LBB0_21
	0x80, 0xf9, 0x2e, //0x0000012f cmpb         $46, %cl
	0x0f, 0x85, 0x0d, 0x00, 0x00, 0x00, //0x00000132 jne          LBB0_22
	//0x00000138 LBB0_21
	0x48, 0x89, 0x1e, //0x00000138 movq         %rbx, (%rsi)
	0x48, 0xc7, 0x02, 0xfa, 0xff, 0xff, 0xff, //0x0000013b movq         $-6, (%rdx)
	0x5b, //0x00000142 popq         %rbx
	0x5d, //0x00000143 popq         %rbp
	0xc3, //0x00000144 retq         
	//0x00000145 LBB0_22
	0x48, 0x89, 0x1e, //0x00000145 movq         %rbx, (%rsi)
	0x48, 0x89, 0x7a, 0x10, //0x00000148 movq         %rdi, $16(%rdx)
	0x5b, //0x0000014c popq         %rbx
	0x5d, //0x0000014d popq         %rbp
	0xc3, //0x0000014e retq         
	0x00, //0x0000014f .p2align 2, 0x00
	//0x00000150 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x00000150 .long 2
}
 
