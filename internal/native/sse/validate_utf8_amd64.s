// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__validate_utf8_entry(SB), NOSPLIT, $40
	NO_LOCAL_POINTERS
	LONG $0xeb0d8d4c; WORD $0xffff; BYTE $0xff  // leaq         $-21(%rip), %r9
	LONG $0x244c894c; BYTE $0x38  // movq         %r9, $56(%rsp)
	LONG $0x30c48348  // addq         $48, %rsp
	BYTE $0xc3  // retq         
	BYTE $0x00
	  // .p2align 4, 0x90
_validate_utf8:
	BYTE $0x55  // pushq        %rbp
	WORD $0x8948; BYTE $0xe5  // movq         %rsp, %rbp
	WORD $0x5741  // pushq        %r15
	WORD $0x5641  // pushq        %r14
	WORD $0x5441  // pushq        %r12
	BYTE $0x53  // pushq        %rbx
	BYTE $0x50  // pushq        %rax
	WORD $0x8b4c; BYTE $0x17  // movq         (%rdi), %r10
	LONG $0x085f8b4c  // movq         $8(%rdi), %r11
	WORD $0x8b48; BYTE $0x0e  // movq         (%rsi), %rcx
	WORD $0x014c; BYTE $0xd1  // addq         %r10, %rcx
	LONG $0x1a448d4f; BYTE $0xfd  // leaq         $-3(%r10,%r11), %r8
	LONG $0x000010e9; BYTE $0x00  // jmp          LBB0_1, $16(%rip)
	QUAD $0x9090909090909090; LONG $0x90909090; BYTE $0x90  // .p2align 4, 0x90
LBB0_19:
	WORD $0x0148; BYTE $0xd9  // addq         %rbx, %rcx
LBB0_1:
	WORD $0x394c; BYTE $0xc1  // cmpq         %r8, %rcx
	LONG $0x00e1830f; WORD $0x0000  // jae          LBB0_2, $225(%rip)
	LONG $0x000001bb; BYTE $0x00  // movl         $1, %ebx
	WORD $0x3980; BYTE $0x00  // cmpb         $0, (%rcx)
	LONG $0xffe6890f; WORD $0xffff  // jns          LBB0_19, $-26(%rip)
	WORD $0x018b  // movl         (%rcx), %eax
	WORD $0xc789  // movl         %eax, %edi
	LONG $0xc0f0e781; WORD $0x00c0  // andl         $12632304, %edi
	LONG $0x80e0ff81; WORD $0x0080  // cmpl         $8421600, %edi
	LONG $0x0030850f; WORD $0x0000  // jne          LBB0_10, $48(%rip)
	WORD $0xc789  // movl         %eax, %edi
	LONG $0x200fe781; WORD $0x0000  // andl         $8207, %edi
	LONG $0x200dff81; WORD $0x0000  // cmpl         $8205, %edi
	LONG $0x001c840f; WORD $0x0000  // je           LBB0_10, $28(%rip)
	LONG $0x000003bb; BYTE $0x00  // movl         $3, %ebx
	WORD $0xff85  // testl        %edi, %edi
	LONG $0xffaf850f; WORD $0xffff  // jne          LBB0_19, $-81(%rip)
	QUAD $0x9090909090909090; LONG $0x90909090; WORD $0x9090; BYTE $0x90  // .p2align 4, 0x90
LBB0_10:
	WORD $0xc789  // movl         %eax, %edi
	LONG $0xc0e0e781; WORD $0x0000  // andl         $49376, %edi
	LONG $0x80c0ff81; WORD $0x0000  // cmpl         $32960, %edi
	LONG $0x0010850f; WORD $0x0000  // jne          LBB0_12, $16(%rip)
	WORD $0xc789  // movl         %eax, %edi
	LONG $0x000002bb; BYTE $0x00  // movl         $2, %ebx
	WORD $0xe783; BYTE $0x1e  // andl         $30, %edi
	LONG $0xff7c850f; WORD $0xffff  // jne          LBB0_19, $-132(%rip)
LBB0_12:
	WORD $0xc789  // movl         %eax, %edi
	LONG $0xc0f8e781; WORD $0xc0c0  // andl         $-1061109512, %edi
	LONG $0x80f0ff81; WORD $0x8080  // cmpl         $-2139062032, %edi
	LONG $0x0026850f; WORD $0x0000  // jne          LBB0_16, $38(%rip)
	WORD $0xc789  // movl         %eax, %edi
	LONG $0x3007e781; WORD $0x0000  // andl         $12295, %edi
	LONG $0x0018840f; WORD $0x0000  // je           LBB0_16, $24(%rip)
	LONG $0x000004bb; BYTE $0x00  // movl         $4, %ebx
	WORD $0x04a8  // testb        $4, %al
	LONG $0xff4d840f; WORD $0xffff  // je           LBB0_19, $-179(%rip)
	LONG $0x00300325; BYTE $0x00  // andl         $12291, %eax
	LONG $0xff42840f; WORD $0xffff  // je           LBB0_19, $-190(%rip)
