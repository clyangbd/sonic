// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT ·__validate_utf8_fast_entry(SB), NOSPLIT, $16
	NO_LOCAL_POINTERS
	LONG $0xeb0d8d4c; WORD $0xffff; BYTE $0xff  // leaq         $-21(%rip), %r9
	LONG $0x244c894c; BYTE $0x20  // movq         %r9, $32(%rsp)
	LONG $0x18c48348  // addq         $24, %rsp
	BYTE $0xc3  // retq         
	BYTE $0x00
	  // .p2align 4, 0x90
_validate_utf8_fast:
	BYTE $0x55  // pushq        %rbp
	WORD $0x8948; BYTE $0xe5  // movq         %rsp, %rbp
	BYTE $0x53  // pushq        %rbx
	BYTE $0x50  // pushq        %rax
	WORD $0x8b4c; BYTE $0x17  // movq         (%rdi), %r10
	LONG $0x085f8b4c  // movq         $8(%rdi), %r11
	LONG $0x1a748d4b; BYTE $0xfd  // leaq         $-3(%r10,%r11), %rsi
	WORD $0x894c; BYTE $0xd0  // movq         %r10, %rax
	WORD $0x3949; BYTE $0xf2  // cmpq         %rsi, %r10
	LONG $0x00e0830f; WORD $0x0000  // jae          LBB0_14, $224(%rip)
	WORD $0x894c; BYTE $0xd0  // movq         %r10, %rax
	LONG $0x000016e9; BYTE $0x00  // jmp          LBB0_3, $22(%rip)
	QUAD $0x9090909090909090; WORD $0x9090  // .p2align 4, 0x90
LBB0_2:
	WORD $0x0148; BYTE $0xd0  // addq         %rdx, %rax
	WORD $0x3948; BYTE $0xf0  // cmpq         %rsi, %rax
	LONG $0x00c2830f; WORD $0x0000  // jae          LBB0_14, $194(%rip)
LBB0_3:
	LONG $0x000001ba; BYTE $0x00  // movl         $1, %edx
	WORD $0x3880; BYTE $0x00  // cmpb         $0, (%rax)
	LONG $0xffe6890f; WORD $0xffff  // jns          LBB0_2, $-26(%rip)
	WORD $0x388b  // movl         (%rax), %edi
	WORD $0xf989  // movl         %edi, %ecx
	LONG $0xc0f0e181; WORD $0x00c0  // andl         $12632304, %ecx
	LONG $0x80e0f981; WORD $0x0080  // cmpl         $8421600, %ecx
	LONG $0x0030850f; WORD $0x0000  // jne          LBB0_7, $48(%rip)
	WORD $0xf989  // movl         %edi, %ecx
	LONG $0x200fe181; WORD $0x0000  // andl         $8207, %ecx
	LONG $0x200df981; WORD $0x0000  // cmpl         $8205, %ecx
	LONG $0x001c840f; WORD $0x0000  // je           LBB0_7, $28(%rip)
	LONG $0x000003ba; BYTE $0x00  // movl         $3, %edx
	WORD $0xc985  // testl        %ecx, %ecx
	LONG $0xffaf850f; WORD $0xffff  // jne          LBB0_2, $-81(%rip)
	QUAD $0x9090909090909090; LONG $0x90909090; WORD $0x9090; BYTE $0x90  // .p2align 4, 0x90
LBB0_7:
	WORD $0xf989  // movl         %edi, %ecx
	LONG $0xc0e0e181; WORD $0x0000  // andl         $49376, %ecx
	LONG $0x80c0f981; WORD $0x0000  // cmpl         $32960, %ecx
	LONG $0x0010850f; WORD $0x0000  // jne          LBB0_9, $16(%rip)
	WORD $0xf989  // movl         %edi, %ecx
	LONG $0x000002ba; BYTE $0x00  // movl         $2, %edx
	WORD $0xe183; BYTE $0x1e  // andl         $30, %ecx
	LONG $0xff7c850f; WORD $0xffff  // jne          LBB0_2, $-132(%rip)