LBB0_16:
	WORD $0x8948; BYTE $0xcf  // movq         %rcx, %rdi
	WORD $0x294c; BYTE $0xd7  // subq         %r10, %rdi
	WORD $0x8b48; BYTE $0x1a  // movq         (%rdx), %rbx
	LONG $0x00fb8148; WORD $0x0010; BYTE $0x00  // cmpq         $4096, %rbx
	LONG $0x0187830f; WORD $0x0000  // jae          LBB0_17, $391(%rip)
	WORD $0x6348; BYTE $0xc7  // movslq       %edi, %rax
	LONG $0x017b8d48  // leaq         $1(%rbx), %rdi
	WORD $0x8948; BYTE $0x3a  // movq         %rdi, (%rdx)
	LONG $0xda448948; BYTE $0x08  // movq         %rax, $8(%rdx,%rbx,8)
	LONG $0x000001bb; BYTE $0x00  // movl         $1, %ebx
	LONG $0xffff13e9; BYTE $0xff  // jmp          LBB0_19, $-237(%rip)
LBB0_2:
	WORD $0x014d; BYTE $0xd3  // addq         %r10, %r11
	WORD $0x394c; BYTE $0xd9  // cmpq         %r11, %rcx
	LONG $0x013e830f; WORD $0x0000  // jae          LBB0_36, $318(%rip)
	LONG $0xdc458d4c  // leaq         $-36(%rbp), %r8
	LONG $0xda4d8d4c  // leaq         $-38(%rbp), %r9
	LONG $0x000016e9; BYTE $0x00  // jmp          LBB0_4, $22(%rip)
	QUAD $0x9090909090909090; WORD $0x9090  // .p2align 4, 0x90
LBB0_5:
	WORD $0xff48; BYTE $0xc1  // incq         %rcx
	WORD $0x394c; BYTE $0xd9  // cmpq         %r11, %rcx
	LONG $0x011b830f; WORD $0x0000  // jae          LBB0_36, $283(%rip)
LBB0_4:
	WORD $0x3980; BYTE $0x00  // cmpb         $0, (%rcx)
	LONG $0xffeb890f; WORD $0xffff  // jns          LBB0_5, $-21(%rip)
	LONG $0x00dc45c6  // movb         $0, $-36(%rbp)
	LONG $0x00da45c6  // movb         $0, $-38(%rbp)
	WORD $0x894c; BYTE $0xdb  // movq         %r11, %rbx
	WORD $0x2948; BYTE $0xcb  // subq         %rcx, %rbx
	LONG $0x02fb8348  // cmpq         $2, %rbx
	LONG $0x0035820f; WORD $0x0000  // jb           LBB0_21, $53(%rip)
	LONG $0x21b60f44  // movzbl       (%rcx), %r12d
	LONG $0x71b60f44; BYTE $0x01  // movzbl       $1(%rcx), %r14d
	LONG $0xdc658844  // movb         %r12b, $-36(%rbp)
	LONG $0x02798d4c  // leaq         $2(%rcx), %r15
	LONG $0xfec38348  // addq         $-2, %rbx
	WORD $0x894c; BYTE $0xcf  // movq         %r9, %rdi
	WORD $0x8548; BYTE $0xdb  // testq        %rbx, %rbx
	LONG $0x0029840f; WORD $0x0000  // je           LBB0_24, $41(%rip)
LBB0_25:
	LONG $0x07b60f41  // movzbl       (%r15), %eax
	WORD $0x0788  // movb         %al, (%rdi)
	LONG $0x65b60f44; BYTE $0xdc  // movzbl       $-36(%rbp), %r12d
	LONG $0xda7db60f  // movzbl       $-38(%rbp), %edi
	LONG $0x000017e9; BYTE $0x00  // jmp          LBB0_26, $23(%rip)
LBB0_21:
	WORD $0x3145; BYTE $0xe4  // xorl         %r12d, %r12d
	WORD $0x3145; BYTE $0xf6  // xorl         %r14d, %r14d
	WORD $0x894c; BYTE $0xc7  // movq         %r8, %rdi
	WORD $0x8949; BYTE $0xcf  // movq         %rcx, %r15
	WORD $0x8548; BYTE $0xdb  // testq        %rbx, %rbx
	LONG $0xffd7850f; WORD $0xffff  // jne          LBB0_25, $-41(%rip)
LBB0_24:
	WORD $0xff31  // xorl         %edi, %edi
LBB0_26:
	LONG $0xc7b60f40  // movzbl       %dil, %eax
	WORD $0xe0c1; BYTE $0x10  // shll         $16, %eax
	LONG $0xdeb60f41  // movzbl       %r14b, %ebx
	WORD $0xe3c1; BYTE $0x08  // shll         $8, %ebx
	LONG $0xfcb60f41  // movzbl       %r12b, %edi
	WORD $0xdf09  // orl          %ebx, %edi
	WORD $0xf809  // orl          %edi, %eax
	LONG $0xc0c0f025; BYTE $0x00  // andl         $12632304, %eax
	LONG $0x8080e03d; BYTE $0x00  // cmpl         $8421600, %eax
	LONG $0x0021850f; WORD $0x0000  // jne          LBB0_29, $33(%rip)
	WORD $0xf889  // movl         %edi, %eax
	LONG $0x00200f25; BYTE $0x00  // andl         $8207, %eax
	LONG $0x00200d3d; BYTE $0x00  // cmpl         $8205, %eax
	LONG $0x000f840f; WORD $0x0000  // je           LBB0_29, $15(%rip)
	LONG $0x000003bb; BYTE $0x00  // movl         $3, %ebx
	WORD $0xc085  // testl        %eax, %eax
	LONG $0x0023850f; WORD $0x0000  // jne          LBB0_34, $35(%rip)
	WORD $0x9090  // .p2align 4, 0x90
LBB0_29:
	LONG $0x1ec4f641  // testb        $30, %r12b
	LONG $0x0028840f; WORD $0x0000  // je           LBB0_31, $40(%rip)
	LONG $0xc0e0e781; WORD $0x0000  // andl         $49376, %edi
	LONG $0x000002bb; BYTE $0x00  // movl         $2, %ebx
	LONG $0x80c0ff81; WORD $0x0000  // cmpl         $32960, %edi
	LONG $0x0011850f; WORD $0x0000  // jne          LBB0_31, $17(%rip)
LBB0_34:
	WORD $0x0148; BYTE $0xd9  // addq         %rbx, %rcx
	WORD $0x394c; BYTE $0xd9  // cmpq         %r11, %rcx
	LONG $0xff1f820f; WORD $0xffff  // jb           LBB0_4, $-225(%rip)
	LONG $0x000035e9; BYTE $0x00  // jmp          LBB0_36, $53(%rip)
LBB0_31:
	WORD $0x8948; BYTE $0xc8  // movq         %rcx, %rax
	WORD $0x294c; BYTE $0xd0  // subq         %r10, %rax
	WORD $0x8b48; BYTE $0x3a  // movq         (%rdx), %rdi
	LONG $0x00ff8148; WORD $0x0010; BYTE $0x00  // cmpq         $4096, %rdi
	LONG $0x0034830f; WORD $0x0000  // jae          LBB0_32, $52(%rip)
	WORD $0x9848  // cltq         
	LONG $0x015f8d48  // leaq         $1(%rdi), %rbx
	WORD $0x8948; BYTE $0x1a  // movq         %rbx, (%rdx)
	LONG $0xfa448948; BYTE $0x08  // movq         %rax, $8(%rdx,%rdi,8)
	LONG $0x000001bb; BYTE $0x00  // movl         $1, %ebx
	WORD $0x0148; BYTE $0xd9  // addq         %rbx, %rcx
	WORD $0x394c; BYTE $0xd9  // cmpq         %r11, %rcx
	LONG $0xfee5820f; WORD $0xffff  // jb           LBB0_4, $-283(%rip)
LBB0_36:
	WORD $0x294c; BYTE $0xd1  // subq         %r10, %rcx
	WORD $0x8948; BYTE $0x0e  // movq         %rcx, (%rsi)
	WORD $0xc031  // xorl         %eax, %eax
LBB0_37:
	LONG $0x08c48348  // addq         $8, %rsp
	BYTE $0x5b  // popq         %rbx
	WORD $0x5c41  // popq         %r12
	WORD $0x5e41  // popq         %r14
	WORD $0x5f41  // popq         %r15
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_32:
	WORD $0x8948; BYTE $0x06  // movq         %rax, (%rsi)
	LONG $0xffc0c748; WORD $0xffff; BYTE $0xff  // movq         $-1, %rax
	LONG $0xffffe4e9; BYTE $0xff  // jmp          LBB0_37, $-28(%rip)
LBB0_17:
	WORD $0x8948; BYTE $0x3e  // movq         %rdi, (%rsi)
	LONG $0xffc0c748; WORD $0xffff; BYTE $0xff  // movq         $-1, %rax
	LONG $0xffffd5e9; BYTE $0xff  // jmp          LBB0_37, $-43(%rip)
	WORD $0x0000  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	LONG $0x00000002  // .long 2

TEXT ·__validate_utf8(SB), NOSPLIT | NOFRAME, $0 - 32
	NO_LOCAL_POINTERS

_entry:
	MOVQ (TLS), R14
	LEAQ -48(SP), R12
	CMPQ R12, 16(R14)
	JBE  _stack_grow

_validate_utf8:
	MOVQ s+0(FP), DI
	MOVQ p+8(FP), SI
	MOVQ m+16(FP), DX
	CALL ·__validate_utf8_entry+32(SB)  // _validate_utf8
	MOVQ AX, ret+24(FP)
	RET

_stack_grow:
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