LBB0_9:
	WORD $0xf989  // movl         %edi, %ecx
	LONG $0xc0f8e181; WORD $0xc0c0  // andl         $-1061109512, %ecx
	LONG $0x80f0f981; WORD $0x8080  // cmpl         $-2139062032, %ecx
	LONG $0x0029850f; WORD $0x0000  // jne          LBB0_13, $41(%rip)
	WORD $0xf989  // movl         %edi, %ecx
	LONG $0x3007e181; WORD $0x0000  // andl         $12295, %ecx
	LONG $0x001b840f; WORD $0x0000  // je           LBB0_13, $27(%rip)
	LONG $0x000004ba; BYTE $0x00  // movl         $4, %edx
	LONG $0x04c7f640  // testb        $4, %dil
	LONG $0xff4b840f; WORD $0xffff  // je           LBB0_2, $-181(%rip)
	LONG $0x3003e781; WORD $0x0000  // andl         $12291, %edi
	LONG $0xff3f840f; WORD $0xffff  // je           LBB0_2, $-193(%rip)
LBB0_13:
	WORD $0xf748; BYTE $0xd0  // notq         %rax
	WORD $0x014c; BYTE $0xd0  // addq         %r10, %rax
	LONG $0x08c48348  // addq         $8, %rsp
	BYTE $0x5b  // popq         %rbx
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
LBB0_14:
	WORD $0x014d; BYTE $0xd3  // addq         %r10, %r11
	WORD $0x394c; BYTE $0xd8  // cmpq         %r11, %rax
	LONG $0x0103830f; WORD $0x0000  // jae          LBB0_30, $259(%rip)
	LONG $0xf4458d4c  // leaq         $-12(%rbp), %r8
	LONG $0xf24d8d4c  // leaq         $-14(%rbp), %r9
	LONG $0x000015e9; BYTE $0x00  // jmp          LBB0_17, $21(%rip)
	QUAD $0x9090909090909090; BYTE $0x90  // .p2align 4, 0x90
LBB0_16:
	WORD $0xff48; BYTE $0xc0  // incq         %rax
	WORD $0x394c; BYTE $0xd8  // cmpq         %r11, %rax
	LONG $0x00e1830f; WORD $0x0000  // jae          LBB0_30, $225(%rip)
LBB0_17:
	WORD $0x3880; BYTE $0x00  // cmpb         $0, (%rax)
	LONG $0xffeb890f; WORD $0xffff  // jns          LBB0_16, $-21(%rip)
	LONG $0x00f445c6  // movb         $0, $-12(%rbp)
	LONG $0x00f245c6  // movb         $0, $-14(%rbp)
	WORD $0x894c; BYTE $0xda  // movq         %r11, %rdx
	WORD $0x2948; BYTE $0xc2  // subq         %rax, %rdx
	LONG $0x02fa8348  // cmpq         $2, %rdx
	LONG $0x0031820f; WORD $0x0000  // jb           LBB0_21, $49(%rip)
	WORD $0xb60f; BYTE $0x30  // movzbl       (%rax), %esi
	LONG $0x0178b60f  // movzbl       $1(%rax), %edi
	LONG $0xf4758840  // movb         %sil, $-12(%rbp)
	LONG $0x02488d48  // leaq         $2(%rax), %rcx
	LONG $0xfec28348  // addq         $-2, %rdx
	WORD $0x894c; BYTE $0xcb  // movq         %r9, %rbx
	WORD $0x8548; BYTE $0xd2  // testq        %rdx, %rdx
	LONG $0x0025840f; WORD $0x0000  // je           LBB0_22, $37(%rip)
LBB0_20:
	WORD $0xb60f; BYTE $0x09  // movzbl       (%rcx), %ecx
	WORD $0x0b88  // movb         %cl, (%rbx)
	LONG $0xf475b60f  // movzbl       $-12(%rbp), %esi
	LONG $0xf24db60f  // movzbl       $-14(%rbp), %ecx
	LONG $0x000015e9; BYTE $0x00  // jmp          LBB0_23, $21(%rip)
LBB0_21:
	WORD $0xf631  // xorl         %esi, %esi
	WORD $0xff31  // xorl         %edi, %edi
	WORD $0x894c; BYTE $0xc3  // movq         %r8, %rbx
	WORD $0x8948; BYTE $0xc1  // movq         %rax, %rcx
	WORD $0x8548; BYTE $0xd2  // testq        %rdx, %rdx
	LONG $0xffdb850f; WORD $0xffff  // jne          LBB0_20, $-37(%rip)
LBB0_22:
	WORD $0xc931  // xorl         %ecx, %ecx
LBB0_23:
	WORD $0xb60f; BYTE $0xc9  // movzbl       %cl, %ecx
	WORD $0xe1c1; BYTE $0x10  // shll         $16, %ecx
	LONG $0xffb60f40  // movzbl       %dil, %edi
	WORD $0xe7c1; BYTE $0x08  // shll         $8, %edi
	LONG $0xd6b60f40  // movzbl       %sil, %edx
	WORD $0xfa09  // orl          %edi, %edx
	WORD $0xd109  // orl          %edx, %ecx
	LONG $0xc0f0e181; WORD $0x00c0  // andl         $12632304, %ecx
	LONG $0x80e0f981; WORD $0x0080  // cmpl         $8421600, %ecx
	LONG $0x0026850f; WORD $0x0000  // jne          LBB0_26, $38(%rip)
	WORD $0xd789  // movl         %edx, %edi
	LONG $0x200fe781; WORD $0x0000  // andl         $8207, %edi
	LONG $0x200dff81; WORD $0x0000  // cmpl         $8205, %edi
	LONG $0x0012840f; WORD $0x0000  // je           LBB0_26, $18(%rip)
	LONG $0x000003b9; BYTE $0x00  // movl         $3, %ecx
	WORD $0xff85  // testl        %edi, %edi
	LONG $0x0026850f; WORD $0x0000  // jne          LBB0_28, $38(%rip)
	LONG $0x90909090; BYTE $0x90  // .p2align 4, 0x90
LBB0_26:
	LONG $0x1ec6f640  // testb        $30, %sil
	LONG $0xff07840f; WORD $0xffff  // je           LBB0_13, $-249(%rip)
	LONG $0xc0e0e281; WORD $0x0000  // andl         $49376, %edx
	LONG $0x000002b9; BYTE $0x00  // movl         $2, %ecx
	LONG $0x80c0fa81; WORD $0x0000  // cmpl         $32960, %edx
	LONG $0xfef0850f; WORD $0xffff  // jne          LBB0_13, $-272(%rip)
LBB0_28:
	WORD $0x0148; BYTE $0xc8  // addq         %rcx, %rax
	WORD $0x394c; BYTE $0xd8  // cmpq         %r11, %rax
	LONG $0xff1f820f; WORD $0xffff  // jb           LBB0_17, $-225(%rip)
LBB0_30:
	WORD $0xc031  // xorl         %eax, %eax
	LONG $0x08c48348  // addq         $8, %rsp
	BYTE $0x5b  // popq         %rbx
	BYTE $0x5d  // popq         %rbp
	BYTE $0xc3  // retq         
	WORD $0x0000  // .p2align 2, 0x00
_MASK_USE_NUMBER:
	LONG $0x00000002  // .long 2

TEXT ·__validate_utf8_fast(SB), NOSPLIT | NOFRAME, $0 - 16
	NO_LOCAL_POINTERS

_entry:
	MOVQ (TLS), R14
	LEAQ -24(SP), R12
	CMPQ R12, 16(R14)
	JBE  _stack_grow

_validate_utf8_fast:
	MOVQ s+0(FP), DI
	CALL ·__validate_utf8_fast_entry+32(SB)  // _validate_utf8_fast
	MOVQ AX, ret+8(FP)
	RET

_stack_grow:
	CALL runtime·morestack_noctxt<>(SB)
	JMP  _entry
