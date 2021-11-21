// This file was generated. DO NOT EDIT.
package arm64

type iclass int

const (
	iclass_none iclass = iota

	// reserved
	iclass_perm_undef // Reserved

	// sme > mortlach_32bit_prod
	iclass_mortlach_f32f32_prod // SME FP32 outer product
	iclass_mortlach_b16f32_prod // SME BF16 outer product
	iclass_mortlach_f16f32_prod // SME FP16 outer product
	iclass_mortlach_i8i32_prod  // SME Int8 outer product

	// sme > mortlach_64bit_prod
	iclass_mortlach_f64f64_prod // SME FP64 outer product
	iclass_mortlach_i16i64_prod // SME Int16 outer product

	// sme > mortlach_ins
	iclass_mortlach_insert_pred // SME move vector to array

	// sme > mortlach_ext
	iclass_mortlach_extract_pred // SME move array to vector

	// sme > mortlach_misc
	iclass_mortlach_zero // SME zero array

	// sme > mortlach_hvadd
	iclass_mortlach_addhv // SME add vector to array

	// sme > mortlach_mem
	iclass_mortlach_contig_load   // SME load array vector (elements)
	iclass_mortlach_contig_store  // SME store array vector (elements)
	iclass_mortlach_ctxt_ldst     // SME save and restore array
	iclass_mortlach_contig_qload  // SME load array vector (quadwords)
	iclass_mortlach_contig_qstore // SME store array vector (quadwords)

	// sve > sve_int_muladd_pred
	iclass_sve_int_mlas_vvv_pred     // SVE integer multiply-accumulate writing addend (predicated)
	iclass_sve_int_mladdsub_vvv_pred // SVE integer multiply-add writing multiplicand (predicated)

	// sve > sve_int_pred_bin
	iclass_sve_int_bin_pred_arit_0 // SVE integer add/subtract vectors (predicated)
	iclass_sve_int_bin_pred_arit_1 // SVE integer min/max/difference (predicated)
	iclass_sve_int_bin_pred_arit_2 // SVE integer multiply vectors (predicated)
	iclass_sve_int_bin_pred_div    // SVE integer divide vectors (predicated)
	iclass_sve_int_bin_pred_log    // SVE bitwise logical operations (predicated)

	// sve > sve_int_pred_red
	iclass_sve_int_reduce_0     // SVE integer add reduction (predicated)
	iclass_sve_int_reduce_1     // SVE integer min/max reduction (predicated)
	iclass_sve_int_movprfx_pred // SVE constructive prefix (predicated)
	iclass_sve_int_reduce_2     // SVE bitwise logical reduction (predicated)

	// sve > sve_int_pred_shift
	iclass_sve_int_bin_pred_shift_0 // SVE bitwise shift by immediate (predicated)
	iclass_sve_int_bin_pred_shift_1 // SVE bitwise shift by vector (predicated)
	iclass_sve_int_bin_pred_shift_2 // SVE bitwise shift by wide elements (predicated)

	// sve > sve_int_pred_un
	iclass_sve_int_un_pred_arit_0 // SVE integer unary operations (predicated)
	iclass_sve_int_un_pred_arit_1 // SVE bitwise unary operations (predicated)

	// sve
	iclass_sve_int_bin_cons_arit_0 // SVE integer add/subtract vectors (unpredicated)

	// sve > sve_int_unpred_logical
	iclass_sve_int_bin_cons_log // SVE bitwise logical operations (unpredicated)
	iclass_sve_int_rotate_imm   // sve_int_rotate_imm
	iclass_sve_int_tern_log     // SVE2 bitwise ternary operations

	// sve > sve_index
	iclass_sve_int_index_ii // SVE index generation (immediate start, immediate increment)
	iclass_sve_int_index_ri // SVE index generation (register start, immediate increment)
	iclass_sve_int_index_ir // SVE index generation (immediate start, register increment)
	iclass_sve_int_index_rr // SVE index generation (register start, register increment)

	// sve > sve_alloca
	iclass_sve_int_arith_vl  // SVE stack frame adjustment
	iclass_sve_int_read_vl_a // SVE stack frame size

	// sve > sve_int_unpred_arit_b
	iclass_sve_int_mul_b   // SVE2 integer multiply vectors (unpredicated)
	iclass_sve_int_sqdmulh // SVE2 signed saturating doubling multiply high (unpredicated)

	// sve > sve_int_unpred_shift
	iclass_sve_int_bin_cons_shift_a // SVE bitwise shift by wide elements (unpredicated)
	iclass_sve_int_bin_cons_shift_b // SVE bitwise shift by immediate (unpredicated)

	// sve
	iclass_sve_int_bin_cons_misc_0_a // SVE address generation

	// sve > sve_int_unpred_misc
	iclass_sve_int_bin_cons_misc_0_b // SVE floating-point trig select coefficient
	iclass_sve_int_bin_cons_misc_0_c // SVE floating-point exponential accelerator
	iclass_sve_int_bin_cons_misc_0_d // SVE constructive prefix (unpredicated)

	// sve > sve_countelt
	iclass_sve_int_countvlv0      // SVE saturating inc/dec vector by element count
	iclass_sve_int_count          // SVE element count
	iclass_sve_int_countvlv1      // SVE inc/dec vector by element count
	iclass_sve_int_pred_pattern_a // SVE inc/dec register by element count
	iclass_sve_int_pred_pattern_b // SVE saturating inc/dec register by element count

	// sve > sve_maskimm
	iclass_sve_int_dup_mask_imm // SVE broadcast bitmask immediate
	iclass_sve_int_log_imm      // SVE bitwise logical with immediate (unpredicated)

	// sve > sve_wideimm_pred
	iclass_sve_int_dup_imm_pred   // SVE copy integer immediate (predicated)
	iclass_sve_int_dup_fpimm_pred // SVE copy floating-point immediate (predicated)

	// sve
	iclass_sve_int_perm_dup_i    // SVE broadcast indexed element
	iclass_sve_int_perm_tbl_3src // SVE table lookup (three sources)
	iclass_sve_int_perm_tbl      // SVE table lookup

	// sve > sve_perm_unpred_d
	iclass_sve_int_perm_dup_r     // SVE broadcast general register
	iclass_sve_int_perm_insrs     // SVE insert general register
	iclass_sve_int_perm_unpk      // SVE unpack vector elements
	iclass_sve_int_perm_insrv     // SVE insert SIMD&FP scalar register
	iclass_sve_int_perm_reverse_z // SVE reverse vector elements

	// sve > sve_perm_predicates
	iclass_sve_int_perm_punpk       // SVE unpack predicate elements
	iclass_sve_int_perm_bin_perm_pp // SVE permute predicate elements
	iclass_sve_int_perm_reverse_p   // SVE reverse predicate elements

	// sve
	iclass_sve_int_perm_bin_perm_zz // SVE permute vector elements

	// sve > sve_perm_pred
	iclass_sve_int_perm_cpy_v    // SVE copy SIMD&FP scalar register to vector (predicated)
	iclass_sve_int_perm_compact  // SVE compress active elements
	iclass_sve_int_perm_last_r   // SVE extract element to general register
	iclass_sve_int_perm_last_v   // SVE extract element to SIMD&FP scalar register
	iclass_sve_int_perm_rev      // SVE reverse within elements
	iclass_sve_int_perm_cpy_r    // SVE copy general register to vector (predicated)
	iclass_sve_int_perm_clast_zz // SVE conditionally broadcast element to vector
	iclass_sve_int_perm_clast_vz // SVE conditionally extract element to SIMD&FP scalar
	iclass_sve_int_perm_splice   // SVE vector splice (destructive)
	iclass_sve_intx_perm_splice  // SVE2 vector splice (constructive)
	iclass_sve_int_perm_revd     // SVE reverse doublewords
	iclass_sve_int_perm_clast_rz // SVE conditionally extract element to general register

	// sve
	iclass_sve_int_sel_vvv // SVE select vector elements (predicated)

	// sve > sve_perm_extract
	iclass_sve_int_perm_extract_i  // SVE extract vector (immediate offset, destructive)
	iclass_sve_intx_perm_extract_i // SVE2 extract vector (immediate offset, constructive)

	// sve > sve_perm_inter_long
	iclass_sve_int_perm_bin_long_perm_zz // SVE permute vector segments

	// sve > sve_cmpvec
	iclass_sve_int_cmp_0 // SVE integer compare vectors
	iclass_sve_int_cmp_1 // SVE integer compare with wide elements

	// sve
	iclass_sve_int_ucmp_vi  // SVE integer compare with unsigned immediate
	iclass_sve_int_scmp_vi  // SVE integer compare with signed immediate
	iclass_sve_int_pred_log // SVE predicate logical operations

	// sve > sve_pred_gen_b
	iclass_sve_int_brkp // SVE propagate break from previous partition

	// sve > sve_pred_gen_c
	iclass_sve_int_brkn  // SVE propagate break to next partition
	iclass_sve_int_break // SVE partition break condition

	// sve > sve_pred_gen_d
	iclass_sve_int_ptest   // SVE predicate test
	iclass_sve_int_pfirst  // SVE predicate first active
	iclass_sve_int_pfalse  // SVE predicate zero
	iclass_sve_int_rdffr   // SVE predicate read from FFR (predicated)
	iclass_sve_int_pnext   // SVE predicate next active
	iclass_sve_int_rdffr_2 // SVE predicate read from FFR (unpredicated)
	iclass_sve_int_ptrue   // SVE predicate initialize

	// sve > sve_cmpgpr
	iclass_sve_int_while_rr // SVE integer compare scalar count and limit
	iclass_sve_int_cterm    // SVE conditionally terminate scalars
	iclass_sve_int_whilenc  // SVE pointer conflict compare

	// sve
	iclass_sve_int_pred_dup // SVE broadcast predicate element

	// sve > sve_wideimm_unpred
	iclass_sve_int_arith_imm0 // SVE integer add/subtract immediate (unpredicated)
	iclass_sve_int_arith_imm1 // SVE integer min/max immediate (unpredicated)
	iclass_sve_int_arith_imm2 // SVE integer multiply immediate (unpredicated)
	iclass_sve_int_dup_imm    // SVE broadcast integer immediate (unpredicated)
	iclass_sve_int_dup_fpimm  // SVE broadcast floating-point immediate (unpredicated)

	// sve > sve_pred_count_a
	iclass_sve_int_pcount_pred // SVE predicate count

	// sve > sve_pred_count_b
	iclass_sve_int_count_v_sat // SVE saturating inc/dec vector by predicate count
	iclass_sve_int_count_r_sat // SVE saturating inc/dec register by predicate count
	iclass_sve_int_count_v     // SVE inc/dec vector by predicate count
	iclass_sve_int_count_r     // SVE inc/dec register by predicate count

	// sve > sve_pred_wrffr
	iclass_sve_int_wrffr  // SVE FFR write from predicate
	iclass_sve_int_setffr // SVE FFR initialise

	// sve > sve_intx_muladd_unpred
	iclass_sve_intx_dot         // SVE integer dot product (unpredicated)
	iclass_sve_intx_qdmlalbt    // SVE2 saturating multiply-add interleaved long
	iclass_sve_intx_cdot        // SVE2 complex integer dot product
	iclass_sve_intx_cmla        // SVE2 complex integer multiply-add
	iclass_sve_intx_mlal_long   // SVE2 integer multiply-add long
	iclass_sve_intx_qdmlal_long // SVE2 saturating  multiply-add long
	iclass_sve_intx_qrdmlah     // SVE2 saturating multiply-add high
	iclass_sve_intx_mixed_dot   // SVE mixed sign dot product

	// sve > sve_intx_predicated
	iclass_sve_intx_accumulate_long_pairs    // SVE2 integer pairwise add and accumulate long
	iclass_sve_intx_pred_arith_unary         // SVE2 integer unary operations (predicated)
	iclass_sve_intx_bin_pred_shift_sat_round // SVE2 saturating/rounding bitwise shift left (predicated)
	iclass_sve_intx_pred_arith_binary        // SVE2 integer halving add/subtract (predicated)
	iclass_sve_intx_arith_binary_pairs       // SVE2 integer pairwise arithmetic
	iclass_sve_intx_pred_arith_binary_sat    // SVE2 saturating add/subtract

	// sve
	iclass_sve_intx_clamp // SVE clamp

	// sve > sve_intx_by_indexed_elem
	iclass_sve_intx_dot_by_indexed_elem        // SVE integer dot product (indexed)
	iclass_sve_intx_mla_by_indexed_elem        // SVE2 integer multiply-add (indexed)
	iclass_sve_intx_qrdmlah_by_indexed_elem    // SVE2 saturating multiply-add high (indexed)
	iclass_sve_intx_mixed_dot_by_indexed_elem  // SVE mixed sign dot product (indexed)
	iclass_sve_intx_qdmla_long_by_indexed_elem // SVE2 saturating multiply-add (indexed)
	iclass_sve_intx_cdot_by_indexed_elem       // SVE2 complex integer dot product (indexed)
	iclass_sve_intx_cmla_by_indexed_elem       // SVE2 complex integer multiply-add (indexed)
	iclass_sve_intx_qrdcmla_by_indexed_elem    // SVE2 complex saturating multiply-add (indexed)
	iclass_sve_intx_mla_long_by_indexed_elem   // SVE2 integer multiply-add long (indexed)
	iclass_sve_intx_mul_long_by_indexed_elem   // SVE2 integer multiply long (indexed)
	iclass_sve_intx_qdmul_long_by_indexed_elem // SVE2 saturating multiply (indexed)
	iclass_sve_intx_qdmulh_by_indexed_elem     // SVE2 saturating  multiply high (indexed)
	iclass_sve_intx_mul_by_indexed_elem        // SVE2 integer multiply (indexed)

	// sve > sve_intx_cons_widening
	iclass_sve_intx_cons_arith_long // SVE2 integer add/subtract long
	iclass_sve_intx_cons_arith_wide // SVE2 integer add/subtract wide
	iclass_sve_intx_cons_mul_long   // SVE2 integer multiply long

	// sve > sve_intx_constructive
	iclass_sve_intx_shift_long // SVE2 bitwise shift left long
	iclass_sve_intx_clong      // SVE2 integer add/subtract interleaved long
	iclass_sve_intx_eorx       // SVE2 bitwise exclusive-or interleaved
	iclass_sve_intx_mmla       // SVE integer matrix multiply accumulate
	iclass_sve_intx_perm_bit   // SVE2 bitwise permute

	// sve > sve_intx_acc
	iclass_sve_intx_cadd         // SVE2 complex integer add
	iclass_sve_intx_aba_long     // SVE2 integer absolute difference and accumulate long
	iclass_sve_intx_adc_long     // SVE2 integer add/subtract long with carry
	iclass_sve_intx_sra          // SVE2 bitwise shift right and accumulate
	iclass_sve_intx_shift_insert // SVE2 bitwise shift and insert
	iclass_sve_intx_aba          // SVE2 integer absolute difference and accumulate

	// sve > sve_intx_narrowing
	iclass_sve_intx_extract_narrow // SVE2 saturating extract narrow
	iclass_sve_intx_shift_narrow   // SVE2 bitwise shift right narrow
	iclass_sve_intx_arith_narrow   // SVE2 integer add/subtract narrow high part

	// sve
	iclass_sve_intx_match // SVE2 character match

	// sve > sve_intx_histseg
	iclass_sve_intx_histseg // SVE2 histogram generation  (segment)

	// sve
	iclass_sve_intx_histcnt // SVE2 histogram generation (vector)

	// sve > sve_intx_crypto
	iclass_sve_crypto_unary        // SVE2 crypto unary operations
	iclass_sve_crypto_binary_dest  // SVE2 crypto destructive binary operations
	iclass_sve_crypto_binary_const // SVE2 crypto constructive binary operations

	// sve
	iclass_sve_fp_fcmla                 // SVE floating-point complex multiply-add (predicated)
	iclass_sve_fp_fcadd                 // SVE floating-point complex add (predicated)
	iclass_sve_fp_fcvt2                 // SVE floating-point convert precision odd elements
	iclass_sve_fp_pairwise              // SVE2 floating-point pairwise operations
	iclass_sve_fp_fma_by_indexed_elem   // SVE floating-point multiply-add (indexed)
	iclass_sve_fp_fcmla_by_indexed_elem // SVE floating-point complex multiply-add (indexed)
	iclass_sve_fp_fmul_by_indexed_elem  // SVE floating-point multiply (indexed)

	// sve > sve_fp_fma_w_by_indexed_elem
	iclass_sve_fp_fdot_by_indexed_elem     // SVE BFloat16 floating-point dot product (indexed)
	iclass_sve_fp_fma_long_by_indexed_elem // SVE floating-point multiply-add long (indexed)

	// sve > sve_fp_fma_w
	iclass_sve_fp_fdot     // SVE BFloat16 floating-point dot product
	iclass_sve_fp_fma_long // SVE floating-point multiply-add long

	// sve
	iclass_sve_fp_fmmla    // SVE floating point matrix multiply accumulate
	iclass_sve_fp_3op_p_pd // SVE floating-point compare vectors
	iclass_sve_fp_3op_u_zd // SVE floating-point arithmetic (unpredicated)

	// sve > sve_fp_pred
	iclass_sve_fp_2op_p_zds   // SVE floating-point arithmetic (predicated)
	iclass_sve_fp_ftmad       // SVE floating-point trig multiply-add coefficient
	iclass_sve_fp_2op_i_p_zds // SVE floating-point arithmetic with immediate (predicated)

	// sve > sve_fp_unary
	iclass_sve_fp_2op_p_zd_a   // SVE floating-point round to integral value
	iclass_sve_fp_2op_p_zd_b_0 // SVE floating-point convert precision
	iclass_sve_fp_2op_p_zd_b_1 // SVE floating-point unary operations
	iclass_sve_fp_2op_p_zd_c   // SVE integer convert to floating-point
	iclass_sve_fp_2op_p_zd_d   // SVE floating-point convert to integer

	// sve
	iclass_sve_fp_fast_red // SVE floating-point recursive reduction

	// sve > sve_fp_unary_unpred
	iclass_sve_fp_2op_u_zd // SVE floating-point reciprocal estimate (unpredicated)

	// sve > sve_fp_cmpzero
	iclass_sve_fp_2op_p_pd // SVE floating-point compare with zero

	// sve > sve_fp_slowreduce
	iclass_sve_fp_2op_p_vd // SVE floating-point serial reduction (predicated)

	// sve > sve_fp_fma
	iclass_sve_fp_3op_p_zds_a // SVE floating-point multiply-accumulate writing addend
	iclass_sve_fp_3op_p_zds_b // SVE floating-point multiply-accumulate writing multiplicand

	// sve > sve_mem32
	iclass_sve_mem_32b_prfm_sv  // SVE 32-bit gather prefetch (scalar plus 32-bit scaled offsets)
	iclass_sve_mem_32b_gld_sv_a // SVE 32-bit gather load halfwords (scalar plus 32-bit scaled offsets)
	iclass_sve_mem_32b_gld_sv_b // SVE 32-bit gather load words (scalar plus 32-bit scaled offsets)
	iclass_sve_mem_32b_pfill    // SVE load predicate register
	iclass_sve_mem_32b_fill     // SVE load vector register
	iclass_sve_mem_prfm_si      // SVE contiguous prefetch (scalar plus immediate)
	iclass_sve_mem_32b_gld_vs   // SVE 32-bit gather load (scalar plus 32-bit unscaled offsets)
	iclass_sve_mem_32b_gldnt_vs // SVE2 32-bit gather non-temporal load (scalar plus 32-bit unscaled offsets)
	iclass_sve_mem_prfm_ss      // SVE contiguous prefetch (scalar plus scalar)
	iclass_sve_mem_32b_prfm_vi  // SVE 32-bit gather prefetch (vector plus immediate)
	iclass_sve_mem_32b_gld_vi   // SVE 32-bit gather load (vector plus immediate)
	iclass_sve_mem_ld_dup       // SVE load and broadcast element

	// sve > sve_memcld
	iclass_sve_mem_cldnt_si // SVE contiguous non-temporal load (scalar plus immediate)
	iclass_sve_mem_cldnt_ss // SVE contiguous non-temporal load (scalar plus scalar)
	iclass_sve_mem_eld_si   // SVE load multiple structures (scalar plus immediate)
	iclass_sve_mem_eld_ss   // SVE load multiple structures (scalar plus scalar)
	iclass_sve_mem_ldqr_si  // SVE load and broadcast quadword (scalar plus immediate)
	iclass_sve_mem_cld_si   // SVE contiguous load (scalar plus immediate)
	iclass_sve_mem_cldnf_si // SVE contiguous non-fault load (scalar plus immediate)
	iclass_sve_mem_ldqr_ss  // SVE load and broadcast quadword (scalar plus scalar)
	iclass_sve_mem_cld_ss   // SVE contiguous load (scalar plus scalar)
	iclass_sve_mem_cldff_ss // SVE contiguous first-fault load (scalar plus scalar)

	// sve > sve_mem64
	iclass_sve_mem_64b_prfm_sv2 // SVE 64-bit gather prefetch (scalar plus 64-bit scaled offsets)
	iclass_sve_mem_64b_prfm_sv  // SVE 64-bit gather prefetch (scalar plus unpacked 32-bit scaled offsets)
	iclass_sve_mem_64b_gld_sv2  // SVE 64-bit gather load (scalar plus 64-bit scaled offsets)
	iclass_sve_mem_64b_gld_sv   // SVE 64-bit gather load (scalar plus 32-bit unpacked scaled offsets)
	iclass_sve_mem_64b_prfm_vi  // SVE 64-bit gather prefetch (vector plus immediate)
	iclass_sve_mem_64b_gldnt_vs // SVE2 64-bit gather non-temporal load (scalar plus unpacked 32-bit unscaled offsets)
	iclass_sve_mem_64b_gld_vi   // SVE 64-bit gather load (vector plus immediate)
	iclass_sve_mem_64b_gld_vs2  // SVE 64-bit gather load (scalar plus 64-bit unscaled offsets)
	iclass_sve_mem_64b_gld_vs   // SVE 64-bit gather load (scalar plus unpacked 32-bit unscaled offsets)

	// sve > sve_memst_cs
	iclass_sve_mem_pspill // SVE store predicate register
	iclass_sve_mem_spill  // SVE store vector register
	iclass_sve_mem_cst_ss // SVE contiguous store (scalar plus scalar)

	// sve > sve_memst_nt
	iclass_sve_mem_sstnt_64b_vs // SVE2 64-bit scatter non-temporal store (vector plus scalar)
	iclass_sve_mem_cstnt_ss     // SVE contiguous non-temporal store (scalar plus scalar)
	iclass_sve_mem_sstnt_32b_vs // SVE2 32-bit scatter non-temporal store (vector plus scalar)
	iclass_sve_mem_est_ss       // SVE store multiple structures (scalar plus scalar)

	// sve > sve_memst_ss
	iclass_sve_mem_sst_vs_a // SVE 64-bit scatter store (scalar plus unpacked 32-bit unscaled offsets)
	iclass_sve_mem_sst_sv_a // SVE 64-bit scatter store (scalar plus unpacked 32-bit scaled offsets)
	iclass_sve_mem_sst_vs_b // SVE 32-bit scatter store (scalar plus 32-bit unscaled offsets)
	iclass_sve_mem_sst_sv_b // SVE 32-bit scatter store (scalar plus 32-bit scaled offsets)

	// sve > sve_memst_ss2
	iclass_sve_mem_sst_vs2  // SVE 64-bit scatter store (scalar plus 64-bit unscaled offsets)
	iclass_sve_mem_sst_sv2  // SVE 64-bit scatter store (scalar plus 64-bit scaled offsets)
	iclass_sve_mem_sst_vi_a // SVE 64-bit scatter store (vector plus immediate)
	iclass_sve_mem_sst_vi_b // SVE 32-bit scatter store (vector plus immediate)

	// sve > sve_memst_si
	iclass_sve_mem_cstnt_si // SVE contiguous non-temporal store (scalar plus immediate)
	iclass_sve_mem_est_si   // SVE store multiple structures (scalar plus immediate)
	iclass_sve_mem_cst_si   // SVE contiguous store (scalar plus immediate)

	// dpimm
	iclass_pcreladdr      // PC-rel. addressing
	iclass_addsub_imm     // Add/subtract (immediate)
	iclass_addsub_immtags // Add/subtract (immediate, with tags)
	iclass_log_imm        // Logical (immediate)
	iclass_movewide       // Move wide (immediate)
	iclass_bitfield       // Bitfield
	iclass_extract        // Extract

	// control
	iclass_condbranch          // Conditional branch (immediate)
	iclass_exception           // Exception generation
	iclass_systeminstrswithreg // System instructions with register argument
	iclass_hints               // Hints
	iclass_barriers            // Barriers
	iclass_pstate              // PSTATE
	iclass_systemresult        // System with result
	iclass_systeminstrs        // System instructions
	iclass_systemmove          // System register move
	iclass_branch_reg          // Unconditional branch (register)
	iclass_branch_imm          // Unconditional branch (immediate)
	iclass_compbranch          // Compare and branch (immediate)
	iclass_testbranch          // Test and branch (immediate)

	// ldst
	iclass_comswappr        // Compare and swap pair
	iclass_asisdlse         // Advanced SIMD load/store multiple structures
	iclass_asisdlsep        // Advanced SIMD load/store multiple structures (post-indexed)
	iclass_asisdlso         // Advanced SIMD load/store single structure
	iclass_asisdlsop        // Advanced SIMD load/store single structure (post-indexed)
	iclass_ldsttags         // Load/store memory tags
	iclass_ldstexclp        // Load/store exclusive pair
	iclass_ldstexclr        // Load/store exclusive register
	iclass_ldstord          // Load/store ordered
	iclass_comswap          // Compare and swap
	iclass_ldapstl_unscaled // LDAPR/STLR (unscaled immediate)
	iclass_loadlit          // Load register (literal)
	iclass_memcms           // Memory Copy and Memory Set
	iclass_ldstnapair_offs  // Load/store no-allocate pair (offset)
	iclass_ldstpair_post    // Load/store register pair (post-indexed)
	iclass_ldstpair_off     // Load/store register pair (offset)
	iclass_ldstpair_pre     // Load/store register pair (pre-indexed)
	iclass_ldst_unscaled    // Load/store register (unscaled immediate)
	iclass_ldst_immpost     // Load/store register (immediate post-indexed)
	iclass_ldst_unpriv      // Load/store register (unprivileged)
	iclass_ldst_immpre      // Load/store register (immediate pre-indexed)
	iclass_memop            // Atomic memory operations
	iclass_ldst_regoff      // Load/store register (register offset)
	iclass_ldst_pac         // Load/store register (pac)
	iclass_ldst_pos         // Load/store register (unsigned immediate)

	// dpreg
	iclass_dp_2src      // Data-processing (2 source)
	iclass_dp_1src      // Data-processing (1 source)
	iclass_log_shift    // Logical (shifted register)
	iclass_addsub_shift // Add/subtract (shifted register)
	iclass_addsub_ext   // Add/subtract (extended register)
	iclass_addsub_carry // Add/subtract (with carry)
	iclass_rmif         // Rotate right into flags
	iclass_setf         // Evaluate into flags
	iclass_condcmp_reg  // Conditional compare (register)
	iclass_condcmp_imm  // Conditional compare (immediate)
	iclass_condsel      // Conditional select
	iclass_dp_3src      // Data-processing (3 source)

	// simd-dp
	iclass_cryptoaes      // Cryptographic AES
	iclass_cryptosha3     // Cryptographic three-register SHA
	iclass_cryptosha2     // Cryptographic two-register SHA
	iclass_asisdone       // Advanced SIMD scalar copy
	iclass_asisdsamefp16  // Advanced SIMD scalar three same FP16
	iclass_asisdmiscfp16  // Advanced SIMD scalar two-register miscellaneous FP16
	iclass_asisdsame2     // Advanced SIMD scalar three same extra
	iclass_asisdmisc      // Advanced SIMD scalar two-register miscellaneous
	iclass_asisdpair      // Advanced SIMD scalar pairwise
	iclass_asisddiff      // Advanced SIMD scalar three different
	iclass_asisdsame      // Advanced SIMD scalar three same
	iclass_asisdshf       // Advanced SIMD scalar shift by immediate
	iclass_asisdelem      // Advanced SIMD scalar x indexed element
	iclass_asimdtbl       // Advanced SIMD table lookup
	iclass_asimdperm      // Advanced SIMD permute
	iclass_asimdext       // Advanced SIMD extract
	iclass_asimdins       // Advanced SIMD copy
	iclass_asimdsamefp16  // Advanced SIMD three same (FP16)
	iclass_asimdmiscfp16  // Advanced SIMD two-register miscellaneous (FP16)
	iclass_asimdsame2     // Advanced SIMD three-register extension
	iclass_asimdmisc      // Advanced SIMD two-register miscellaneous
	iclass_asimdall       // Advanced SIMD across lanes
	iclass_asimddiff      // Advanced SIMD three different
	iclass_asimdsame      // Advanced SIMD three same
	iclass_asimdimm       // Advanced SIMD modified immediate
	iclass_asimdshf       // Advanced SIMD shift by immediate
	iclass_asimdelem      // Advanced SIMD vector x indexed element
	iclass_crypto3_imm2   // Cryptographic three-register, imm2
	iclass_cryptosha512_3 // Cryptographic three-register SHA 512
	iclass_crypto4        // Cryptographic four-register
	iclass_crypto3_imm6   // Cryptographic three-register, imm6
	iclass_cryptosha512_2 // Cryptographic two-register SHA 512
	iclass_float2fix      // Conversion between floating-point and fixed-point
	iclass_float2int      // Conversion between floating-point and integer
	iclass_floatdp1       // Floating-point data-processing (1 source)
	iclass_floatcmp       // Floating-point compare
	iclass_floatimm       // Floating-point immediate
	iclass_floatccmp      // Floating-point conditional compare
	iclass_floatdp2       // Floating-point data-processing (2 source)
	iclass_floatsel       // Floating-point conditional select
	iclass_floatdp3       // Floating-point data-processing (3 source)
)

type encoding int

const (
	encoding_none encoding = iota

	// iclass_perm_undef
	encoding_UDF_only_perm_undef // UDF

	// iclass_mortlach_f32f32_prod
	encoding_fmopa_za_pp_zz_32 // FMOPA (non-widening)
	encoding_fmops_za_pp_zz_32 // FMOPS (non-widening)

	// iclass_mortlach_b16f32_prod
	encoding_bfmopa_za32_pp_zz_ // BFMOPA
	encoding_bfmops_za32_pp_zz_ // BFMOPS

	// iclass_mortlach_f16f32_prod
	encoding_fmopa_za32_pp_zz_16 // FMOPA (widening)
	encoding_fmops_za32_pp_zz_16 // FMOPS (widening)

	// iclass_mortlach_i8i32_prod
	encoding_smopa_za_pp_zz_32  // SMOPA
	encoding_smops_za_pp_zz_32  // SMOPS
	encoding_sumopa_za_pp_zz_32 // SUMOPA
	encoding_sumops_za_pp_zz_32 // SUMOPS
	encoding_usmopa_za_pp_zz_32 // USMOPA
	encoding_usmops_za_pp_zz_32 // USMOPS
	encoding_umopa_za_pp_zz_32  // UMOPA
	encoding_umops_za_pp_zz_32  // UMOPS

	// iclass_mortlach_f64f64_prod
	encoding_fmopa_za_pp_zz_64 // FMOPA (non-widening)
	encoding_fmops_za_pp_zz_64 // FMOPS (non-widening)

	// iclass_mortlach_i16i64_prod
	encoding_smopa_za_pp_zz_64  // SMOPA
	encoding_smops_za_pp_zz_64  // SMOPS
	encoding_sumopa_za_pp_zz_64 // SUMOPA
	encoding_sumops_za_pp_zz_64 // SUMOPS
	encoding_usmopa_za_pp_zz_64 // USMOPA
	encoding_usmops_za_pp_zz_64 // USMOPS
	encoding_umopa_za_pp_zz_64  // UMOPA
	encoding_umops_za_pp_zz_64  // UMOPS

	// iclass_mortlach_insert_pred
	encoding_mova_za_p_rz_b // MOVA (vector to tile)
	encoding_mova_za_p_rz_h // MOVA (vector to tile)
	encoding_mova_za_p_rz_w // MOVA (vector to tile)
	encoding_mova_za_p_rz_d // MOVA (vector to tile)
	encoding_mova_za_p_rz_q // MOVA (vector to tile)

	// iclass_mortlach_extract_pred
	encoding_mova_z_p_rza_b // MOVA (tile to vector)
	encoding_mova_z_p_rza_h // MOVA (tile to vector)
	encoding_mova_z_p_rza_w // MOVA (tile to vector)
	encoding_mova_z_p_rza_d // MOVA (tile to vector)
	encoding_mova_z_p_rza_q // MOVA (tile to vector)

	// iclass_mortlach_zero
	encoding_zero_za_i_ // ZERO

	// iclass_mortlach_addhv
	encoding_addha_za_pp_z_32 // ADDHA
	encoding_addva_za_pp_z_32 // ADDVA
	encoding_addha_za_pp_z_64 // ADDHA
	encoding_addva_za_pp_z_64 // ADDVA

	// iclass_mortlach_contig_load
	encoding_ld1b_za_p_rrr_ // LD1B
	encoding_ld1h_za_p_rrr_ // LD1H
	encoding_ld1w_za_p_rrr_ // LD1W
	encoding_ld1d_za_p_rrr_ // LD1D

	// iclass_mortlach_contig_store
	encoding_st1b_za_p_rrr_ // ST1B
	encoding_st1h_za_p_rrr_ // ST1H
	encoding_st1w_za_p_rrr_ // ST1W
	encoding_st1d_za_p_rrr_ // ST1D

	// iclass_mortlach_ctxt_ldst
	encoding_ldr_za_ri_ // LDR
	encoding_str_za_ri_ // STR

	// iclass_mortlach_contig_qload
	encoding_ld1q_za_p_rrr_ // LD1Q

	// iclass_mortlach_contig_qstore
	encoding_st1q_za_p_rrr_ // ST1Q

	// iclass_sve_int_mlas_vvv_pred
	encoding_mla_z_p_zzz_ // MLA (vectors)
	encoding_mls_z_p_zzz_ // MLS (vectors)

	// iclass_sve_int_mladdsub_vvv_pred
	encoding_mad_z_p_zzz_ // MAD
	encoding_msb_z_p_zzz_ // MSB

	// iclass_sve_int_bin_pred_arit_0
	encoding_add_z_p_zz_  // ADD (vectors, predicated)
	encoding_sub_z_p_zz_  // SUB (vectors, predicated)
	encoding_subr_z_p_zz_ // SUBR (vectors)

	// iclass_sve_int_bin_pred_arit_1
	encoding_smax_z_p_zz_ // SMAX (vectors)
	encoding_umax_z_p_zz_ // UMAX (vectors)
	encoding_smin_z_p_zz_ // SMIN (vectors)
	encoding_umin_z_p_zz_ // UMIN (vectors)
	encoding_sabd_z_p_zz_ // SABD
	encoding_uabd_z_p_zz_ // UABD

	// iclass_sve_int_bin_pred_arit_2
	encoding_mul_z_p_zz_   // MUL (vectors, predicated)
	encoding_smulh_z_p_zz_ // SMULH (predicated)
	encoding_umulh_z_p_zz_ // UMULH (predicated)

	// iclass_sve_int_bin_pred_div
	encoding_sdiv_z_p_zz_  // SDIV
	encoding_udiv_z_p_zz_  // UDIV
	encoding_sdivr_z_p_zz_ // SDIVR
	encoding_udivr_z_p_zz_ // UDIVR

	// iclass_sve_int_bin_pred_log
	encoding_orr_z_p_zz_ // ORR (vectors, predicated)
	encoding_eor_z_p_zz_ // EOR (vectors, predicated)
	encoding_and_z_p_zz_ // AND (vectors, predicated)
	encoding_bic_z_p_zz_ // BIC (vectors, predicated)

	// iclass_sve_int_reduce_0
	encoding_saddv_r_p_z_ // SADDV
	encoding_uaddv_r_p_z_ // UADDV

	// iclass_sve_int_reduce_1
	encoding_smaxv_r_p_z_ // SMAXV
	encoding_umaxv_r_p_z_ // UMAXV
	encoding_sminv_r_p_z_ // SMINV
	encoding_uminv_r_p_z_ // UMINV

	// iclass_sve_int_movprfx_pred
	encoding_movprfx_z_p_z_ // MOVPRFX (predicated)

	// iclass_sve_int_reduce_2
	encoding_orv_r_p_z_  // ORV
	encoding_eorv_r_p_z_ // EORV
	encoding_andv_r_p_z_ // ANDV

	// iclass_sve_int_bin_pred_shift_0
	encoding_asr_z_p_zi_    // ASR (immediate, predicated)
	encoding_lsr_z_p_zi_    // LSR (immediate, predicated)
	encoding_lsl_z_p_zi_    // LSL (immediate, predicated)
	encoding_asrd_z_p_zi_   // ASRD
	encoding_sqshl_z_p_zi_  // SQSHL (immediate)
	encoding_uqshl_z_p_zi_  // UQSHL (immediate)
	encoding_srshr_z_p_zi_  // SRSHR
	encoding_urshr_z_p_zi_  // URSHR
	encoding_sqshlu_z_p_zi_ // SQSHLU

	// iclass_sve_int_bin_pred_shift_1
	encoding_asr_z_p_zz_  // ASR (vectors)
	encoding_lsr_z_p_zz_  // LSR (vectors)
	encoding_lsl_z_p_zz_  // LSL (vectors)
	encoding_asrr_z_p_zz_ // ASRR
	encoding_lsrr_z_p_zz_ // LSRR
	encoding_lslr_z_p_zz_ // LSLR

	// iclass_sve_int_bin_pred_shift_2
	encoding_asr_z_p_zw_ // ASR (wide elements, predicated)
	encoding_lsr_z_p_zw_ // LSR (wide elements, predicated)
	encoding_lsl_z_p_zw_ // LSL (wide elements, predicated)

	// iclass_sve_int_un_pred_arit_0
	encoding_sxtb_z_p_z_ // SXTB, SXTH, SXTW
	encoding_uxtb_z_p_z_ // UXTB, UXTH, UXTW
	encoding_sxth_z_p_z_ // SXTB, SXTH, SXTW
	encoding_uxth_z_p_z_ // UXTB, UXTH, UXTW
	encoding_sxtw_z_p_z_ // SXTB, SXTH, SXTW
	encoding_uxtw_z_p_z_ // UXTB, UXTH, UXTW
	encoding_abs_z_p_z_  // ABS
	encoding_neg_z_p_z_  // NEG

	// iclass_sve_int_un_pred_arit_1
	encoding_cls_z_p_z_  // CLS
	encoding_clz_z_p_z_  // CLZ
	encoding_cnt_z_p_z_  // CNT
	encoding_cnot_z_p_z_ // CNOT
	encoding_fabs_z_p_z_ // FABS
	encoding_fneg_z_p_z_ // FNEG
	encoding_not_z_p_z_  // NOT (vector)

	// iclass_sve_int_bin_cons_arit_0
	encoding_add_z_zz_   // ADD (vectors, unpredicated)
	encoding_sub_z_zz_   // SUB (vectors, unpredicated)
	encoding_sqadd_z_zz_ // SQADD (vectors, unpredicated)
	encoding_uqadd_z_zz_ // UQADD (vectors, unpredicated)
	encoding_sqsub_z_zz_ // SQSUB (vectors, unpredicated)
	encoding_uqsub_z_zz_ // UQSUB (vectors, unpredicated)

	// iclass_sve_int_bin_cons_log
	encoding_and_z_zz_ // AND (vectors, unpredicated)
	encoding_orr_z_zz_ // ORR (vectors, unpredicated)
	encoding_eor_z_zz_ // EOR (vectors, unpredicated)
	encoding_bic_z_zz_ // BIC (vectors, unpredicated)

	// iclass_sve_int_rotate_imm
	encoding_xar_z_zzi_ // XAR

	// iclass_sve_int_tern_log
	encoding_eor3_z_zzz_  // EOR3
	encoding_bsl_z_zzz_   // BSL
	encoding_bcax_z_zzz_  // BCAX
	encoding_bsl1n_z_zzz_ // BSL1N
	encoding_bsl2n_z_zzz_ // BSL2N
	encoding_nbsl_z_zzz_  // NBSL

	// iclass_sve_int_index_ii
	encoding_index_z_ii_ // INDEX (immediates)

	// iclass_sve_int_index_ri
	encoding_index_z_ri_ // INDEX (scalar, immediate)

	// iclass_sve_int_index_ir
	encoding_index_z_ir_ // INDEX (immediate, scalar)

	// iclass_sve_int_index_rr
	encoding_index_z_rr_ // INDEX (scalars)

	// iclass_sve_int_arith_vl
	encoding_addvl_r_ri_ // ADDVL
	encoding_addpl_r_ri_ // ADDPL

	// iclass_sve_int_read_vl_a
	encoding_rdvl_r_i_ // RDVL

	// iclass_sve_int_mul_b
	encoding_mul_z_zz_   // MUL (vectors, unpredicated)
	encoding_smulh_z_zz_ // SMULH (unpredicated)
	encoding_umulh_z_zz_ // UMULH (unpredicated)
	encoding_pmul_z_zz_  // PMUL

	// iclass_sve_int_sqdmulh
	encoding_sqdmulh_z_zz_  // SQDMULH (vectors)
	encoding_sqrdmulh_z_zz_ // SQRDMULH (vectors)

	// iclass_sve_int_bin_cons_shift_a
	encoding_asr_z_zw_ // ASR (wide elements, unpredicated)
	encoding_lsr_z_zw_ // LSR (wide elements, unpredicated)
	encoding_lsl_z_zw_ // LSL (wide elements, unpredicated)

	// iclass_sve_int_bin_cons_shift_b
	encoding_asr_z_zi_ // ASR (immediate, unpredicated)
	encoding_lsr_z_zi_ // LSR (immediate, unpredicated)
	encoding_lsl_z_zi_ // LSL (immediate, unpredicated)

	// iclass_sve_int_bin_cons_misc_0_a
	encoding_adr_z_az_d_s32_scaled   // ADR
	encoding_adr_z_az_d_u32_scaled   // ADR
	encoding_adr_z_az_sd_same_scaled // ADR

	// iclass_sve_int_bin_cons_misc_0_b
	encoding_ftssel_z_zz_ // FTSSEL

	// iclass_sve_int_bin_cons_misc_0_c
	encoding_fexpa_z_z_ // FEXPA

	// iclass_sve_int_bin_cons_misc_0_d
	encoding_movprfx_z_z_ // MOVPRFX (unpredicated)

	// iclass_sve_int_countvlv0
	encoding_sqinch_z_zs_ // SQINCH (vector)
	encoding_uqinch_z_zs_ // UQINCH (vector)
	encoding_sqdech_z_zs_ // SQDECH (vector)
	encoding_uqdech_z_zs_ // UQDECH (vector)
	encoding_sqincw_z_zs_ // SQINCW (vector)
	encoding_uqincw_z_zs_ // UQINCW (vector)
	encoding_sqdecw_z_zs_ // SQDECW (vector)
	encoding_uqdecw_z_zs_ // UQDECW (vector)
	encoding_sqincd_z_zs_ // SQINCD (vector)
	encoding_uqincd_z_zs_ // UQINCD (vector)
	encoding_sqdecd_z_zs_ // SQDECD (vector)
	encoding_uqdecd_z_zs_ // UQDECD (vector)

	// iclass_sve_int_count
	encoding_cntb_r_s_ // CNTB, CNTD, CNTH, CNTW
	encoding_cnth_r_s_ // CNTB, CNTD, CNTH, CNTW
	encoding_cntw_r_s_ // CNTB, CNTD, CNTH, CNTW
	encoding_cntd_r_s_ // CNTB, CNTD, CNTH, CNTW

	// iclass_sve_int_countvlv1
	encoding_inch_z_zs_ // INCD, INCH, INCW (vector)
	encoding_dech_z_zs_ // DECD, DECH, DECW (vector)
	encoding_incw_z_zs_ // INCD, INCH, INCW (vector)
	encoding_decw_z_zs_ // DECD, DECH, DECW (vector)
	encoding_incd_z_zs_ // INCD, INCH, INCW (vector)
	encoding_decd_z_zs_ // DECD, DECH, DECW (vector)

	// iclass_sve_int_pred_pattern_a
	encoding_incb_r_rs_ // INCB, INCD, INCH, INCW (scalar)
	encoding_decb_r_rs_ // DECB, DECD, DECH, DECW (scalar)
	encoding_inch_r_rs_ // INCB, INCD, INCH, INCW (scalar)
	encoding_dech_r_rs_ // DECB, DECD, DECH, DECW (scalar)
	encoding_incw_r_rs_ // INCB, INCD, INCH, INCW (scalar)
	encoding_decw_r_rs_ // DECB, DECD, DECH, DECW (scalar)
	encoding_incd_r_rs_ // INCB, INCD, INCH, INCW (scalar)
	encoding_decd_r_rs_ // DECB, DECD, DECH, DECW (scalar)

	// iclass_sve_int_pred_pattern_b
	encoding_sqincb_r_rs_sx // SQINCB
	encoding_uqincb_r_rs_uw // UQINCB
	encoding_sqdecb_r_rs_sx // SQDECB
	encoding_uqdecb_r_rs_uw // UQDECB
	encoding_sqincb_r_rs_x  // SQINCB
	encoding_uqincb_r_rs_x  // UQINCB
	encoding_sqdecb_r_rs_x  // SQDECB
	encoding_uqdecb_r_rs_x  // UQDECB
	encoding_sqinch_r_rs_sx // SQINCH (scalar)
	encoding_uqinch_r_rs_uw // UQINCH (scalar)
	encoding_sqdech_r_rs_sx // SQDECH (scalar)
	encoding_uqdech_r_rs_uw // UQDECH (scalar)
	encoding_sqinch_r_rs_x  // SQINCH (scalar)
	encoding_uqinch_r_rs_x  // UQINCH (scalar)
	encoding_sqdech_r_rs_x  // SQDECH (scalar)
	encoding_uqdech_r_rs_x  // UQDECH (scalar)
	encoding_sqincw_r_rs_sx // SQINCW (scalar)
	encoding_uqincw_r_rs_uw // UQINCW (scalar)
	encoding_sqdecw_r_rs_sx // SQDECW (scalar)
	encoding_uqdecw_r_rs_uw // UQDECW (scalar)
	encoding_sqincw_r_rs_x  // SQINCW (scalar)
	encoding_uqincw_r_rs_x  // UQINCW (scalar)
	encoding_sqdecw_r_rs_x  // SQDECW (scalar)
	encoding_uqdecw_r_rs_x  // UQDECW (scalar)
	encoding_sqincd_r_rs_sx // SQINCD (scalar)
	encoding_uqincd_r_rs_uw // UQINCD (scalar)
	encoding_sqdecd_r_rs_sx // SQDECD (scalar)
	encoding_uqdecd_r_rs_uw // UQDECD (scalar)
	encoding_sqincd_r_rs_x  // SQINCD (scalar)
	encoding_uqincd_r_rs_x  // UQINCD (scalar)
	encoding_sqdecd_r_rs_x  // SQDECD (scalar)
	encoding_uqdecd_r_rs_x  // UQDECD (scalar)

	// iclass_sve_int_dup_mask_imm
	encoding_dupm_z_i_ // DUPM

	// iclass_sve_int_log_imm
	encoding_orr_z_zi_ // ORR (immediate)
	encoding_eor_z_zi_ // EOR (immediate)
	encoding_and_z_zi_ // AND (immediate)

	// iclass_sve_int_dup_imm_pred
	encoding_cpy_z_o_i_ // CPY (immediate, zeroing)
	encoding_cpy_z_p_i_ // CPY (immediate, merging)

	// iclass_sve_int_dup_fpimm_pred
	encoding_fcpy_z_p_i_ // FCPY

	// iclass_sve_int_perm_dup_i
	encoding_dup_z_zi_ // DUP (indexed)

	// iclass_sve_int_perm_tbl_3src
	encoding_tbl_z_zz_2 // TBL
	encoding_tbx_z_zz_  // TBX

	// iclass_sve_int_perm_tbl
	encoding_tbl_z_zz_1 // TBL

	// iclass_sve_int_perm_dup_r
	encoding_dup_z_r_ // DUP (scalar)

	// iclass_sve_int_perm_insrs
	encoding_insr_z_r_ // INSR (scalar)

	// iclass_sve_int_perm_unpk
	encoding_sunpklo_z_z_ // SUNPKHI, SUNPKLO
	encoding_sunpkhi_z_z_ // SUNPKHI, SUNPKLO
	encoding_uunpklo_z_z_ // UUNPKHI, UUNPKLO
	encoding_uunpkhi_z_z_ // UUNPKHI, UUNPKLO

	// iclass_sve_int_perm_insrv
	encoding_insr_z_v_ // INSR (SIMD&FP scalar)

	// iclass_sve_int_perm_reverse_z
	encoding_rev_z_z_ // REV (vector)

	// iclass_sve_int_perm_punpk
	encoding_punpklo_p_p_ // PUNPKHI, PUNPKLO
	encoding_punpkhi_p_p_ // PUNPKHI, PUNPKLO

	// iclass_sve_int_perm_bin_perm_pp
	encoding_zip1_p_pp_ // ZIP1, ZIP2 (predicates)
	encoding_zip2_p_pp_ // ZIP1, ZIP2 (predicates)
	encoding_uzp1_p_pp_ // UZP1, UZP2 (predicates)
	encoding_uzp2_p_pp_ // UZP1, UZP2 (predicates)
	encoding_trn1_p_pp_ // TRN1, TRN2 (predicates)
	encoding_trn2_p_pp_ // TRN1, TRN2 (predicates)

	// iclass_sve_int_perm_reverse_p
	encoding_rev_p_p_ // REV (predicate)

	// iclass_sve_int_perm_bin_perm_zz
	encoding_zip1_z_zz_ // ZIP1, ZIP2 (vectors)
	encoding_zip2_z_zz_ // ZIP1, ZIP2 (vectors)
	encoding_uzp1_z_zz_ // UZP1, UZP2 (vectors)
	encoding_uzp2_z_zz_ // UZP1, UZP2 (vectors)
	encoding_trn1_z_zz_ // TRN1, TRN2 (vectors)
	encoding_trn2_z_zz_ // TRN1, TRN2 (vectors)

	// iclass_sve_int_perm_cpy_v
	encoding_cpy_z_p_v_ // CPY (SIMD&FP scalar)

	// iclass_sve_int_perm_compact
	encoding_compact_z_p_z_ // COMPACT

	// iclass_sve_int_perm_last_r
	encoding_lasta_r_p_z_ // LASTA (scalar)
	encoding_lastb_r_p_z_ // LASTB (scalar)

	// iclass_sve_int_perm_last_v
	encoding_lasta_v_p_z_ // LASTA (SIMD&FP scalar)
	encoding_lastb_v_p_z_ // LASTB (SIMD&FP scalar)

	// iclass_sve_int_perm_rev
	encoding_revb_z_z_   // REVB, REVH, REVW
	encoding_revh_z_z_   // REVB, REVH, REVW
	encoding_revw_z_z_   // REVB, REVH, REVW
	encoding_rbit_z_p_z_ // RBIT

	// iclass_sve_int_perm_cpy_r
	encoding_cpy_z_p_r_ // CPY (scalar)

	// iclass_sve_int_perm_clast_zz
	encoding_clasta_z_p_zz_ // CLASTA (vectors)
	encoding_clastb_z_p_zz_ // CLASTB (vectors)

	// iclass_sve_int_perm_clast_vz
	encoding_clasta_v_p_z_ // CLASTA (SIMD&FP scalar)
	encoding_clastb_v_p_z_ // CLASTB (SIMD&FP scalar)

	// iclass_sve_int_perm_splice
	encoding_splice_z_p_zz_des // SPLICE

	// iclass_sve_intx_perm_splice
	encoding_splice_z_p_zz_con // SPLICE

	// iclass_sve_int_perm_revd
	encoding_revd_z_p_z_ // REVD

	// iclass_sve_int_perm_clast_rz
	encoding_clasta_r_p_z_ // CLASTA (scalar)
	encoding_clastb_r_p_z_ // CLASTB (scalar)

	// iclass_sve_int_sel_vvv
	encoding_sel_z_p_zz_ // SEL (vectors)

	// iclass_sve_int_perm_extract_i
	encoding_ext_z_zi_des // EXT

	// iclass_sve_intx_perm_extract_i
	encoding_ext_z_zi_con // EXT

	// iclass_sve_int_perm_bin_long_perm_zz
	encoding_zip1_z_zz_q // ZIP1, ZIP2 (vectors)
	encoding_zip2_z_zz_q // ZIP1, ZIP2 (vectors)
	encoding_uzp1_z_zz_q // UZP1, UZP2 (vectors)
	encoding_uzp2_z_zz_q // UZP1, UZP2 (vectors)
	encoding_trn1_z_zz_q // TRN1, TRN2 (vectors)
	encoding_trn2_z_zz_q // TRN1, TRN2 (vectors)

	// iclass_sve_int_cmp_0
	encoding_cmphs_p_p_zz_ // CMP<cc> (vectors)
	encoding_cmphi_p_p_zz_ // CMP<cc> (vectors)
	encoding_cmpeq_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmpne_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmpge_p_p_zz_ // CMP<cc> (vectors)
	encoding_cmpgt_p_p_zz_ // CMP<cc> (vectors)
	encoding_cmpeq_p_p_zz_ // CMP<cc> (vectors)
	encoding_cmpne_p_p_zz_ // CMP<cc> (vectors)

	// iclass_sve_int_cmp_1
	encoding_cmpge_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmpgt_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmplt_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmple_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmphs_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmphi_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmplo_p_p_zw_ // CMP<cc> (wide elements)
	encoding_cmpls_p_p_zw_ // CMP<cc> (wide elements)

	// iclass_sve_int_ucmp_vi
	encoding_cmphs_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmphi_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmplo_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmpls_p_p_zi_ // CMP<cc> (immediate)

	// iclass_sve_int_scmp_vi
	encoding_cmpge_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmpgt_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmplt_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmple_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmpeq_p_p_zi_ // CMP<cc> (immediate)
	encoding_cmpne_p_p_zi_ // CMP<cc> (immediate)

	// iclass_sve_int_pred_log
	encoding_and_p_p_pp_z   // AND (predicates)
	encoding_bic_p_p_pp_z   // BIC (predicates)
	encoding_eor_p_p_pp_z   // EOR (predicates)
	encoding_sel_p_p_pp_    // SEL (predicates)
	encoding_ands_p_p_pp_z  // ANDS
	encoding_bics_p_p_pp_z  // BICS
	encoding_eors_p_p_pp_z  // EORS
	encoding_orr_p_p_pp_z   // ORR (predicates)
	encoding_orn_p_p_pp_z   // ORN (predicates)
	encoding_nor_p_p_pp_z   // NOR
	encoding_nand_p_p_pp_z  // NAND
	encoding_orrs_p_p_pp_z  // ORRS
	encoding_orns_p_p_pp_z  // ORNS
	encoding_nors_p_p_pp_z  // NORS
	encoding_nands_p_p_pp_z // NANDS

	// iclass_sve_int_brkp
	encoding_brkpa_p_p_pp_  // BRKPA
	encoding_brkpb_p_p_pp_  // BRKPB
	encoding_brkpas_p_p_pp_ // BRKPAS
	encoding_brkpbs_p_p_pp_ // BRKPBS

	// iclass_sve_int_brkn
	encoding_brkn_p_p_pp_  // BRKN
	encoding_brkns_p_p_pp_ // BRKNS

	// iclass_sve_int_break
	encoding_brka_p_p_p_   // BRKA
	encoding_brkas_p_p_p_z // BRKAS
	encoding_brkb_p_p_p_   // BRKB
	encoding_brkbs_p_p_p_z // BRKBS

	// iclass_sve_int_ptest
	encoding_ptest_p_p_ // PTEST

	// iclass_sve_int_pfirst
	encoding_pfirst_p_p_p_ // PFIRST

	// iclass_sve_int_pfalse
	encoding_pfalse_p_ // PFALSE

	// iclass_sve_int_rdffr
	encoding_rdffr_p_p_f_  // RDFFR (predicated)
	encoding_rdffrs_p_p_f_ // RDFFRS

	// iclass_sve_int_pnext
	encoding_pnext_p_p_p_ // PNEXT

	// iclass_sve_int_rdffr_2
	encoding_rdffr_p_f_ // RDFFR (unpredicated)

	// iclass_sve_int_ptrue
	encoding_ptrue_p_s_  // PTRUE
	encoding_ptrues_p_s_ // PTRUES

	// iclass_sve_int_while_rr
	encoding_whilege_p_p_rr_ // WHILEGE
	encoding_whilegt_p_p_rr_ // WHILEGT
	encoding_whilelt_p_p_rr_ // WHILELT
	encoding_whilele_p_p_rr_ // WHILELE
	encoding_whilehs_p_p_rr_ // WHILEHS
	encoding_whilehi_p_p_rr_ // WHILEHI
	encoding_whilelo_p_p_rr_ // WHILELO
	encoding_whilels_p_p_rr_ // WHILELS

	// iclass_sve_int_cterm
	encoding_ctermeq_rr_ // CTERMEQ, CTERMNE
	encoding_ctermne_rr_ // CTERMEQ, CTERMNE

	// iclass_sve_int_whilenc
	encoding_whilewr_p_rr_ // WHILEWR
	encoding_whilerw_p_rr_ // WHILERW

	// iclass_sve_int_pred_dup
	encoding_psel_p_ppi_ // PSEL

	// iclass_sve_int_arith_imm0
	encoding_add_z_zi_   // ADD (immediate)
	encoding_sub_z_zi_   // SUB (immediate)
	encoding_subr_z_zi_  // SUBR (immediate)
	encoding_sqadd_z_zi_ // SQADD (immediate)
	encoding_uqadd_z_zi_ // UQADD (immediate)
	encoding_sqsub_z_zi_ // SQSUB (immediate)
	encoding_uqsub_z_zi_ // UQSUB (immediate)

	// iclass_sve_int_arith_imm1
	encoding_smax_z_zi_ // SMAX (immediate)
	encoding_umax_z_zi_ // UMAX (immediate)
	encoding_smin_z_zi_ // SMIN (immediate)
	encoding_umin_z_zi_ // UMIN (immediate)

	// iclass_sve_int_arith_imm2
	encoding_mul_z_zi_ // MUL (immediate)

	// iclass_sve_int_dup_imm
	encoding_dup_z_i_ // DUP (immediate)

	// iclass_sve_int_dup_fpimm
	encoding_fdup_z_i_ // FDUP

	// iclass_sve_int_pcount_pred
	encoding_cntp_r_p_p_ // CNTP

	// iclass_sve_int_count_v_sat
	encoding_sqincp_z_p_z_ // SQINCP (vector)
	encoding_uqincp_z_p_z_ // UQINCP (vector)
	encoding_sqdecp_z_p_z_ // SQDECP (vector)
	encoding_uqdecp_z_p_z_ // UQDECP (vector)

	// iclass_sve_int_count_r_sat
	encoding_sqincp_r_p_r_sx // SQINCP (scalar)
	encoding_sqincp_r_p_r_x  // SQINCP (scalar)
	encoding_uqincp_r_p_r_uw // UQINCP (scalar)
	encoding_uqincp_r_p_r_x  // UQINCP (scalar)
	encoding_sqdecp_r_p_r_sx // SQDECP (scalar)
	encoding_sqdecp_r_p_r_x  // SQDECP (scalar)
	encoding_uqdecp_r_p_r_uw // UQDECP (scalar)
	encoding_uqdecp_r_p_r_x  // UQDECP (scalar)

	// iclass_sve_int_count_v
	encoding_incp_z_p_z_ // INCP (vector)
	encoding_decp_z_p_z_ // DECP (vector)

	// iclass_sve_int_count_r
	encoding_incp_r_p_r_ // INCP (scalar)
	encoding_decp_r_p_r_ // DECP (scalar)

	// iclass_sve_int_wrffr
	encoding_wrffr_f_p_ // WRFFR

	// iclass_sve_int_setffr
	encoding_setffr_f_ // SETFFR

	// iclass_sve_intx_dot
	encoding_sdot_z_zzz_ // SDOT (vectors)
	encoding_udot_z_zzz_ // UDOT (vectors)

	// iclass_sve_intx_qdmlalbt
	encoding_sqdmlalbt_z_zzz_ // SQDMLALBT
	encoding_sqdmlslbt_z_zzz_ // SQDMLSLBT

	// iclass_sve_intx_cdot
	encoding_cdot_z_zzz_ // CDOT (vectors)

	// iclass_sve_intx_cmla
	encoding_cmla_z_zzz_      // CMLA (vectors)
	encoding_sqrdcmlah_z_zzz_ // SQRDCMLAH (vectors)

	// iclass_sve_intx_mlal_long
	encoding_smlalb_z_zzz_ // SMLALB (vectors)
	encoding_smlalt_z_zzz_ // SMLALT (vectors)
	encoding_umlalb_z_zzz_ // UMLALB (vectors)
	encoding_umlalt_z_zzz_ // UMLALT (vectors)
	encoding_smlslb_z_zzz_ // SMLSLB (vectors)
	encoding_smlslt_z_zzz_ // SMLSLT (vectors)
	encoding_umlslb_z_zzz_ // UMLSLB (vectors)
	encoding_umlslt_z_zzz_ // UMLSLT (vectors)

	// iclass_sve_intx_qdmlal_long
	encoding_sqdmlalb_z_zzz_ // SQDMLALB (vectors)
	encoding_sqdmlalt_z_zzz_ // SQDMLALT (vectors)
	encoding_sqdmlslb_z_zzz_ // SQDMLSLB (vectors)
	encoding_sqdmlslt_z_zzz_ // SQDMLSLT (vectors)

	// iclass_sve_intx_qrdmlah
	encoding_sqrdmlah_z_zzz_ // SQRDMLAH (vectors)
	encoding_sqrdmlsh_z_zzz_ // SQRDMLSH (vectors)

	// iclass_sve_intx_mixed_dot
	encoding_usdot_z_zzz_s // USDOT (vectors)

	// iclass_sve_intx_accumulate_long_pairs
	encoding_sadalp_z_p_z_ // SADALP
	encoding_uadalp_z_p_z_ // UADALP

	// iclass_sve_intx_pred_arith_unary
	encoding_urecpe_z_p_z_  // URECPE
	encoding_ursqrte_z_p_z_ // URSQRTE
	encoding_sqabs_z_p_z_   // SQABS
	encoding_sqneg_z_p_z_   // SQNEG

	// iclass_sve_intx_bin_pred_shift_sat_round
	encoding_srshl_z_p_zz_   // SRSHL
	encoding_urshl_z_p_zz_   // URSHL
	encoding_srshlr_z_p_zz_  // SRSHLR
	encoding_urshlr_z_p_zz_  // URSHLR
	encoding_sqshl_z_p_zz_   // SQSHL (vectors)
	encoding_uqshl_z_p_zz_   // UQSHL (vectors)
	encoding_sqrshl_z_p_zz_  // SQRSHL
	encoding_uqrshl_z_p_zz_  // UQRSHL
	encoding_sqshlr_z_p_zz_  // SQSHLR
	encoding_uqshlr_z_p_zz_  // UQSHLR
	encoding_sqrshlr_z_p_zz_ // SQRSHLR
	encoding_uqrshlr_z_p_zz_ // UQRSHLR

	// iclass_sve_intx_pred_arith_binary
	encoding_shadd_z_p_zz_  // SHADD
	encoding_uhadd_z_p_zz_  // UHADD
	encoding_shsub_z_p_zz_  // SHSUB
	encoding_uhsub_z_p_zz_  // UHSUB
	encoding_srhadd_z_p_zz_ // SRHADD
	encoding_urhadd_z_p_zz_ // URHADD
	encoding_shsubr_z_p_zz_ // SHSUBR
	encoding_uhsubr_z_p_zz_ // UHSUBR

	// iclass_sve_intx_arith_binary_pairs
	encoding_addp_z_p_zz_  // ADDP
	encoding_smaxp_z_p_zz_ // SMAXP
	encoding_umaxp_z_p_zz_ // UMAXP
	encoding_sminp_z_p_zz_ // SMINP
	encoding_uminp_z_p_zz_ // UMINP

	// iclass_sve_intx_pred_arith_binary_sat
	encoding_sqadd_z_p_zz_  // SQADD (vectors, predicated)
	encoding_uqadd_z_p_zz_  // UQADD (vectors, predicated)
	encoding_sqsub_z_p_zz_  // SQSUB (vectors, predicated)
	encoding_uqsub_z_p_zz_  // UQSUB (vectors, predicated)
	encoding_suqadd_z_p_zz_ // SUQADD
	encoding_usqadd_z_p_zz_ // USQADD
	encoding_sqsubr_z_p_zz_ // SQSUBR
	encoding_uqsubr_z_p_zz_ // UQSUBR

	// iclass_sve_intx_clamp
	encoding_sclamp_z_zz_ // SCLAMP
	encoding_uclamp_z_zz_ // UCLAMP

	// iclass_sve_intx_dot_by_indexed_elem
	encoding_sdot_z_zzzi_s // SDOT (indexed)
	encoding_udot_z_zzzi_s // UDOT (indexed)
	encoding_sdot_z_zzzi_d // SDOT (indexed)
	encoding_udot_z_zzzi_d // UDOT (indexed)

	// iclass_sve_intx_mla_by_indexed_elem
	encoding_mla_z_zzzi_h // MLA (indexed)
	encoding_mls_z_zzzi_h // MLS (indexed)
	encoding_mla_z_zzzi_s // MLA (indexed)
	encoding_mls_z_zzzi_s // MLS (indexed)
	encoding_mla_z_zzzi_d // MLA (indexed)
	encoding_mls_z_zzzi_d // MLS (indexed)

	// iclass_sve_intx_qrdmlah_by_indexed_elem
	encoding_sqrdmlah_z_zzzi_h // SQRDMLAH (indexed)
	encoding_sqrdmlsh_z_zzzi_h // SQRDMLSH (indexed)
	encoding_sqrdmlah_z_zzzi_s // SQRDMLAH (indexed)
	encoding_sqrdmlsh_z_zzzi_s // SQRDMLSH (indexed)
	encoding_sqrdmlah_z_zzzi_d // SQRDMLAH (indexed)
	encoding_sqrdmlsh_z_zzzi_d // SQRDMLSH (indexed)

	// iclass_sve_intx_mixed_dot_by_indexed_elem
	encoding_usdot_z_zzzi_s // USDOT (indexed)
	encoding_sudot_z_zzzi_s // SUDOT

	// iclass_sve_intx_qdmla_long_by_indexed_elem
	encoding_sqdmlalb_z_zzzi_s // SQDMLALB (indexed)
	encoding_sqdmlalt_z_zzzi_s // SQDMLALT (indexed)
	encoding_sqdmlslb_z_zzzi_s // SQDMLSLB (indexed)
	encoding_sqdmlslt_z_zzzi_s // SQDMLSLT (indexed)
	encoding_sqdmlalb_z_zzzi_d // SQDMLALB (indexed)
	encoding_sqdmlalt_z_zzzi_d // SQDMLALT (indexed)
	encoding_sqdmlslb_z_zzzi_d // SQDMLSLB (indexed)
	encoding_sqdmlslt_z_zzzi_d // SQDMLSLT (indexed)

	// iclass_sve_intx_cdot_by_indexed_elem
	encoding_cdot_z_zzzi_s // CDOT (indexed)
	encoding_cdot_z_zzzi_d // CDOT (indexed)

	// iclass_sve_intx_cmla_by_indexed_elem
	encoding_cmla_z_zzzi_h // CMLA (indexed)
	encoding_cmla_z_zzzi_s // CMLA (indexed)

	// iclass_sve_intx_qrdcmla_by_indexed_elem
	encoding_sqrdcmlah_z_zzzi_h // SQRDCMLAH (indexed)
	encoding_sqrdcmlah_z_zzzi_s // SQRDCMLAH (indexed)

	// iclass_sve_intx_mla_long_by_indexed_elem
	encoding_smlalb_z_zzzi_s // SMLALB (indexed)
	encoding_smlalt_z_zzzi_s // SMLALT (indexed)
	encoding_umlalb_z_zzzi_s // UMLALB (indexed)
	encoding_umlalt_z_zzzi_s // UMLALT (indexed)
	encoding_smlslb_z_zzzi_s // SMLSLB (indexed)
	encoding_smlslt_z_zzzi_s // SMLSLT (indexed)
	encoding_umlslb_z_zzzi_s // UMLSLB (indexed)
	encoding_umlslt_z_zzzi_s // UMLSLT (indexed)
	encoding_smlalb_z_zzzi_d // SMLALB (indexed)
	encoding_smlalt_z_zzzi_d // SMLALT (indexed)
	encoding_umlalb_z_zzzi_d // UMLALB (indexed)
	encoding_umlalt_z_zzzi_d // UMLALT (indexed)
	encoding_smlslb_z_zzzi_d // SMLSLB (indexed)
	encoding_smlslt_z_zzzi_d // SMLSLT (indexed)
	encoding_umlslb_z_zzzi_d // UMLSLB (indexed)
	encoding_umlslt_z_zzzi_d // UMLSLT (indexed)

	// iclass_sve_intx_mul_long_by_indexed_elem
	encoding_smullb_z_zzi_s // SMULLB (indexed)
	encoding_smullt_z_zzi_s // SMULLT (indexed)
	encoding_umullb_z_zzi_s // UMULLB (indexed)
	encoding_umullt_z_zzi_s // UMULLT (indexed)
	encoding_smullb_z_zzi_d // SMULLB (indexed)
	encoding_smullt_z_zzi_d // SMULLT (indexed)
	encoding_umullb_z_zzi_d // UMULLB (indexed)
	encoding_umullt_z_zzi_d // UMULLT (indexed)

	// iclass_sve_intx_qdmul_long_by_indexed_elem
	encoding_sqdmullb_z_zzi_s // SQDMULLB (indexed)
	encoding_sqdmullt_z_zzi_s // SQDMULLT (indexed)
	encoding_sqdmullb_z_zzi_d // SQDMULLB (indexed)
	encoding_sqdmullt_z_zzi_d // SQDMULLT (indexed)

	// iclass_sve_intx_qdmulh_by_indexed_elem
	encoding_sqdmulh_z_zzi_h  // SQDMULH (indexed)
	encoding_sqrdmulh_z_zzi_h // SQRDMULH (indexed)
	encoding_sqdmulh_z_zzi_s  // SQDMULH (indexed)
	encoding_sqrdmulh_z_zzi_s // SQRDMULH (indexed)
	encoding_sqdmulh_z_zzi_d  // SQDMULH (indexed)
	encoding_sqrdmulh_z_zzi_d // SQRDMULH (indexed)

	// iclass_sve_intx_mul_by_indexed_elem
	encoding_mul_z_zzi_h // MUL (indexed)
	encoding_mul_z_zzi_s // MUL (indexed)
	encoding_mul_z_zzi_d // MUL (indexed)

	// iclass_sve_intx_cons_arith_long
	encoding_saddlb_z_zz_ // SADDLB
	encoding_saddlt_z_zz_ // SADDLT
	encoding_uaddlb_z_zz_ // UADDLB
	encoding_uaddlt_z_zz_ // UADDLT
	encoding_ssublb_z_zz_ // SSUBLB
	encoding_ssublt_z_zz_ // SSUBLT
	encoding_usublb_z_zz_ // USUBLB
	encoding_usublt_z_zz_ // USUBLT
	encoding_sabdlb_z_zz_ // SABDLB
	encoding_sabdlt_z_zz_ // SABDLT
	encoding_uabdlb_z_zz_ // UABDLB
	encoding_uabdlt_z_zz_ // UABDLT

	// iclass_sve_intx_cons_arith_wide
	encoding_saddwb_z_zz_ // SADDWB
	encoding_saddwt_z_zz_ // SADDWT
	encoding_uaddwb_z_zz_ // UADDWB
	encoding_uaddwt_z_zz_ // UADDWT
	encoding_ssubwb_z_zz_ // SSUBWB
	encoding_ssubwt_z_zz_ // SSUBWT
	encoding_usubwb_z_zz_ // USUBWB
	encoding_usubwt_z_zz_ // USUBWT

	// iclass_sve_intx_cons_mul_long
	encoding_sqdmullb_z_zz_ // SQDMULLB (vectors)
	encoding_sqdmullt_z_zz_ // SQDMULLT (vectors)
	encoding_pmullb_z_zz_   // PMULLB
	encoding_pmullt_z_zz_   // PMULLT
	encoding_smullb_z_zz_   // SMULLB (vectors)
	encoding_smullt_z_zz_   // SMULLT (vectors)
	encoding_umullb_z_zz_   // UMULLB (vectors)
	encoding_umullt_z_zz_   // UMULLT (vectors)

	// iclass_sve_intx_shift_long
	encoding_sshllb_z_zi_ // SSHLLB
	encoding_sshllt_z_zi_ // SSHLLT
	encoding_ushllb_z_zi_ // USHLLB
	encoding_ushllt_z_zi_ // USHLLT

	// iclass_sve_intx_clong
	encoding_saddlbt_z_zz_ // SADDLBT
	encoding_ssublbt_z_zz_ // SSUBLBT
	encoding_ssubltb_z_zz_ // SSUBLTB

	// iclass_sve_intx_eorx
	encoding_eorbt_z_zz_ // EORBT
	encoding_eortb_z_zz_ // EORTB

	// iclass_sve_intx_mmla
	encoding_smmla_z_zzz_  // SMMLA
	encoding_usmmla_z_zzz_ // USMMLA
	encoding_ummla_z_zzz_  // UMMLA

	// iclass_sve_intx_perm_bit
	encoding_bext_z_zz_ // BEXT
	encoding_bdep_z_zz_ // BDEP
	encoding_bgrp_z_zz_ // BGRP

	// iclass_sve_intx_cadd
	encoding_cadd_z_zz_   // CADD
	encoding_sqcadd_z_zz_ // SQCADD

	// iclass_sve_intx_aba_long
	encoding_sabalb_z_zzz_ // SABALB
	encoding_sabalt_z_zzz_ // SABALT
	encoding_uabalb_z_zzz_ // UABALB
	encoding_uabalt_z_zzz_ // UABALT

	// iclass_sve_intx_adc_long
	encoding_adclb_z_zzz_ // ADCLB
	encoding_adclt_z_zzz_ // ADCLT
	encoding_sbclb_z_zzz_ // SBCLB
	encoding_sbclt_z_zzz_ // SBCLT

	// iclass_sve_intx_sra
	encoding_ssra_z_zi_  // SSRA
	encoding_usra_z_zi_  // USRA
	encoding_srsra_z_zi_ // SRSRA
	encoding_ursra_z_zi_ // URSRA

	// iclass_sve_intx_shift_insert
	encoding_sri_z_zzi_ // SRI
	encoding_sli_z_zzi_ // SLI

	// iclass_sve_intx_aba
	encoding_saba_z_zzz_ // SABA
	encoding_uaba_z_zzz_ // UABA

	// iclass_sve_intx_extract_narrow
	encoding_sqxtnb_z_zz_  // SQXTNB
	encoding_sqxtnt_z_zz_  // SQXTNT
	encoding_uqxtnb_z_zz_  // UQXTNB
	encoding_uqxtnt_z_zz_  // UQXTNT
	encoding_sqxtunb_z_zz_ // SQXTUNB
	encoding_sqxtunt_z_zz_ // SQXTUNT

	// iclass_sve_intx_shift_narrow
	encoding_sqshrunb_z_zi_  // SQSHRUNB
	encoding_sqshrunt_z_zi_  // SQSHRUNT
	encoding_sqrshrunb_z_zi_ // SQRSHRUNB
	encoding_sqrshrunt_z_zi_ // SQRSHRUNT
	encoding_shrnb_z_zi_     // SHRNB
	encoding_shrnt_z_zi_     // SHRNT
	encoding_rshrnb_z_zi_    // RSHRNB
	encoding_rshrnt_z_zi_    // RSHRNT
	encoding_sqshrnb_z_zi_   // SQSHRNB
	encoding_sqshrnt_z_zi_   // SQSHRNT
	encoding_sqrshrnb_z_zi_  // SQRSHRNB
	encoding_sqrshrnt_z_zi_  // SQRSHRNT
	encoding_uqshrnb_z_zi_   // UQSHRNB
	encoding_uqshrnt_z_zi_   // UQSHRNT
	encoding_uqrshrnb_z_zi_  // UQRSHRNB
	encoding_uqrshrnt_z_zi_  // UQRSHRNT

	// iclass_sve_intx_arith_narrow
	encoding_addhnb_z_zz_  // ADDHNB
	encoding_addhnt_z_zz_  // ADDHNT
	encoding_raddhnb_z_zz_ // RADDHNB
	encoding_raddhnt_z_zz_ // RADDHNT
	encoding_subhnb_z_zz_  // SUBHNB
	encoding_subhnt_z_zz_  // SUBHNT
	encoding_rsubhnb_z_zz_ // RSUBHNB
	encoding_rsubhnt_z_zz_ // RSUBHNT

	// iclass_sve_intx_match
	encoding_match_p_p_zz_  // MATCH
	encoding_nmatch_p_p_zz_ // NMATCH

	// iclass_sve_intx_histseg
	encoding_histseg_z_zz_ // HISTSEG

	// iclass_sve_intx_histcnt
	encoding_histcnt_z_p_zz_ // HISTCNT

	// iclass_sve_crypto_unary
	encoding_aesmc_z_z_  // AESMC
	encoding_aesimc_z_z_ // AESIMC

	// iclass_sve_crypto_binary_dest
	encoding_aese_z_zz_ // AESE
	encoding_aesd_z_zz_ // AESD
	encoding_sm4e_z_zz_ // SM4E

	// iclass_sve_crypto_binary_const
	encoding_sm4ekey_z_zz_ // SM4EKEY
	encoding_rax1_z_zz_    // RAX1

	// iclass_sve_fp_fcmla
	encoding_fcmla_z_p_zzz_ // FCMLA (vectors)

	// iclass_sve_fp_fcadd
	encoding_fcadd_z_p_zz_ // FCADD

	// iclass_sve_fp_fcvt2
	encoding_fcvtxnt_z_p_z_d2s  // FCVTXNT
	encoding_fcvtnt_z_p_z_s2h   // FCVTNT
	encoding_fcvtlt_z_p_z_h2s   // FCVTLT
	encoding_bfcvtnt_z_p_z_s2bf // BFCVTNT
	encoding_fcvtnt_z_p_z_d2s   // FCVTNT
	encoding_fcvtlt_z_p_z_s2d   // FCVTLT

	// iclass_sve_fp_pairwise
	encoding_faddp_z_p_zz_   // FADDP
	encoding_fmaxnmp_z_p_zz_ // FMAXNMP
	encoding_fminnmp_z_p_zz_ // FMINNMP
	encoding_fmaxp_z_p_zz_   // FMAXP
	encoding_fminp_z_p_zz_   // FMINP

	// iclass_sve_fp_fma_by_indexed_elem
	encoding_fmla_z_zzzi_h // FMLA (indexed)
	encoding_fmls_z_zzzi_h // FMLS (indexed)
	encoding_fmla_z_zzzi_s // FMLA (indexed)
	encoding_fmls_z_zzzi_s // FMLS (indexed)
	encoding_fmla_z_zzzi_d // FMLA (indexed)
	encoding_fmls_z_zzzi_d // FMLS (indexed)

	// iclass_sve_fp_fcmla_by_indexed_elem
	encoding_fcmla_z_zzzi_h // FCMLA (indexed)
	encoding_fcmla_z_zzzi_s // FCMLA (indexed)

	// iclass_sve_fp_fmul_by_indexed_elem
	encoding_fmul_z_zzi_h // FMUL (indexed)
	encoding_fmul_z_zzi_s // FMUL (indexed)
	encoding_fmul_z_zzi_d // FMUL (indexed)

	// iclass_sve_fp_fdot_by_indexed_elem
	encoding_bfdot_z_zzzi_ // BFDOT (indexed)

	// iclass_sve_fp_fma_long_by_indexed_elem
	encoding_fmlalb_z_zzzi_s // FMLALB (indexed)
	encoding_fmlalt_z_zzzi_s // FMLALT (indexed)
	encoding_fmlslb_z_zzzi_s // FMLSLB (indexed)
	encoding_fmlslt_z_zzzi_s // FMLSLT (indexed)
	encoding_bfmlalb_z_zzzi_ // BFMLALB (indexed)
	encoding_bfmlalt_z_zzzi_ // BFMLALT (indexed)

	// iclass_sve_fp_fdot
	encoding_bfdot_z_zzz_ // BFDOT (vectors)

	// iclass_sve_fp_fma_long
	encoding_fmlalb_z_zzz_  // FMLALB (vectors)
	encoding_fmlalt_z_zzz_  // FMLALT (vectors)
	encoding_fmlslb_z_zzz_  // FMLSLB (vectors)
	encoding_fmlslt_z_zzz_  // FMLSLT (vectors)
	encoding_bfmlalb_z_zzz_ // BFMLALB (vectors)
	encoding_bfmlalt_z_zzz_ // BFMLALT (vectors)

	// iclass_sve_fp_fmmla
	encoding_bfmmla_z_zzz_ // BFMMLA
	encoding_fmmla_z_zzz_s // FMMLA
	encoding_fmmla_z_zzz_d // FMMLA

	// iclass_sve_fp_3op_p_pd
	encoding_fcmge_p_p_zz_ // FCM<cc> (vectors)
	encoding_fcmgt_p_p_zz_ // FCM<cc> (vectors)
	encoding_fcmeq_p_p_zz_ // FCM<cc> (vectors)
	encoding_fcmne_p_p_zz_ // FCM<cc> (vectors)
	encoding_fcmuo_p_p_zz_ // FCM<cc> (vectors)
	encoding_facge_p_p_zz_ // FAC<cc>
	encoding_facgt_p_p_zz_ // FAC<cc>

	// iclass_sve_fp_3op_u_zd
	encoding_fadd_z_zz_    // FADD (vectors, unpredicated)
	encoding_fsub_z_zz_    // FSUB (vectors, unpredicated)
	encoding_fmul_z_zz_    // FMUL (vectors, unpredicated)
	encoding_ftsmul_z_zz_  // FTSMUL
	encoding_frecps_z_zz_  // FRECPS
	encoding_frsqrts_z_zz_ // FRSQRTS

	// iclass_sve_fp_2op_p_zds
	encoding_fadd_z_p_zz_   // FADD (vectors, predicated)
	encoding_fsub_z_p_zz_   // FSUB (vectors, predicated)
	encoding_fmul_z_p_zz_   // FMUL (vectors, predicated)
	encoding_fsubr_z_p_zz_  // FSUBR (vectors)
	encoding_fmaxnm_z_p_zz_ // FMAXNM (vectors)
	encoding_fminnm_z_p_zz_ // FMINNM (vectors)
	encoding_fmax_z_p_zz_   // FMAX (vectors)
	encoding_fmin_z_p_zz_   // FMIN (vectors)
	encoding_fabd_z_p_zz_   // FABD
	encoding_fscale_z_p_zz_ // FSCALE
	encoding_fmulx_z_p_zz_  // FMULX
	encoding_fdivr_z_p_zz_  // FDIVR
	encoding_fdiv_z_p_zz_   // FDIV

	// iclass_sve_fp_ftmad
	encoding_ftmad_z_zzi_ // FTMAD

	// iclass_sve_fp_2op_i_p_zds
	encoding_fadd_z_p_zs_   // FADD (immediate)
	encoding_fsub_z_p_zs_   // FSUB (immediate)
	encoding_fmul_z_p_zs_   // FMUL (immediate)
	encoding_fsubr_z_p_zs_  // FSUBR (immediate)
	encoding_fmaxnm_z_p_zs_ // FMAXNM (immediate)
	encoding_fminnm_z_p_zs_ // FMINNM (immediate)
	encoding_fmax_z_p_zs_   // FMAX (immediate)
	encoding_fmin_z_p_zs_   // FMIN (immediate)

	// iclass_sve_fp_2op_p_zd_a
	encoding_frintn_z_p_z_ // FRINT<r>
	encoding_frintp_z_p_z_ // FRINT<r>
	encoding_frintm_z_p_z_ // FRINT<r>
	encoding_frintz_z_p_z_ // FRINT<r>
	encoding_frinta_z_p_z_ // FRINT<r>
	encoding_frintx_z_p_z_ // FRINT<r>
	encoding_frinti_z_p_z_ // FRINT<r>

	// iclass_sve_fp_2op_p_zd_b_0
	encoding_fcvtx_z_p_z_d2s  // FCVTX
	encoding_fcvt_z_p_z_s2h   // FCVT
	encoding_fcvt_z_p_z_h2s   // FCVT
	encoding_bfcvt_z_p_z_s2bf // BFCVT
	encoding_fcvt_z_p_z_d2h   // FCVT
	encoding_fcvt_z_p_z_h2d   // FCVT
	encoding_fcvt_z_p_z_d2s   // FCVT
	encoding_fcvt_z_p_z_s2d   // FCVT

	// iclass_sve_fp_2op_p_zd_b_1
	encoding_frecpx_z_p_z_ // FRECPX
	encoding_fsqrt_z_p_z_  // FSQRT

	// iclass_sve_fp_2op_p_zd_c
	encoding_scvtf_z_p_z_h2fp16 // SCVTF
	encoding_ucvtf_z_p_z_h2fp16 // UCVTF
	encoding_scvtf_z_p_z_w2fp16 // SCVTF
	encoding_ucvtf_z_p_z_w2fp16 // UCVTF
	encoding_scvtf_z_p_z_x2fp16 // SCVTF
	encoding_ucvtf_z_p_z_x2fp16 // UCVTF
	encoding_scvtf_z_p_z_w2s    // SCVTF
	encoding_ucvtf_z_p_z_w2s    // UCVTF
	encoding_scvtf_z_p_z_w2d    // SCVTF
	encoding_ucvtf_z_p_z_w2d    // UCVTF
	encoding_scvtf_z_p_z_x2s    // SCVTF
	encoding_ucvtf_z_p_z_x2s    // UCVTF
	encoding_scvtf_z_p_z_x2d    // SCVTF
	encoding_ucvtf_z_p_z_x2d    // UCVTF

	// iclass_sve_fp_2op_p_zd_d
	encoding_flogb_z_p_z_        // FLOGB
	encoding_fcvtzs_z_p_z_fp162h // FCVTZS
	encoding_fcvtzu_z_p_z_fp162h // FCVTZU
	encoding_fcvtzs_z_p_z_fp162w // FCVTZS
	encoding_fcvtzu_z_p_z_fp162w // FCVTZU
	encoding_fcvtzs_z_p_z_fp162x // FCVTZS
	encoding_fcvtzu_z_p_z_fp162x // FCVTZU
	encoding_fcvtzs_z_p_z_s2w    // FCVTZS
	encoding_fcvtzu_z_p_z_s2w    // FCVTZU
	encoding_fcvtzs_z_p_z_d2w    // FCVTZS
	encoding_fcvtzu_z_p_z_d2w    // FCVTZU
	encoding_fcvtzs_z_p_z_s2x    // FCVTZS
	encoding_fcvtzu_z_p_z_s2x    // FCVTZU
	encoding_fcvtzs_z_p_z_d2x    // FCVTZS
	encoding_fcvtzu_z_p_z_d2x    // FCVTZU

	// iclass_sve_fp_fast_red
	encoding_faddv_v_p_z_   // FADDV
	encoding_fmaxnmv_v_p_z_ // FMAXNMV
	encoding_fminnmv_v_p_z_ // FMINNMV
	encoding_fmaxv_v_p_z_   // FMAXV
	encoding_fminv_v_p_z_   // FMINV

	// iclass_sve_fp_2op_u_zd
	encoding_frecpe_z_z_  // FRECPE
	encoding_frsqrte_z_z_ // FRSQRTE

	// iclass_sve_fp_2op_p_pd
	encoding_fcmge_p_p_z0_ // FCM<cc> (zero)
	encoding_fcmgt_p_p_z0_ // FCM<cc> (zero)
	encoding_fcmlt_p_p_z0_ // FCM<cc> (zero)
	encoding_fcmle_p_p_z0_ // FCM<cc> (zero)
	encoding_fcmeq_p_p_z0_ // FCM<cc> (zero)
	encoding_fcmne_p_p_z0_ // FCM<cc> (zero)

	// iclass_sve_fp_2op_p_vd
	encoding_fadda_v_p_z_ // FADDA

	// iclass_sve_fp_3op_p_zds_a
	encoding_fmla_z_p_zzz_  // FMLA (vectors)
	encoding_fmls_z_p_zzz_  // FMLS (vectors)
	encoding_fnmla_z_p_zzz_ // FNMLA
	encoding_fnmls_z_p_zzz_ // FNMLS

	// iclass_sve_fp_3op_p_zds_b
	encoding_fmad_z_p_zzz_  // FMAD
	encoding_fmsb_z_p_zzz_  // FMSB
	encoding_fnmad_z_p_zzz_ // FNMAD
	encoding_fnmsb_z_p_zzz_ // FNMSB

	// iclass_sve_mem_32b_prfm_sv
	encoding_prfb_i_p_bz_s_x32_scaled // PRFB (scalar plus vector)
	encoding_prfh_i_p_bz_s_x32_scaled // PRFH (scalar plus vector)
	encoding_prfw_i_p_bz_s_x32_scaled // PRFW (scalar plus vector)
	encoding_prfd_i_p_bz_s_x32_scaled // PRFD (scalar plus vector)

	// iclass_sve_mem_32b_gld_sv_a
	encoding_ld1sh_z_p_bz_s_x32_scaled   // LD1SH (scalar plus vector)
	encoding_ldff1sh_z_p_bz_s_x32_scaled // LDFF1SH (scalar plus vector)
	encoding_ld1h_z_p_bz_s_x32_scaled    // LD1H (scalar plus vector)
	encoding_ldff1h_z_p_bz_s_x32_scaled  // LDFF1H (scalar plus vector)

	// iclass_sve_mem_32b_gld_sv_b
	encoding_ld1w_z_p_bz_s_x32_scaled   // LD1W (scalar plus vector)
	encoding_ldff1w_z_p_bz_s_x32_scaled // LDFF1W (scalar plus vector)

	// iclass_sve_mem_32b_pfill
	encoding_ldr_p_bi_ // LDR (predicate)

	// iclass_sve_mem_32b_fill
	encoding_ldr_z_bi_ // LDR (vector)

	// iclass_sve_mem_prfm_si
	encoding_prfb_i_p_bi_s // PRFB (scalar plus immediate)
	encoding_prfh_i_p_bi_s // PRFH (scalar plus immediate)
	encoding_prfw_i_p_bi_s // PRFW (scalar plus immediate)
	encoding_prfd_i_p_bi_s // PRFD (scalar plus immediate)

	// iclass_sve_mem_32b_gld_vs
	encoding_ld1sb_z_p_bz_s_x32_unscaled   // LD1SB (scalar plus vector)
	encoding_ldff1sb_z_p_bz_s_x32_unscaled // LDFF1SB (scalar plus vector)
	encoding_ld1b_z_p_bz_s_x32_unscaled    // LD1B (scalar plus vector)
	encoding_ldff1b_z_p_bz_s_x32_unscaled  // LDFF1B (scalar plus vector)
	encoding_ld1sh_z_p_bz_s_x32_unscaled   // LD1SH (scalar plus vector)
	encoding_ldff1sh_z_p_bz_s_x32_unscaled // LDFF1SH (scalar plus vector)
	encoding_ld1h_z_p_bz_s_x32_unscaled    // LD1H (scalar plus vector)
	encoding_ldff1h_z_p_bz_s_x32_unscaled  // LDFF1H (scalar plus vector)
	encoding_ld1w_z_p_bz_s_x32_unscaled    // LD1W (scalar plus vector)
	encoding_ldff1w_z_p_bz_s_x32_unscaled  // LDFF1W (scalar plus vector)

	// iclass_sve_mem_32b_gldnt_vs
	encoding_ldnt1sb_z_p_ar_s_x32_unscaled // LDNT1SB
	encoding_ldnt1b_z_p_ar_s_x32_unscaled  // LDNT1B (vector plus scalar)
	encoding_ldnt1sh_z_p_ar_s_x32_unscaled // LDNT1SH
	encoding_ldnt1h_z_p_ar_s_x32_unscaled  // LDNT1H (vector plus scalar)
	encoding_ldnt1w_z_p_ar_s_x32_unscaled  // LDNT1W (vector plus scalar)

	// iclass_sve_mem_prfm_ss
	encoding_prfb_i_p_br_s // PRFB (scalar plus scalar)
	encoding_prfh_i_p_br_s // PRFH (scalar plus scalar)
	encoding_prfw_i_p_br_s // PRFW (scalar plus scalar)
	encoding_prfd_i_p_br_s // PRFD (scalar plus scalar)

	// iclass_sve_mem_32b_prfm_vi
	encoding_prfb_i_p_ai_s // PRFB (vector plus immediate)
	encoding_prfh_i_p_ai_s // PRFH (vector plus immediate)
	encoding_prfw_i_p_ai_s // PRFW (vector plus immediate)
	encoding_prfd_i_p_ai_s // PRFD (vector plus immediate)

	// iclass_sve_mem_32b_gld_vi
	encoding_ld1sb_z_p_ai_s   // LD1SB (vector plus immediate)
	encoding_ldff1sb_z_p_ai_s // LDFF1SB (vector plus immediate)
	encoding_ld1b_z_p_ai_s    // LD1B (vector plus immediate)
	encoding_ldff1b_z_p_ai_s  // LDFF1B (vector plus immediate)
	encoding_ld1sh_z_p_ai_s   // LD1SH (vector plus immediate)
	encoding_ldff1sh_z_p_ai_s // LDFF1SH (vector plus immediate)
	encoding_ld1h_z_p_ai_s    // LD1H (vector plus immediate)
	encoding_ldff1h_z_p_ai_s  // LDFF1H (vector plus immediate)
	encoding_ld1w_z_p_ai_s    // LD1W (vector plus immediate)
	encoding_ldff1w_z_p_ai_s  // LDFF1W (vector plus immediate)

	// iclass_sve_mem_ld_dup
	encoding_ld1rb_z_p_bi_u8   // LD1RB
	encoding_ld1rb_z_p_bi_u16  // LD1RB
	encoding_ld1rb_z_p_bi_u32  // LD1RB
	encoding_ld1rb_z_p_bi_u64  // LD1RB
	encoding_ld1rsw_z_p_bi_s64 // LD1RSW
	encoding_ld1rh_z_p_bi_u16  // LD1RH
	encoding_ld1rh_z_p_bi_u32  // LD1RH
	encoding_ld1rh_z_p_bi_u64  // LD1RH
	encoding_ld1rsh_z_p_bi_s64 // LD1RSH
	encoding_ld1rsh_z_p_bi_s32 // LD1RSH
	encoding_ld1rw_z_p_bi_u32  // LD1RW
	encoding_ld1rw_z_p_bi_u64  // LD1RW
	encoding_ld1rsb_z_p_bi_s64 // LD1RSB
	encoding_ld1rsb_z_p_bi_s32 // LD1RSB
	encoding_ld1rsb_z_p_bi_s16 // LD1RSB
	encoding_ld1rd_z_p_bi_u64  // LD1RD

	// iclass_sve_mem_cldnt_si
	encoding_ldnt1b_z_p_bi_contiguous // LDNT1B (scalar plus immediate)
	encoding_ldnt1h_z_p_bi_contiguous // LDNT1H (scalar plus immediate)
	encoding_ldnt1w_z_p_bi_contiguous // LDNT1W (scalar plus immediate)
	encoding_ldnt1d_z_p_bi_contiguous // LDNT1D (scalar plus immediate)

	// iclass_sve_mem_cldnt_ss
	encoding_ldnt1b_z_p_br_contiguous // LDNT1B (scalar plus scalar)
	encoding_ldnt1h_z_p_br_contiguous // LDNT1H (scalar plus scalar)
	encoding_ldnt1w_z_p_br_contiguous // LDNT1W (scalar plus scalar)
	encoding_ldnt1d_z_p_br_contiguous // LDNT1D (scalar plus scalar)

	// iclass_sve_mem_eld_si
	encoding_ld2b_z_p_bi_contiguous // LD2B (scalar plus immediate)
	encoding_ld3b_z_p_bi_contiguous // LD3B (scalar plus immediate)
	encoding_ld4b_z_p_bi_contiguous // LD4B (scalar plus immediate)
	encoding_ld2h_z_p_bi_contiguous // LD2H (scalar plus immediate)
	encoding_ld3h_z_p_bi_contiguous // LD3H (scalar plus immediate)
	encoding_ld4h_z_p_bi_contiguous // LD4H (scalar plus immediate)
	encoding_ld2w_z_p_bi_contiguous // LD2W (scalar plus immediate)
	encoding_ld3w_z_p_bi_contiguous // LD3W (scalar plus immediate)
	encoding_ld4w_z_p_bi_contiguous // LD4W (scalar plus immediate)
	encoding_ld2d_z_p_bi_contiguous // LD2D (scalar plus immediate)
	encoding_ld3d_z_p_bi_contiguous // LD3D (scalar plus immediate)
	encoding_ld4d_z_p_bi_contiguous // LD4D (scalar plus immediate)

	// iclass_sve_mem_eld_ss
	encoding_ld2b_z_p_br_contiguous // LD2B (scalar plus scalar)
	encoding_ld3b_z_p_br_contiguous // LD3B (scalar plus scalar)
	encoding_ld4b_z_p_br_contiguous // LD4B (scalar plus scalar)
	encoding_ld2h_z_p_br_contiguous // LD2H (scalar plus scalar)
	encoding_ld3h_z_p_br_contiguous // LD3H (scalar plus scalar)
	encoding_ld4h_z_p_br_contiguous // LD4H (scalar plus scalar)
	encoding_ld2w_z_p_br_contiguous // LD2W (scalar plus scalar)
	encoding_ld3w_z_p_br_contiguous // LD3W (scalar plus scalar)
	encoding_ld4w_z_p_br_contiguous // LD4W (scalar plus scalar)
	encoding_ld2d_z_p_br_contiguous // LD2D (scalar plus scalar)
	encoding_ld3d_z_p_br_contiguous // LD3D (scalar plus scalar)
	encoding_ld4d_z_p_br_contiguous // LD4D (scalar plus scalar)

	// iclass_sve_mem_ldqr_si
	encoding_ld1rqb_z_p_bi_u8  // LD1RQB (scalar plus immediate)
	encoding_ld1rob_z_p_bi_u8  // LD1ROB (scalar plus immediate)
	encoding_ld1rqh_z_p_bi_u16 // LD1RQH (scalar plus immediate)
	encoding_ld1roh_z_p_bi_u16 // LD1ROH (scalar plus immediate)
	encoding_ld1rqw_z_p_bi_u32 // LD1RQW (scalar plus immediate)
	encoding_ld1row_z_p_bi_u32 // LD1ROW (scalar plus immediate)
	encoding_ld1rqd_z_p_bi_u64 // LD1RQD (scalar plus immediate)
	encoding_ld1rod_z_p_bi_u64 // LD1ROD (scalar plus immediate)

	// iclass_sve_mem_cld_si
	encoding_ld1b_z_p_bi_u8   // LD1B (scalar plus immediate)
	encoding_ld1b_z_p_bi_u16  // LD1B (scalar plus immediate)
	encoding_ld1b_z_p_bi_u32  // LD1B (scalar plus immediate)
	encoding_ld1b_z_p_bi_u64  // LD1B (scalar plus immediate)
	encoding_ld1sw_z_p_bi_s64 // LD1SW (scalar plus immediate)
	encoding_ld1h_z_p_bi_u16  // LD1H (scalar plus immediate)
	encoding_ld1h_z_p_bi_u32  // LD1H (scalar plus immediate)
	encoding_ld1h_z_p_bi_u64  // LD1H (scalar plus immediate)
	encoding_ld1sh_z_p_bi_s64 // LD1SH (scalar plus immediate)
	encoding_ld1sh_z_p_bi_s32 // LD1SH (scalar plus immediate)
	encoding_ld1w_z_p_bi_u32  // LD1W (scalar plus immediate)
	encoding_ld1w_z_p_bi_u64  // LD1W (scalar plus immediate)
	encoding_ld1sb_z_p_bi_s64 // LD1SB (scalar plus immediate)
	encoding_ld1sb_z_p_bi_s32 // LD1SB (scalar plus immediate)
	encoding_ld1sb_z_p_bi_s16 // LD1SB (scalar plus immediate)
	encoding_ld1d_z_p_bi_u64  // LD1D (scalar plus immediate)

	// iclass_sve_mem_cldnf_si
	encoding_ldnf1b_z_p_bi_u8   // LDNF1B
	encoding_ldnf1b_z_p_bi_u16  // LDNF1B
	encoding_ldnf1b_z_p_bi_u32  // LDNF1B
	encoding_ldnf1b_z_p_bi_u64  // LDNF1B
	encoding_ldnf1sw_z_p_bi_s64 // LDNF1SW
	encoding_ldnf1h_z_p_bi_u16  // LDNF1H
	encoding_ldnf1h_z_p_bi_u32  // LDNF1H
	encoding_ldnf1h_z_p_bi_u64  // LDNF1H
	encoding_ldnf1sh_z_p_bi_s64 // LDNF1SH
	encoding_ldnf1sh_z_p_bi_s32 // LDNF1SH
	encoding_ldnf1w_z_p_bi_u32  // LDNF1W
	encoding_ldnf1w_z_p_bi_u64  // LDNF1W
	encoding_ldnf1sb_z_p_bi_s64 // LDNF1SB
	encoding_ldnf1sb_z_p_bi_s32 // LDNF1SB
	encoding_ldnf1sb_z_p_bi_s16 // LDNF1SB
	encoding_ldnf1d_z_p_bi_u64  // LDNF1D

	// iclass_sve_mem_ldqr_ss
	encoding_ld1rqb_z_p_br_contiguous // LD1RQB (scalar plus scalar)
	encoding_ld1rob_z_p_br_contiguous // LD1ROB (scalar plus scalar)
	encoding_ld1rqh_z_p_br_contiguous // LD1RQH (scalar plus scalar)
	encoding_ld1roh_z_p_br_contiguous // LD1ROH (scalar plus scalar)
	encoding_ld1rqw_z_p_br_contiguous // LD1RQW (scalar plus scalar)
	encoding_ld1row_z_p_br_contiguous // LD1ROW (scalar plus scalar)
	encoding_ld1rqd_z_p_br_contiguous // LD1RQD (scalar plus scalar)
	encoding_ld1rod_z_p_br_contiguous // LD1ROD (scalar plus scalar)

	// iclass_sve_mem_cld_ss
	encoding_ld1b_z_p_br_u8   // LD1B (scalar plus scalar)
	encoding_ld1b_z_p_br_u16  // LD1B (scalar plus scalar)
	encoding_ld1b_z_p_br_u32  // LD1B (scalar plus scalar)
	encoding_ld1b_z_p_br_u64  // LD1B (scalar plus scalar)
	encoding_ld1sw_z_p_br_s64 // LD1SW (scalar plus scalar)
	encoding_ld1h_z_p_br_u16  // LD1H (scalar plus scalar)
	encoding_ld1h_z_p_br_u32  // LD1H (scalar plus scalar)
	encoding_ld1h_z_p_br_u64  // LD1H (scalar plus scalar)
	encoding_ld1sh_z_p_br_s64 // LD1SH (scalar plus scalar)
	encoding_ld1sh_z_p_br_s32 // LD1SH (scalar plus scalar)
	encoding_ld1w_z_p_br_u32  // LD1W (scalar plus scalar)
	encoding_ld1w_z_p_br_u64  // LD1W (scalar plus scalar)
	encoding_ld1sb_z_p_br_s64 // LD1SB (scalar plus scalar)
	encoding_ld1sb_z_p_br_s32 // LD1SB (scalar plus scalar)
	encoding_ld1sb_z_p_br_s16 // LD1SB (scalar plus scalar)
	encoding_ld1d_z_p_br_u64  // LD1D (scalar plus scalar)

	// iclass_sve_mem_cldff_ss
	encoding_ldff1b_z_p_br_u8   // LDFF1B (scalar plus scalar)
	encoding_ldff1b_z_p_br_u16  // LDFF1B (scalar plus scalar)
	encoding_ldff1b_z_p_br_u32  // LDFF1B (scalar plus scalar)
	encoding_ldff1b_z_p_br_u64  // LDFF1B (scalar plus scalar)
	encoding_ldff1sw_z_p_br_s64 // LDFF1SW (scalar plus scalar)
	encoding_ldff1h_z_p_br_u16  // LDFF1H (scalar plus scalar)
	encoding_ldff1h_z_p_br_u32  // LDFF1H (scalar plus scalar)
	encoding_ldff1h_z_p_br_u64  // LDFF1H (scalar plus scalar)
	encoding_ldff1sh_z_p_br_s64 // LDFF1SH (scalar plus scalar)
	encoding_ldff1sh_z_p_br_s32 // LDFF1SH (scalar plus scalar)
	encoding_ldff1w_z_p_br_u32  // LDFF1W (scalar plus scalar)
	encoding_ldff1w_z_p_br_u64  // LDFF1W (scalar plus scalar)
	encoding_ldff1sb_z_p_br_s64 // LDFF1SB (scalar plus scalar)
	encoding_ldff1sb_z_p_br_s32 // LDFF1SB (scalar plus scalar)
	encoding_ldff1sb_z_p_br_s16 // LDFF1SB (scalar plus scalar)
	encoding_ldff1d_z_p_br_u64  // LDFF1D (scalar plus scalar)

	// iclass_sve_mem_64b_prfm_sv2
	encoding_prfb_i_p_bz_d_64_scaled // PRFB (scalar plus vector)
	encoding_prfh_i_p_bz_d_64_scaled // PRFH (scalar plus vector)
	encoding_prfw_i_p_bz_d_64_scaled // PRFW (scalar plus vector)
	encoding_prfd_i_p_bz_d_64_scaled // PRFD (scalar plus vector)

	// iclass_sve_mem_64b_prfm_sv
	encoding_prfb_i_p_bz_d_x32_scaled // PRFB (scalar plus vector)
	encoding_prfh_i_p_bz_d_x32_scaled // PRFH (scalar plus vector)
	encoding_prfw_i_p_bz_d_x32_scaled // PRFW (scalar plus vector)
	encoding_prfd_i_p_bz_d_x32_scaled // PRFD (scalar plus vector)

	// iclass_sve_mem_64b_gld_sv2
	encoding_ld1sh_z_p_bz_d_64_scaled   // LD1SH (scalar plus vector)
	encoding_ldff1sh_z_p_bz_d_64_scaled // LDFF1SH (scalar plus vector)
	encoding_ld1h_z_p_bz_d_64_scaled    // LD1H (scalar plus vector)
	encoding_ldff1h_z_p_bz_d_64_scaled  // LDFF1H (scalar plus vector)
	encoding_ld1sw_z_p_bz_d_64_scaled   // LD1SW (scalar plus vector)
	encoding_ldff1sw_z_p_bz_d_64_scaled // LDFF1SW (scalar plus vector)
	encoding_ld1w_z_p_bz_d_64_scaled    // LD1W (scalar plus vector)
	encoding_ldff1w_z_p_bz_d_64_scaled  // LDFF1W (scalar plus vector)
	encoding_ld1d_z_p_bz_d_64_scaled    // LD1D (scalar plus vector)
	encoding_ldff1d_z_p_bz_d_64_scaled  // LDFF1D (scalar plus vector)

	// iclass_sve_mem_64b_gld_sv
	encoding_ld1sh_z_p_bz_d_x32_scaled   // LD1SH (scalar plus vector)
	encoding_ldff1sh_z_p_bz_d_x32_scaled // LDFF1SH (scalar plus vector)
	encoding_ld1h_z_p_bz_d_x32_scaled    // LD1H (scalar plus vector)
	encoding_ldff1h_z_p_bz_d_x32_scaled  // LDFF1H (scalar plus vector)
	encoding_ld1sw_z_p_bz_d_x32_scaled   // LD1SW (scalar plus vector)
	encoding_ldff1sw_z_p_bz_d_x32_scaled // LDFF1SW (scalar plus vector)
	encoding_ld1w_z_p_bz_d_x32_scaled    // LD1W (scalar plus vector)
	encoding_ldff1w_z_p_bz_d_x32_scaled  // LDFF1W (scalar plus vector)
	encoding_ld1d_z_p_bz_d_x32_scaled    // LD1D (scalar plus vector)
	encoding_ldff1d_z_p_bz_d_x32_scaled  // LDFF1D (scalar plus vector)

	// iclass_sve_mem_64b_prfm_vi
	encoding_prfb_i_p_ai_d // PRFB (vector plus immediate)
	encoding_prfh_i_p_ai_d // PRFH (vector plus immediate)
	encoding_prfw_i_p_ai_d // PRFW (vector plus immediate)
	encoding_prfd_i_p_ai_d // PRFD (vector plus immediate)

	// iclass_sve_mem_64b_gldnt_vs
	encoding_ldnt1sb_z_p_ar_d_64_unscaled // LDNT1SB
	encoding_ldnt1b_z_p_ar_d_64_unscaled  // LDNT1B (vector plus scalar)
	encoding_ldnt1sh_z_p_ar_d_64_unscaled // LDNT1SH
	encoding_ldnt1h_z_p_ar_d_64_unscaled  // LDNT1H (vector plus scalar)
	encoding_ldnt1sw_z_p_ar_d_64_unscaled // LDNT1SW
	encoding_ldnt1w_z_p_ar_d_64_unscaled  // LDNT1W (vector plus scalar)
	encoding_ldnt1d_z_p_ar_d_64_unscaled  // LDNT1D (vector plus scalar)

	// iclass_sve_mem_64b_gld_vi
	encoding_ld1sb_z_p_ai_d   // LD1SB (vector plus immediate)
	encoding_ldff1sb_z_p_ai_d // LDFF1SB (vector plus immediate)
	encoding_ld1b_z_p_ai_d    // LD1B (vector plus immediate)
	encoding_ldff1b_z_p_ai_d  // LDFF1B (vector plus immediate)
	encoding_ld1sh_z_p_ai_d   // LD1SH (vector plus immediate)
	encoding_ldff1sh_z_p_ai_d // LDFF1SH (vector plus immediate)
	encoding_ld1h_z_p_ai_d    // LD1H (vector plus immediate)
	encoding_ldff1h_z_p_ai_d  // LDFF1H (vector plus immediate)
	encoding_ld1sw_z_p_ai_d   // LD1SW (vector plus immediate)
	encoding_ldff1sw_z_p_ai_d // LDFF1SW (vector plus immediate)
	encoding_ld1w_z_p_ai_d    // LD1W (vector plus immediate)
	encoding_ldff1w_z_p_ai_d  // LDFF1W (vector plus immediate)
	encoding_ld1d_z_p_ai_d    // LD1D (vector plus immediate)
	encoding_ldff1d_z_p_ai_d  // LDFF1D (vector plus immediate)

	// iclass_sve_mem_64b_gld_vs2
	encoding_ld1sb_z_p_bz_d_64_unscaled   // LD1SB (scalar plus vector)
	encoding_ldff1sb_z_p_bz_d_64_unscaled // LDFF1SB (scalar plus vector)
	encoding_ld1b_z_p_bz_d_64_unscaled    // LD1B (scalar plus vector)
	encoding_ldff1b_z_p_bz_d_64_unscaled  // LDFF1B (scalar plus vector)
	encoding_ld1sh_z_p_bz_d_64_unscaled   // LD1SH (scalar plus vector)
	encoding_ldff1sh_z_p_bz_d_64_unscaled // LDFF1SH (scalar plus vector)
	encoding_ld1h_z_p_bz_d_64_unscaled    // LD1H (scalar plus vector)
	encoding_ldff1h_z_p_bz_d_64_unscaled  // LDFF1H (scalar plus vector)
	encoding_ld1sw_z_p_bz_d_64_unscaled   // LD1SW (scalar plus vector)
	encoding_ldff1sw_z_p_bz_d_64_unscaled // LDFF1SW (scalar plus vector)
	encoding_ld1w_z_p_bz_d_64_unscaled    // LD1W (scalar plus vector)
	encoding_ldff1w_z_p_bz_d_64_unscaled  // LDFF1W (scalar plus vector)
	encoding_ld1d_z_p_bz_d_64_unscaled    // LD1D (scalar plus vector)
	encoding_ldff1d_z_p_bz_d_64_unscaled  // LDFF1D (scalar plus vector)

	// iclass_sve_mem_64b_gld_vs
	encoding_ld1sb_z_p_bz_d_x32_unscaled   // LD1SB (scalar plus vector)
	encoding_ldff1sb_z_p_bz_d_x32_unscaled // LDFF1SB (scalar plus vector)
	encoding_ld1b_z_p_bz_d_x32_unscaled    // LD1B (scalar plus vector)
	encoding_ldff1b_z_p_bz_d_x32_unscaled  // LDFF1B (scalar plus vector)
	encoding_ld1sh_z_p_bz_d_x32_unscaled   // LD1SH (scalar plus vector)
	encoding_ldff1sh_z_p_bz_d_x32_unscaled // LDFF1SH (scalar plus vector)
	encoding_ld1h_z_p_bz_d_x32_unscaled    // LD1H (scalar plus vector)
	encoding_ldff1h_z_p_bz_d_x32_unscaled  // LDFF1H (scalar plus vector)
	encoding_ld1sw_z_p_bz_d_x32_unscaled   // LD1SW (scalar plus vector)
	encoding_ldff1sw_z_p_bz_d_x32_unscaled // LDFF1SW (scalar plus vector)
	encoding_ld1w_z_p_bz_d_x32_unscaled    // LD1W (scalar plus vector)
	encoding_ldff1w_z_p_bz_d_x32_unscaled  // LDFF1W (scalar plus vector)
	encoding_ld1d_z_p_bz_d_x32_unscaled    // LD1D (scalar plus vector)
	encoding_ldff1d_z_p_bz_d_x32_unscaled  // LDFF1D (scalar plus vector)

	// iclass_sve_mem_pspill
	encoding_str_p_bi_ // STR (predicate)

	// iclass_sve_mem_spill
	encoding_str_z_bi_ // STR (vector)

	// iclass_sve_mem_cst_ss
	encoding_st1b_z_p_br_ // ST1B (scalar plus scalar)
	encoding_st1h_z_p_br_ // ST1H (scalar plus scalar)
	encoding_st1w_z_p_br_ // ST1W (scalar plus scalar)
	encoding_st1d_z_p_br_ // ST1D (scalar plus scalar)

	// iclass_sve_mem_sstnt_64b_vs
	encoding_stnt1b_z_p_ar_d_64_unscaled // STNT1B (vector plus scalar)
	encoding_stnt1h_z_p_ar_d_64_unscaled // STNT1H (vector plus scalar)
	encoding_stnt1w_z_p_ar_d_64_unscaled // STNT1W (vector plus scalar)
	encoding_stnt1d_z_p_ar_d_64_unscaled // STNT1D (vector plus scalar)

	// iclass_sve_mem_cstnt_ss
	encoding_stnt1b_z_p_br_contiguous // STNT1B (scalar plus scalar)
	encoding_stnt1h_z_p_br_contiguous // STNT1H (scalar plus scalar)
	encoding_stnt1w_z_p_br_contiguous // STNT1W (scalar plus scalar)
	encoding_stnt1d_z_p_br_contiguous // STNT1D (scalar plus scalar)

	// iclass_sve_mem_sstnt_32b_vs
	encoding_stnt1b_z_p_ar_s_x32_unscaled // STNT1B (vector plus scalar)
	encoding_stnt1h_z_p_ar_s_x32_unscaled // STNT1H (vector plus scalar)
	encoding_stnt1w_z_p_ar_s_x32_unscaled // STNT1W (vector plus scalar)

	// iclass_sve_mem_est_ss
	encoding_st2b_z_p_br_contiguous // ST2B (scalar plus scalar)
	encoding_st3b_z_p_br_contiguous // ST3B (scalar plus scalar)
	encoding_st4b_z_p_br_contiguous // ST4B (scalar plus scalar)
	encoding_st2h_z_p_br_contiguous // ST2H (scalar plus scalar)
	encoding_st3h_z_p_br_contiguous // ST3H (scalar plus scalar)
	encoding_st4h_z_p_br_contiguous // ST4H (scalar plus scalar)
	encoding_st2w_z_p_br_contiguous // ST2W (scalar plus scalar)
	encoding_st3w_z_p_br_contiguous // ST3W (scalar plus scalar)
	encoding_st4w_z_p_br_contiguous // ST4W (scalar plus scalar)
	encoding_st2d_z_p_br_contiguous // ST2D (scalar plus scalar)
	encoding_st3d_z_p_br_contiguous // ST3D (scalar plus scalar)
	encoding_st4d_z_p_br_contiguous // ST4D (scalar plus scalar)

	// iclass_sve_mem_sst_vs_a
	encoding_st1b_z_p_bz_d_x32_unscaled // ST1B (scalar plus vector)
	encoding_st1h_z_p_bz_d_x32_unscaled // ST1H (scalar plus vector)
	encoding_st1w_z_p_bz_d_x32_unscaled // ST1W (scalar plus vector)
	encoding_st1d_z_p_bz_d_x32_unscaled // ST1D (scalar plus vector)

	// iclass_sve_mem_sst_sv_a
	encoding_st1h_z_p_bz_d_x32_scaled // ST1H (scalar plus vector)
	encoding_st1w_z_p_bz_d_x32_scaled // ST1W (scalar plus vector)
	encoding_st1d_z_p_bz_d_x32_scaled // ST1D (scalar plus vector)

	// iclass_sve_mem_sst_vs_b
	encoding_st1b_z_p_bz_s_x32_unscaled // ST1B (scalar plus vector)
	encoding_st1h_z_p_bz_s_x32_unscaled // ST1H (scalar plus vector)
	encoding_st1w_z_p_bz_s_x32_unscaled // ST1W (scalar plus vector)

	// iclass_sve_mem_sst_sv_b
	encoding_st1h_z_p_bz_s_x32_scaled // ST1H (scalar plus vector)
	encoding_st1w_z_p_bz_s_x32_scaled // ST1W (scalar plus vector)

	// iclass_sve_mem_sst_vs2
	encoding_st1b_z_p_bz_d_64_unscaled // ST1B (scalar plus vector)
	encoding_st1h_z_p_bz_d_64_unscaled // ST1H (scalar plus vector)
	encoding_st1w_z_p_bz_d_64_unscaled // ST1W (scalar plus vector)
	encoding_st1d_z_p_bz_d_64_unscaled // ST1D (scalar plus vector)

	// iclass_sve_mem_sst_sv2
	encoding_st1h_z_p_bz_d_64_scaled // ST1H (scalar plus vector)
	encoding_st1w_z_p_bz_d_64_scaled // ST1W (scalar plus vector)
	encoding_st1d_z_p_bz_d_64_scaled // ST1D (scalar plus vector)

	// iclass_sve_mem_sst_vi_a
	encoding_st1b_z_p_ai_d // ST1B (vector plus immediate)
	encoding_st1h_z_p_ai_d // ST1H (vector plus immediate)
	encoding_st1w_z_p_ai_d // ST1W (vector plus immediate)
	encoding_st1d_z_p_ai_d // ST1D (vector plus immediate)

	// iclass_sve_mem_sst_vi_b
	encoding_st1b_z_p_ai_s // ST1B (vector plus immediate)
	encoding_st1h_z_p_ai_s // ST1H (vector plus immediate)
	encoding_st1w_z_p_ai_s // ST1W (vector plus immediate)

	// iclass_sve_mem_cstnt_si
	encoding_stnt1b_z_p_bi_contiguous // STNT1B (scalar plus immediate)
	encoding_stnt1h_z_p_bi_contiguous // STNT1H (scalar plus immediate)
	encoding_stnt1w_z_p_bi_contiguous // STNT1W (scalar plus immediate)
	encoding_stnt1d_z_p_bi_contiguous // STNT1D (scalar plus immediate)

	// iclass_sve_mem_est_si
	encoding_st2b_z_p_bi_contiguous // ST2B (scalar plus immediate)
	encoding_st3b_z_p_bi_contiguous // ST3B (scalar plus immediate)
	encoding_st4b_z_p_bi_contiguous // ST4B (scalar plus immediate)
	encoding_st2h_z_p_bi_contiguous // ST2H (scalar plus immediate)
	encoding_st3h_z_p_bi_contiguous // ST3H (scalar plus immediate)
	encoding_st4h_z_p_bi_contiguous // ST4H (scalar plus immediate)
	encoding_st2w_z_p_bi_contiguous // ST2W (scalar plus immediate)
	encoding_st3w_z_p_bi_contiguous // ST3W (scalar plus immediate)
	encoding_st4w_z_p_bi_contiguous // ST4W (scalar plus immediate)
	encoding_st2d_z_p_bi_contiguous // ST2D (scalar plus immediate)
	encoding_st3d_z_p_bi_contiguous // ST3D (scalar plus immediate)
	encoding_st4d_z_p_bi_contiguous // ST4D (scalar plus immediate)

	// iclass_sve_mem_cst_si
	encoding_st1b_z_p_bi_ // ST1B (scalar plus immediate)
	encoding_st1h_z_p_bi_ // ST1H (scalar plus immediate)
	encoding_st1w_z_p_bi_ // ST1W (scalar plus immediate)
	encoding_st1d_z_p_bi_ // ST1D (scalar plus immediate)

	// iclass_pcreladdr
	encoding_ADR_only_pcreladdr  // ADR
	encoding_ADRP_only_pcreladdr // ADRP

	// iclass_addsub_imm
	encoding_ADD_32_addsub_imm   // ADD (immediate)
	encoding_ADDS_32S_addsub_imm // ADDS (immediate)
	encoding_SUB_32_addsub_imm   // SUB (immediate)
	encoding_SUBS_32S_addsub_imm // SUBS (immediate)
	encoding_ADD_64_addsub_imm   // ADD (immediate)
	encoding_ADDS_64S_addsub_imm // ADDS (immediate)
	encoding_SUB_64_addsub_imm   // SUB (immediate)
	encoding_SUBS_64S_addsub_imm // SUBS (immediate)

	// iclass_addsub_immtags
	encoding_ADDG_64_addsub_immtags // ADDG
	encoding_SUBG_64_addsub_immtags // SUBG

	// iclass_log_imm
	encoding_AND_32_log_imm   // AND (immediate)
	encoding_ORR_32_log_imm   // ORR (immediate)
	encoding_EOR_32_log_imm   // EOR (immediate)
	encoding_ANDS_32S_log_imm // ANDS (immediate)
	encoding_AND_64_log_imm   // AND (immediate)
	encoding_ORR_64_log_imm   // ORR (immediate)
	encoding_EOR_64_log_imm   // EOR (immediate)
	encoding_ANDS_64S_log_imm // ANDS (immediate)

	// iclass_movewide
	encoding_MOVN_32_movewide // MOVN
	encoding_MOVZ_32_movewide // MOVZ
	encoding_MOVK_32_movewide // MOVK
	encoding_MOVN_64_movewide // MOVN
	encoding_MOVZ_64_movewide // MOVZ
	encoding_MOVK_64_movewide // MOVK

	// iclass_bitfield
	encoding_SBFM_32M_bitfield // SBFM
	encoding_BFM_32M_bitfield  // BFM
	encoding_UBFM_32M_bitfield // UBFM
	encoding_SBFM_64M_bitfield // SBFM
	encoding_BFM_64M_bitfield  // BFM
	encoding_UBFM_64M_bitfield // UBFM

	// iclass_extract
	encoding_EXTR_32_extract // EXTR
	encoding_EXTR_64_extract // EXTR

	// iclass_condbranch
	encoding_B_only_condbranch  // B.cond
	encoding_BC_only_condbranch // BC.cond

	// iclass_exception
	encoding_SVC_EX_exception     // SVC
	encoding_HVC_EX_exception     // HVC
	encoding_SMC_EX_exception     // SMC
	encoding_BRK_EX_exception     // BRK
	encoding_HLT_EX_exception     // HLT
	encoding_TCANCEL_EX_exception // TCANCEL
	encoding_DCPS1_DC_exception   // DCPS1
	encoding_DCPS2_DC_exception   // DCPS2
	encoding_DCPS3_DC_exception   // DCPS3

	// iclass_systeminstrswithreg
	encoding_WFET_only_systeminstrswithreg // WFET
	encoding_WFIT_only_systeminstrswithreg // WFIT

	// iclass_hints
	encoding_HINT_HM_hints      // HINT
	encoding_NOP_HI_hints       // NOP
	encoding_YIELD_HI_hints     // YIELD
	encoding_WFE_HI_hints       // WFE
	encoding_WFI_HI_hints       // WFI
	encoding_SEV_HI_hints       // SEV
	encoding_SEVL_HI_hints      // SEVL
	encoding_DGH_HI_hints       // DGH
	encoding_XPACLRI_HI_hints   // XPACD, XPACI, XPACLRI
	encoding_PACIA1716_HI_hints // PACIA, PACIA1716, PACIASP, PACIAZ, PACIZA
	encoding_PACIB1716_HI_hints // PACIB, PACIB1716, PACIBSP, PACIBZ, PACIZB
	encoding_AUTIA1716_HI_hints // AUTIA, AUTIA1716, AUTIASP, AUTIAZ, AUTIZA
	encoding_AUTIB1716_HI_hints // AUTIB, AUTIB1716, AUTIBSP, AUTIBZ, AUTIZB
	encoding_ESB_HI_hints       // ESB
	encoding_PSB_HC_hints       // PSB CSYNC
	encoding_TSB_HC_hints       // TSB CSYNC
	encoding_CSDB_HI_hints      // CSDB
	encoding_PACIAZ_HI_hints    // PACIA, PACIA1716, PACIASP, PACIAZ, PACIZA
	encoding_PACIASP_HI_hints   // PACIA, PACIA1716, PACIASP, PACIAZ, PACIZA
	encoding_PACIBZ_HI_hints    // PACIB, PACIB1716, PACIBSP, PACIBZ, PACIZB
	encoding_PACIBSP_HI_hints   // PACIB, PACIB1716, PACIBSP, PACIBZ, PACIZB
	encoding_AUTIAZ_HI_hints    // AUTIA, AUTIA1716, AUTIASP, AUTIAZ, AUTIZA
	encoding_AUTIASP_HI_hints   // AUTIA, AUTIA1716, AUTIASP, AUTIAZ, AUTIZA
	encoding_AUTIBZ_HI_hints    // AUTIB, AUTIB1716, AUTIBSP, AUTIBZ, AUTIZB
	encoding_AUTIBSP_HI_hints   // AUTIB, AUTIB1716, AUTIBSP, AUTIBZ, AUTIZB
	encoding_BTI_HB_hints       // BTI

	// iclass_barriers
	encoding_CLREX_BN_barriers     // CLREX
	encoding_DSB_BO_barriers       // DSB
	encoding_DMB_BO_barriers       // DMB
	encoding_ISB_BI_barriers       // ISB
	encoding_SB_only_barriers      // SB
	encoding_DSB_BOn_barriers      // DSB
	encoding_TCOMMIT_only_barriers // TCOMMIT

	// iclass_pstate
	encoding_MSR_SI_pstate   // MSR (immediate)
	encoding_CFINV_M_pstate  // CFINV
	encoding_XAFLAG_M_pstate // XAFLAG
	encoding_AXFLAG_M_pstate // AXFLAG

	// iclass_systemresult
	encoding_TSTART_BR_systemresult // TSTART
	encoding_TTEST_BR_systemresult  // TTEST

	// iclass_systeminstrs
	encoding_SYS_CR_systeminstrs  // SYS
	encoding_SYSL_RC_systeminstrs // SYSL

	// iclass_systemmove
	encoding_MSR_SR_systemmove // MSR (register)
	encoding_MRS_RS_systemmove // MRS

	// iclass_branch_reg
	encoding_BR_64_branch_reg      // BR
	encoding_BRAAZ_64_branch_reg   // BRAA, BRAAZ, BRAB, BRABZ
	encoding_BRABZ_64_branch_reg   // BRAA, BRAAZ, BRAB, BRABZ
	encoding_BLR_64_branch_reg     // BLR
	encoding_BLRAAZ_64_branch_reg  // BLRAA, BLRAAZ, BLRAB, BLRABZ
	encoding_BLRABZ_64_branch_reg  // BLRAA, BLRAAZ, BLRAB, BLRABZ
	encoding_RET_64R_branch_reg    // RET
	encoding_RETAA_64E_branch_reg  // RETAA, RETAB
	encoding_RETAB_64E_branch_reg  // RETAA, RETAB
	encoding_ERET_64E_branch_reg   // ERET
	encoding_ERETAA_64E_branch_reg // ERETAA, ERETAB
	encoding_ERETAB_64E_branch_reg // ERETAA, ERETAB
	encoding_DRPS_64E_branch_reg   // DRPS
	encoding_BRAA_64P_branch_reg   // BRAA, BRAAZ, BRAB, BRABZ
	encoding_BRAB_64P_branch_reg   // BRAA, BRAAZ, BRAB, BRABZ
	encoding_BLRAA_64P_branch_reg  // BLRAA, BLRAAZ, BLRAB, BLRABZ
	encoding_BLRAB_64P_branch_reg  // BLRAA, BLRAAZ, BLRAB, BLRABZ

	// iclass_branch_imm
	encoding_B_only_branch_imm  // B
	encoding_BL_only_branch_imm // BL

	// iclass_compbranch
	encoding_CBZ_32_compbranch  // CBZ
	encoding_CBNZ_32_compbranch // CBNZ
	encoding_CBZ_64_compbranch  // CBZ
	encoding_CBNZ_64_compbranch // CBNZ

	// iclass_testbranch
	encoding_TBZ_only_testbranch  // TBZ
	encoding_TBNZ_only_testbranch // TBNZ

	// iclass_comswappr
	encoding_CASP_CP32_comswappr   // CASP, CASPA, CASPAL, CASPL
	encoding_CASPL_CP32_comswappr  // CASP, CASPA, CASPAL, CASPL
	encoding_CASPA_CP32_comswappr  // CASP, CASPA, CASPAL, CASPL
	encoding_CASPAL_CP32_comswappr // CASP, CASPA, CASPAL, CASPL
	encoding_CASP_CP64_comswappr   // CASP, CASPA, CASPAL, CASPL
	encoding_CASPL_CP64_comswappr  // CASP, CASPA, CASPAL, CASPL
	encoding_CASPA_CP64_comswappr  // CASP, CASPA, CASPAL, CASPL
	encoding_CASPAL_CP64_comswappr // CASP, CASPA, CASPAL, CASPL

	// iclass_asisdlse
	encoding_ST4_asisdlse_R4    // ST4 (multiple structures)
	encoding_ST1_asisdlse_R4_4v // ST1 (multiple structures)
	encoding_ST3_asisdlse_R3    // ST3 (multiple structures)
	encoding_ST1_asisdlse_R3_3v // ST1 (multiple structures)
	encoding_ST1_asisdlse_R1_1v // ST1 (multiple structures)
	encoding_ST2_asisdlse_R2    // ST2 (multiple structures)
	encoding_ST1_asisdlse_R2_2v // ST1 (multiple structures)
	encoding_LD4_asisdlse_R4    // LD4 (multiple structures)
	encoding_LD1_asisdlse_R4_4v // LD1 (multiple structures)
	encoding_LD3_asisdlse_R3    // LD3 (multiple structures)
	encoding_LD1_asisdlse_R3_3v // LD1 (multiple structures)
	encoding_LD1_asisdlse_R1_1v // LD1 (multiple structures)
	encoding_LD2_asisdlse_R2    // LD2 (multiple structures)
	encoding_LD1_asisdlse_R2_2v // LD1 (multiple structures)

	// iclass_asisdlsep
	encoding_ST4_asisdlsep_R4_r  // ST4 (multiple structures)
	encoding_ST1_asisdlsep_R4_r4 // ST1 (multiple structures)
	encoding_ST3_asisdlsep_R3_r  // ST3 (multiple structures)
	encoding_ST1_asisdlsep_R3_r3 // ST1 (multiple structures)
	encoding_ST1_asisdlsep_R1_r1 // ST1 (multiple structures)
	encoding_ST2_asisdlsep_R2_r  // ST2 (multiple structures)
	encoding_ST1_asisdlsep_R2_r2 // ST1 (multiple structures)
	encoding_ST4_asisdlsep_I4_i  // ST4 (multiple structures)
	encoding_ST1_asisdlsep_I4_i4 // ST1 (multiple structures)
	encoding_ST3_asisdlsep_I3_i  // ST3 (multiple structures)
	encoding_ST1_asisdlsep_I3_i3 // ST1 (multiple structures)
	encoding_ST1_asisdlsep_I1_i1 // ST1 (multiple structures)
	encoding_ST2_asisdlsep_I2_i  // ST2 (multiple structures)
	encoding_ST1_asisdlsep_I2_i2 // ST1 (multiple structures)
	encoding_LD4_asisdlsep_R4_r  // LD4 (multiple structures)
	encoding_LD1_asisdlsep_R4_r4 // LD1 (multiple structures)
	encoding_LD3_asisdlsep_R3_r  // LD3 (multiple structures)
	encoding_LD1_asisdlsep_R3_r3 // LD1 (multiple structures)
	encoding_LD1_asisdlsep_R1_r1 // LD1 (multiple structures)
	encoding_LD2_asisdlsep_R2_r  // LD2 (multiple structures)
	encoding_LD1_asisdlsep_R2_r2 // LD1 (multiple structures)
	encoding_LD4_asisdlsep_I4_i  // LD4 (multiple structures)
	encoding_LD1_asisdlsep_I4_i4 // LD1 (multiple structures)
	encoding_LD3_asisdlsep_I3_i  // LD3 (multiple structures)
	encoding_LD1_asisdlsep_I3_i3 // LD1 (multiple structures)
	encoding_LD1_asisdlsep_I1_i1 // LD1 (multiple structures)
	encoding_LD2_asisdlsep_I2_i  // LD2 (multiple structures)
	encoding_LD1_asisdlsep_I2_i2 // LD1 (multiple structures)

	// iclass_asisdlso
	encoding_ST1_asisdlso_B1_1b // ST1 (single structure)
	encoding_ST3_asisdlso_B3_3b // ST3 (single structure)
	encoding_ST1_asisdlso_H1_1h // ST1 (single structure)
	encoding_ST3_asisdlso_H3_3h // ST3 (single structure)
	encoding_ST1_asisdlso_S1_1s // ST1 (single structure)
	encoding_ST1_asisdlso_D1_1d // ST1 (single structure)
	encoding_ST3_asisdlso_S3_3s // ST3 (single structure)
	encoding_ST3_asisdlso_D3_3d // ST3 (single structure)
	encoding_ST2_asisdlso_B2_2b // ST2 (single structure)
	encoding_ST4_asisdlso_B4_4b // ST4 (single structure)
	encoding_ST2_asisdlso_H2_2h // ST2 (single structure)
	encoding_ST4_asisdlso_H4_4h // ST4 (single structure)
	encoding_ST2_asisdlso_S2_2s // ST2 (single structure)
	encoding_ST2_asisdlso_D2_2d // ST2 (single structure)
	encoding_ST4_asisdlso_S4_4s // ST4 (single structure)
	encoding_ST4_asisdlso_D4_4d // ST4 (single structure)
	encoding_LD1_asisdlso_B1_1b // LD1 (single structure)
	encoding_LD3_asisdlso_B3_3b // LD3 (single structure)
	encoding_LD1_asisdlso_H1_1h // LD1 (single structure)
	encoding_LD3_asisdlso_H3_3h // LD3 (single structure)
	encoding_LD1_asisdlso_S1_1s // LD1 (single structure)
	encoding_LD1_asisdlso_D1_1d // LD1 (single structure)
	encoding_LD3_asisdlso_S3_3s // LD3 (single structure)
	encoding_LD3_asisdlso_D3_3d // LD3 (single structure)
	encoding_LD1R_asisdlso_R1   // LD1R
	encoding_LD3R_asisdlso_R3   // LD3R
	encoding_LD2_asisdlso_B2_2b // LD2 (single structure)
	encoding_LD4_asisdlso_B4_4b // LD4 (single structure)
	encoding_LD2_asisdlso_H2_2h // LD2 (single structure)
	encoding_LD4_asisdlso_H4_4h // LD4 (single structure)
	encoding_LD2_asisdlso_S2_2s // LD2 (single structure)
	encoding_LD2_asisdlso_D2_2d // LD2 (single structure)
	encoding_LD4_asisdlso_S4_4s // LD4 (single structure)
	encoding_LD4_asisdlso_D4_4d // LD4 (single structure)
	encoding_LD2R_asisdlso_R2   // LD2R
	encoding_LD4R_asisdlso_R4   // LD4R

	// iclass_asisdlsop
	encoding_ST1_asisdlsop_BX1_r1b // ST1 (single structure)
	encoding_ST3_asisdlsop_BX3_r3b // ST3 (single structure)
	encoding_ST1_asisdlsop_HX1_r1h // ST1 (single structure)
	encoding_ST3_asisdlsop_HX3_r3h // ST3 (single structure)
	encoding_ST1_asisdlsop_SX1_r1s // ST1 (single structure)
	encoding_ST1_asisdlsop_DX1_r1d // ST1 (single structure)
	encoding_ST3_asisdlsop_SX3_r3s // ST3 (single structure)
	encoding_ST3_asisdlsop_DX3_r3d // ST3 (single structure)
	encoding_ST1_asisdlsop_B1_i1b  // ST1 (single structure)
	encoding_ST3_asisdlsop_B3_i3b  // ST3 (single structure)
	encoding_ST1_asisdlsop_H1_i1h  // ST1 (single structure)
	encoding_ST3_asisdlsop_H3_i3h  // ST3 (single structure)
	encoding_ST1_asisdlsop_S1_i1s  // ST1 (single structure)
	encoding_ST1_asisdlsop_D1_i1d  // ST1 (single structure)
	encoding_ST3_asisdlsop_S3_i3s  // ST3 (single structure)
	encoding_ST3_asisdlsop_D3_i3d  // ST3 (single structure)
	encoding_ST2_asisdlsop_BX2_r2b // ST2 (single structure)
	encoding_ST4_asisdlsop_BX4_r4b // ST4 (single structure)
	encoding_ST2_asisdlsop_HX2_r2h // ST2 (single structure)
	encoding_ST4_asisdlsop_HX4_r4h // ST4 (single structure)
	encoding_ST2_asisdlsop_SX2_r2s // ST2 (single structure)
	encoding_ST2_asisdlsop_DX2_r2d // ST2 (single structure)
	encoding_ST4_asisdlsop_SX4_r4s // ST4 (single structure)
	encoding_ST4_asisdlsop_DX4_r4d // ST4 (single structure)
	encoding_ST2_asisdlsop_B2_i2b  // ST2 (single structure)
	encoding_ST4_asisdlsop_B4_i4b  // ST4 (single structure)
	encoding_ST2_asisdlsop_H2_i2h  // ST2 (single structure)
	encoding_ST4_asisdlsop_H4_i4h  // ST4 (single structure)
	encoding_ST2_asisdlsop_S2_i2s  // ST2 (single structure)
	encoding_ST2_asisdlsop_D2_i2d  // ST2 (single structure)
	encoding_ST4_asisdlsop_S4_i4s  // ST4 (single structure)
	encoding_ST4_asisdlsop_D4_i4d  // ST4 (single structure)
	encoding_LD1_asisdlsop_BX1_r1b // LD1 (single structure)
	encoding_LD3_asisdlsop_BX3_r3b // LD3 (single structure)
	encoding_LD1_asisdlsop_HX1_r1h // LD1 (single structure)
	encoding_LD3_asisdlsop_HX3_r3h // LD3 (single structure)
	encoding_LD1_asisdlsop_SX1_r1s // LD1 (single structure)
	encoding_LD1_asisdlsop_DX1_r1d // LD1 (single structure)
	encoding_LD3_asisdlsop_SX3_r3s // LD3 (single structure)
	encoding_LD3_asisdlsop_DX3_r3d // LD3 (single structure)
	encoding_LD1R_asisdlsop_RX1_r  // LD1R
	encoding_LD3R_asisdlsop_RX3_r  // LD3R
	encoding_LD1_asisdlsop_B1_i1b  // LD1 (single structure)
	encoding_LD3_asisdlsop_B3_i3b  // LD3 (single structure)
	encoding_LD1_asisdlsop_H1_i1h  // LD1 (single structure)
	encoding_LD3_asisdlsop_H3_i3h  // LD3 (single structure)
	encoding_LD1_asisdlsop_S1_i1s  // LD1 (single structure)
	encoding_LD1_asisdlsop_D1_i1d  // LD1 (single structure)
	encoding_LD3_asisdlsop_S3_i3s  // LD3 (single structure)
	encoding_LD3_asisdlsop_D3_i3d  // LD3 (single structure)
	encoding_LD1R_asisdlsop_R1_i   // LD1R
	encoding_LD3R_asisdlsop_R3_i   // LD3R
	encoding_LD2_asisdlsop_BX2_r2b // LD2 (single structure)
	encoding_LD4_asisdlsop_BX4_r4b // LD4 (single structure)
	encoding_LD2_asisdlsop_HX2_r2h // LD2 (single structure)
	encoding_LD4_asisdlsop_HX4_r4h // LD4 (single structure)
	encoding_LD2_asisdlsop_SX2_r2s // LD2 (single structure)
	encoding_LD2_asisdlsop_DX2_r2d // LD2 (single structure)
	encoding_LD4_asisdlsop_SX4_r4s // LD4 (single structure)
	encoding_LD4_asisdlsop_DX4_r4d // LD4 (single structure)
	encoding_LD2R_asisdlsop_RX2_r  // LD2R
	encoding_LD4R_asisdlsop_RX4_r  // LD4R
	encoding_LD2_asisdlsop_B2_i2b  // LD2 (single structure)
	encoding_LD4_asisdlsop_B4_i4b  // LD4 (single structure)
	encoding_LD2_asisdlsop_H2_i2h  // LD2 (single structure)
	encoding_LD4_asisdlsop_H4_i4h  // LD4 (single structure)
	encoding_LD2_asisdlsop_S2_i2s  // LD2 (single structure)
	encoding_LD2_asisdlsop_D2_i2d  // LD2 (single structure)
	encoding_LD4_asisdlsop_S4_i4s  // LD4 (single structure)
	encoding_LD4_asisdlsop_D4_i4d  // LD4 (single structure)
	encoding_LD2R_asisdlsop_R2_i   // LD2R
	encoding_LD4R_asisdlsop_R4_i   // LD4R

	// iclass_ldsttags
	encoding_STG_64Spost_ldsttags     // STG
	encoding_STG_64Soffset_ldsttags   // STG
	encoding_STG_64Spre_ldsttags      // STG
	encoding_STZGM_64bulk_ldsttags    // STZGM
	encoding_LDG_64Loffset_ldsttags   // LDG
	encoding_STZG_64Spost_ldsttags    // STZG
	encoding_STZG_64Soffset_ldsttags  // STZG
	encoding_STZG_64Spre_ldsttags     // STZG
	encoding_ST2G_64Spost_ldsttags    // ST2G
	encoding_ST2G_64Soffset_ldsttags  // ST2G
	encoding_ST2G_64Spre_ldsttags     // ST2G
	encoding_STGM_64bulk_ldsttags     // STGM
	encoding_STZ2G_64Spost_ldsttags   // STZ2G
	encoding_STZ2G_64Soffset_ldsttags // STZ2G
	encoding_STZ2G_64Spre_ldsttags    // STZ2G
	encoding_LDGM_64bulk_ldsttags     // LDGM

	// iclass_ldstexclp
	encoding_STXP_SP32_ldstexclp  // STXP
	encoding_STLXP_SP32_ldstexclp // STLXP
	encoding_LDXP_LP32_ldstexclp  // LDXP
	encoding_LDAXP_LP32_ldstexclp // LDAXP
	encoding_STXP_SP64_ldstexclp  // STXP
	encoding_STLXP_SP64_ldstexclp // STLXP
	encoding_LDXP_LP64_ldstexclp  // LDXP
	encoding_LDAXP_LP64_ldstexclp // LDAXP

	// iclass_ldstexclr
	encoding_STXRB_SR32_ldstexclr  // STXRB
	encoding_STLXRB_SR32_ldstexclr // STLXRB
	encoding_LDXRB_LR32_ldstexclr  // LDXRB
	encoding_LDAXRB_LR32_ldstexclr // LDAXRB
	encoding_STXRH_SR32_ldstexclr  // STXRH
	encoding_STLXRH_SR32_ldstexclr // STLXRH
	encoding_LDXRH_LR32_ldstexclr  // LDXRH
	encoding_LDAXRH_LR32_ldstexclr // LDAXRH
	encoding_STXR_SR32_ldstexclr   // STXR
	encoding_STLXR_SR32_ldstexclr  // STLXR
	encoding_LDXR_LR32_ldstexclr   // LDXR
	encoding_LDAXR_LR32_ldstexclr  // LDAXR
	encoding_STXR_SR64_ldstexclr   // STXR
	encoding_STLXR_SR64_ldstexclr  // STLXR
	encoding_LDXR_LR64_ldstexclr   // LDXR
	encoding_LDAXR_LR64_ldstexclr  // LDAXR

	// iclass_ldstord
	encoding_STLLRB_SL32_ldstord // STLLRB
	encoding_STLRB_SL32_ldstord  // STLRB
	encoding_LDLARB_LR32_ldstord // LDLARB
	encoding_LDARB_LR32_ldstord  // LDARB
	encoding_STLLRH_SL32_ldstord // STLLRH
	encoding_STLRH_SL32_ldstord  // STLRH
	encoding_LDLARH_LR32_ldstord // LDLARH
	encoding_LDARH_LR32_ldstord  // LDARH
	encoding_STLLR_SL32_ldstord  // STLLR
	encoding_STLR_SL32_ldstord   // STLR
	encoding_LDLAR_LR32_ldstord  // LDLAR
	encoding_LDAR_LR32_ldstord   // LDAR
	encoding_STLLR_SL64_ldstord  // STLLR
	encoding_STLR_SL64_ldstord   // STLR
	encoding_LDLAR_LR64_ldstord  // LDLAR
	encoding_LDAR_LR64_ldstord   // LDAR

	// iclass_comswap
	encoding_CASB_C32_comswap   // CASB, CASAB, CASALB, CASLB
	encoding_CASLB_C32_comswap  // CASB, CASAB, CASALB, CASLB
	encoding_CASAB_C32_comswap  // CASB, CASAB, CASALB, CASLB
	encoding_CASALB_C32_comswap // CASB, CASAB, CASALB, CASLB
	encoding_CASH_C32_comswap   // CASH, CASAH, CASALH, CASLH
	encoding_CASLH_C32_comswap  // CASH, CASAH, CASALH, CASLH
	encoding_CASAH_C32_comswap  // CASH, CASAH, CASALH, CASLH
	encoding_CASALH_C32_comswap // CASH, CASAH, CASALH, CASLH
	encoding_CAS_C32_comswap    // CAS, CASA, CASAL, CASL
	encoding_CASL_C32_comswap   // CAS, CASA, CASAL, CASL
	encoding_CASA_C32_comswap   // CAS, CASA, CASAL, CASL
	encoding_CASAL_C32_comswap  // CAS, CASA, CASAL, CASL
	encoding_CAS_C64_comswap    // CAS, CASA, CASAL, CASL
	encoding_CASL_C64_comswap   // CAS, CASA, CASAL, CASL
	encoding_CASA_C64_comswap   // CAS, CASA, CASAL, CASL
	encoding_CASAL_C64_comswap  // CAS, CASA, CASAL, CASL

	// iclass_ldapstl_unscaled
	encoding_STLURB_32_ldapstl_unscaled   // STLURB
	encoding_LDAPURB_32_ldapstl_unscaled  // LDAPURB
	encoding_LDAPURSB_64_ldapstl_unscaled // LDAPURSB
	encoding_LDAPURSB_32_ldapstl_unscaled // LDAPURSB
	encoding_STLURH_32_ldapstl_unscaled   // STLURH
	encoding_LDAPURH_32_ldapstl_unscaled  // LDAPURH
	encoding_LDAPURSH_64_ldapstl_unscaled // LDAPURSH
	encoding_LDAPURSH_32_ldapstl_unscaled // LDAPURSH
	encoding_STLUR_32_ldapstl_unscaled    // STLUR
	encoding_LDAPUR_32_ldapstl_unscaled   // LDAPUR
	encoding_LDAPURSW_64_ldapstl_unscaled // LDAPURSW
	encoding_STLUR_64_ldapstl_unscaled    // STLUR
	encoding_LDAPUR_64_ldapstl_unscaled   // LDAPUR

	// iclass_loadlit
	encoding_LDR_32_loadlit   // LDR (literal)
	encoding_LDR_S_loadlit    // LDR (literal, SIMD&FP)
	encoding_LDR_64_loadlit   // LDR (literal)
	encoding_LDR_D_loadlit    // LDR (literal, SIMD&FP)
	encoding_LDRSW_64_loadlit // LDRSW (literal)
	encoding_LDR_Q_loadlit    // LDR (literal, SIMD&FP)
	encoding_PRFM_P_loadlit   // PRFM (literal)

	// iclass_memcms
	encoding_CPYFP_CPY_memcms     // CPYFP, CPYFM, CPYFE
	encoding_CPYFPWT_CPY_memcms   // CPYFPWT, CPYFMWT, CPYFEWT
	encoding_CPYFPRT_CPY_memcms   // CPYFPRT, CPYFMRT, CPYFERT
	encoding_CPYFPT_CPY_memcms    // CPYFPT, CPYFMT, CPYFET
	encoding_CPYFPWN_CPY_memcms   // CPYFPWN, CPYFMWN, CPYFEWN
	encoding_CPYFPWTWN_CPY_memcms // CPYFPWTWN, CPYFMWTWN, CPYFEWTWN
	encoding_CPYFPRTWN_CPY_memcms // CPYFPRTWN, CPYFMRTWN, CPYFERTWN
	encoding_CPYFPTWN_CPY_memcms  // CPYFPTWN, CPYFMTWN, CPYFETWN
	encoding_CPYFPRN_CPY_memcms   // CPYFPRN, CPYFMRN, CPYFERN
	encoding_CPYFPWTRN_CPY_memcms // CPYFPWTRN, CPYFMWTRN, CPYFEWTRN
	encoding_CPYFPRTRN_CPY_memcms // CPYFPRTRN, CPYFMRTRN, CPYFERTRN
	encoding_CPYFPTRN_CPY_memcms  // CPYFPTRN, CPYFMTRN, CPYFETRN
	encoding_CPYFPN_CPY_memcms    // CPYFPN, CPYFMN, CPYFEN
	encoding_CPYFPWTN_CPY_memcms  // CPYFPWTN, CPYFMWTN, CPYFEWTN
	encoding_CPYFPRTN_CPY_memcms  // CPYFPRTN, CPYFMRTN, CPYFERTN
	encoding_CPYFPTN_CPY_memcms   // CPYFPTN, CPYFMTN, CPYFETN
	encoding_CPYFM_CPY_memcms     // CPYFP, CPYFM, CPYFE
	encoding_CPYFMWT_CPY_memcms   // CPYFPWT, CPYFMWT, CPYFEWT
	encoding_CPYFMRT_CPY_memcms   // CPYFPRT, CPYFMRT, CPYFERT
	encoding_CPYFMT_CPY_memcms    // CPYFPT, CPYFMT, CPYFET
	encoding_CPYFMWN_CPY_memcms   // CPYFPWN, CPYFMWN, CPYFEWN
	encoding_CPYFMWTWN_CPY_memcms // CPYFPWTWN, CPYFMWTWN, CPYFEWTWN
	encoding_CPYFMRTWN_CPY_memcms // CPYFPRTWN, CPYFMRTWN, CPYFERTWN
	encoding_CPYFMTWN_CPY_memcms  // CPYFPTWN, CPYFMTWN, CPYFETWN
	encoding_CPYFMRN_CPY_memcms   // CPYFPRN, CPYFMRN, CPYFERN
	encoding_CPYFMWTRN_CPY_memcms // CPYFPWTRN, CPYFMWTRN, CPYFEWTRN
	encoding_CPYFMRTRN_CPY_memcms // CPYFPRTRN, CPYFMRTRN, CPYFERTRN
	encoding_CPYFMTRN_CPY_memcms  // CPYFPTRN, CPYFMTRN, CPYFETRN
	encoding_CPYFMN_CPY_memcms    // CPYFPN, CPYFMN, CPYFEN
	encoding_CPYFMWTN_CPY_memcms  // CPYFPWTN, CPYFMWTN, CPYFEWTN
	encoding_CPYFMRTN_CPY_memcms  // CPYFPRTN, CPYFMRTN, CPYFERTN
	encoding_CPYFMTN_CPY_memcms   // CPYFPTN, CPYFMTN, CPYFETN
	encoding_CPYFE_CPY_memcms     // CPYFP, CPYFM, CPYFE
	encoding_CPYFEWT_CPY_memcms   // CPYFPWT, CPYFMWT, CPYFEWT
	encoding_CPYFERT_CPY_memcms   // CPYFPRT, CPYFMRT, CPYFERT
	encoding_CPYFET_CPY_memcms    // CPYFPT, CPYFMT, CPYFET
	encoding_CPYFEWN_CPY_memcms   // CPYFPWN, CPYFMWN, CPYFEWN
	encoding_CPYFEWTWN_CPY_memcms // CPYFPWTWN, CPYFMWTWN, CPYFEWTWN
	encoding_CPYFERTWN_CPY_memcms // CPYFPRTWN, CPYFMRTWN, CPYFERTWN
	encoding_CPYFETWN_CPY_memcms  // CPYFPTWN, CPYFMTWN, CPYFETWN
	encoding_CPYFERN_CPY_memcms   // CPYFPRN, CPYFMRN, CPYFERN
	encoding_CPYFEWTRN_CPY_memcms // CPYFPWTRN, CPYFMWTRN, CPYFEWTRN
	encoding_CPYFERTRN_CPY_memcms // CPYFPRTRN, CPYFMRTRN, CPYFERTRN
	encoding_CPYFETRN_CPY_memcms  // CPYFPTRN, CPYFMTRN, CPYFETRN
	encoding_CPYFEN_CPY_memcms    // CPYFPN, CPYFMN, CPYFEN
	encoding_CPYFEWTN_CPY_memcms  // CPYFPWTN, CPYFMWTN, CPYFEWTN
	encoding_CPYFERTN_CPY_memcms  // CPYFPRTN, CPYFMRTN, CPYFERTN
	encoding_CPYFETN_CPY_memcms   // CPYFPTN, CPYFMTN, CPYFETN
	encoding_SETP_SET_memcms      // SETP, SETM, SETE
	encoding_SETPT_SET_memcms     // SETPT, SETMT, SETET
	encoding_SETPN_SET_memcms     // SETPN, SETMN, SETEN
	encoding_SETPTN_SET_memcms    // SETPTN, SETMTN, SETETN
	encoding_SETM_SET_memcms      // SETP, SETM, SETE
	encoding_SETMT_SET_memcms     // SETPT, SETMT, SETET
	encoding_SETMN_SET_memcms     // SETPN, SETMN, SETEN
	encoding_SETMTN_SET_memcms    // SETPTN, SETMTN, SETETN
	encoding_SETE_SET_memcms      // SETP, SETM, SETE
	encoding_SETET_SET_memcms     // SETPT, SETMT, SETET
	encoding_SETEN_SET_memcms     // SETPN, SETMN, SETEN
	encoding_SETETN_SET_memcms    // SETPTN, SETMTN, SETETN
	encoding_CPYP_CPY_memcms      // CPYP, CPYM, CPYE
	encoding_CPYPWT_CPY_memcms    // CPYPWT, CPYMWT, CPYEWT
	encoding_CPYPRT_CPY_memcms    // CPYPRT, CPYMRT, CPYERT
	encoding_CPYPT_CPY_memcms     // CPYPT, CPYMT, CPYET
	encoding_CPYPWN_CPY_memcms    // CPYPWN, CPYMWN, CPYEWN
	encoding_CPYPWTWN_CPY_memcms  // CPYPWTWN, CPYMWTWN, CPYEWTWN
	encoding_CPYPRTWN_CPY_memcms  // CPYPRTWN, CPYMRTWN, CPYERTWN
	encoding_CPYPTWN_CPY_memcms   // CPYPTWN, CPYMTWN, CPYETWN
	encoding_CPYPRN_CPY_memcms    // CPYPRN, CPYMRN, CPYERN
	encoding_CPYPWTRN_CPY_memcms  // CPYPWTRN, CPYMWTRN, CPYEWTRN
	encoding_CPYPRTRN_CPY_memcms  // CPYPRTRN, CPYMRTRN, CPYERTRN
	encoding_CPYPTRN_CPY_memcms   // CPYPTRN, CPYMTRN, CPYETRN
	encoding_CPYPN_CPY_memcms     // CPYPN, CPYMN, CPYEN
	encoding_CPYPWTN_CPY_memcms   // CPYPWTN, CPYMWTN, CPYEWTN
	encoding_CPYPRTN_CPY_memcms   // CPYPRTN, CPYMRTN, CPYERTN
	encoding_CPYPTN_CPY_memcms    // CPYPTN, CPYMTN, CPYETN
	encoding_CPYM_CPY_memcms      // CPYP, CPYM, CPYE
	encoding_CPYMWT_CPY_memcms    // CPYPWT, CPYMWT, CPYEWT
	encoding_CPYMRT_CPY_memcms    // CPYPRT, CPYMRT, CPYERT
	encoding_CPYMT_CPY_memcms     // CPYPT, CPYMT, CPYET
	encoding_CPYMWN_CPY_memcms    // CPYPWN, CPYMWN, CPYEWN
	encoding_CPYMWTWN_CPY_memcms  // CPYPWTWN, CPYMWTWN, CPYEWTWN
	encoding_CPYMRTWN_CPY_memcms  // CPYPRTWN, CPYMRTWN, CPYERTWN
	encoding_CPYMTWN_CPY_memcms   // CPYPTWN, CPYMTWN, CPYETWN
	encoding_CPYMRN_CPY_memcms    // CPYPRN, CPYMRN, CPYERN
	encoding_CPYMWTRN_CPY_memcms  // CPYPWTRN, CPYMWTRN, CPYEWTRN
	encoding_CPYMRTRN_CPY_memcms  // CPYPRTRN, CPYMRTRN, CPYERTRN
	encoding_CPYMTRN_CPY_memcms   // CPYPTRN, CPYMTRN, CPYETRN
	encoding_CPYMN_CPY_memcms     // CPYPN, CPYMN, CPYEN
	encoding_CPYMWTN_CPY_memcms   // CPYPWTN, CPYMWTN, CPYEWTN
	encoding_CPYMRTN_CPY_memcms   // CPYPRTN, CPYMRTN, CPYERTN
	encoding_CPYMTN_CPY_memcms    // CPYPTN, CPYMTN, CPYETN
	encoding_CPYE_CPY_memcms      // CPYP, CPYM, CPYE
	encoding_CPYEWT_CPY_memcms    // CPYPWT, CPYMWT, CPYEWT
	encoding_CPYERT_CPY_memcms    // CPYPRT, CPYMRT, CPYERT
	encoding_CPYET_CPY_memcms     // CPYPT, CPYMT, CPYET
	encoding_CPYEWN_CPY_memcms    // CPYPWN, CPYMWN, CPYEWN
	encoding_CPYEWTWN_CPY_memcms  // CPYPWTWN, CPYMWTWN, CPYEWTWN
	encoding_CPYERTWN_CPY_memcms  // CPYPRTWN, CPYMRTWN, CPYERTWN
	encoding_CPYETWN_CPY_memcms   // CPYPTWN, CPYMTWN, CPYETWN
	encoding_CPYERN_CPY_memcms    // CPYPRN, CPYMRN, CPYERN
	encoding_CPYEWTRN_CPY_memcms  // CPYPWTRN, CPYMWTRN, CPYEWTRN
	encoding_CPYERTRN_CPY_memcms  // CPYPRTRN, CPYMRTRN, CPYERTRN
	encoding_CPYETRN_CPY_memcms   // CPYPTRN, CPYMTRN, CPYETRN
	encoding_CPYEN_CPY_memcms     // CPYPN, CPYMN, CPYEN
	encoding_CPYEWTN_CPY_memcms   // CPYPWTN, CPYMWTN, CPYEWTN
	encoding_CPYERTN_CPY_memcms   // CPYPRTN, CPYMRTN, CPYERTN
	encoding_CPYETN_CPY_memcms    // CPYPTN, CPYMTN, CPYETN
	encoding_SETGP_SET_memcms     // SETGP, SETGM, SETGE
	encoding_SETGPT_SET_memcms    // SETGPT, SETGMT, SETGET
	encoding_SETGPN_SET_memcms    // SETGPN, SETGMN, SETGEN
	encoding_SETGPTN_SET_memcms   // SETGPTN, SETGMTN, SETGETN
	encoding_SETGM_SET_memcms     // SETGP, SETGM, SETGE
	encoding_SETGMT_SET_memcms    // SETGPT, SETGMT, SETGET
	encoding_SETGMN_SET_memcms    // SETGPN, SETGMN, SETGEN
	encoding_SETGMTN_SET_memcms   // SETGPTN, SETGMTN, SETGETN
	encoding_SETGE_SET_memcms     // SETGP, SETGM, SETGE
	encoding_SETGET_SET_memcms    // SETGPT, SETGMT, SETGET
	encoding_SETGEN_SET_memcms    // SETGPN, SETGMN, SETGEN
	encoding_SETGETN_SET_memcms   // SETGPTN, SETGMTN, SETGETN

	// iclass_ldstnapair_offs
	encoding_STNP_32_ldstnapair_offs // STNP
	encoding_LDNP_32_ldstnapair_offs // LDNP
	encoding_STNP_S_ldstnapair_offs  // STNP (SIMD&FP)
	encoding_LDNP_S_ldstnapair_offs  // LDNP (SIMD&FP)
	encoding_STNP_D_ldstnapair_offs  // STNP (SIMD&FP)
	encoding_LDNP_D_ldstnapair_offs  // LDNP (SIMD&FP)
	encoding_STNP_64_ldstnapair_offs // STNP
	encoding_LDNP_64_ldstnapair_offs // LDNP
	encoding_STNP_Q_ldstnapair_offs  // STNP (SIMD&FP)
	encoding_LDNP_Q_ldstnapair_offs  // LDNP (SIMD&FP)

	// iclass_ldstpair_post
	encoding_STP_32_ldstpair_post   // STP
	encoding_LDP_32_ldstpair_post   // LDP
	encoding_STP_S_ldstpair_post    // STP (SIMD&FP)
	encoding_LDP_S_ldstpair_post    // LDP (SIMD&FP)
	encoding_STGP_64_ldstpair_post  // STGP
	encoding_LDPSW_64_ldstpair_post // LDPSW
	encoding_STP_D_ldstpair_post    // STP (SIMD&FP)
	encoding_LDP_D_ldstpair_post    // LDP (SIMD&FP)
	encoding_STP_64_ldstpair_post   // STP
	encoding_LDP_64_ldstpair_post   // LDP
	encoding_STP_Q_ldstpair_post    // STP (SIMD&FP)
	encoding_LDP_Q_ldstpair_post    // LDP (SIMD&FP)

	// iclass_ldstpair_off
	encoding_STP_32_ldstpair_off   // STP
	encoding_LDP_32_ldstpair_off   // LDP
	encoding_STP_S_ldstpair_off    // STP (SIMD&FP)
	encoding_LDP_S_ldstpair_off    // LDP (SIMD&FP)
	encoding_STGP_64_ldstpair_off  // STGP
	encoding_LDPSW_64_ldstpair_off // LDPSW
	encoding_STP_D_ldstpair_off    // STP (SIMD&FP)
	encoding_LDP_D_ldstpair_off    // LDP (SIMD&FP)
	encoding_STP_64_ldstpair_off   // STP
	encoding_LDP_64_ldstpair_off   // LDP
	encoding_STP_Q_ldstpair_off    // STP (SIMD&FP)
	encoding_LDP_Q_ldstpair_off    // LDP (SIMD&FP)

	// iclass_ldstpair_pre
	encoding_STP_32_ldstpair_pre   // STP
	encoding_LDP_32_ldstpair_pre   // LDP
	encoding_STP_S_ldstpair_pre    // STP (SIMD&FP)
	encoding_LDP_S_ldstpair_pre    // LDP (SIMD&FP)
	encoding_STGP_64_ldstpair_pre  // STGP
	encoding_LDPSW_64_ldstpair_pre // LDPSW
	encoding_STP_D_ldstpair_pre    // STP (SIMD&FP)
	encoding_LDP_D_ldstpair_pre    // LDP (SIMD&FP)
	encoding_STP_64_ldstpair_pre   // STP
	encoding_LDP_64_ldstpair_pre   // LDP
	encoding_STP_Q_ldstpair_pre    // STP (SIMD&FP)
	encoding_LDP_Q_ldstpair_pre    // LDP (SIMD&FP)

	// iclass_ldst_unscaled
	encoding_STURB_32_ldst_unscaled  // STURB
	encoding_LDURB_32_ldst_unscaled  // LDURB
	encoding_LDURSB_64_ldst_unscaled // LDURSB
	encoding_LDURSB_32_ldst_unscaled // LDURSB
	encoding_STUR_B_ldst_unscaled    // STUR (SIMD&FP)
	encoding_LDUR_B_ldst_unscaled    // LDUR (SIMD&FP)
	encoding_STUR_Q_ldst_unscaled    // STUR (SIMD&FP)
	encoding_LDUR_Q_ldst_unscaled    // LDUR (SIMD&FP)
	encoding_STURH_32_ldst_unscaled  // STURH
	encoding_LDURH_32_ldst_unscaled  // LDURH
	encoding_LDURSH_64_ldst_unscaled // LDURSH
	encoding_LDURSH_32_ldst_unscaled // LDURSH
	encoding_STUR_H_ldst_unscaled    // STUR (SIMD&FP)
	encoding_LDUR_H_ldst_unscaled    // LDUR (SIMD&FP)
	encoding_STUR_32_ldst_unscaled   // STUR
	encoding_LDUR_32_ldst_unscaled   // LDUR
	encoding_LDURSW_64_ldst_unscaled // LDURSW
	encoding_STUR_S_ldst_unscaled    // STUR (SIMD&FP)
	encoding_LDUR_S_ldst_unscaled    // LDUR (SIMD&FP)
	encoding_STUR_64_ldst_unscaled   // STUR
	encoding_LDUR_64_ldst_unscaled   // LDUR
	encoding_PRFUM_P_ldst_unscaled   // PRFUM
	encoding_STUR_D_ldst_unscaled    // STUR (SIMD&FP)
	encoding_LDUR_D_ldst_unscaled    // LDUR (SIMD&FP)

	// iclass_ldst_immpost
	encoding_STRB_32_ldst_immpost  // STRB (immediate)
	encoding_LDRB_32_ldst_immpost  // LDRB (immediate)
	encoding_LDRSB_64_ldst_immpost // LDRSB (immediate)
	encoding_LDRSB_32_ldst_immpost // LDRSB (immediate)
	encoding_STR_B_ldst_immpost    // STR (immediate, SIMD&FP)
	encoding_LDR_B_ldst_immpost    // LDR (immediate, SIMD&FP)
	encoding_STR_Q_ldst_immpost    // STR (immediate, SIMD&FP)
	encoding_LDR_Q_ldst_immpost    // LDR (immediate, SIMD&FP)
	encoding_STRH_32_ldst_immpost  // STRH (immediate)
	encoding_LDRH_32_ldst_immpost  // LDRH (immediate)
	encoding_LDRSH_64_ldst_immpost // LDRSH (immediate)
	encoding_LDRSH_32_ldst_immpost // LDRSH (immediate)
	encoding_STR_H_ldst_immpost    // STR (immediate, SIMD&FP)
	encoding_LDR_H_ldst_immpost    // LDR (immediate, SIMD&FP)
	encoding_STR_32_ldst_immpost   // STR (immediate)
	encoding_LDR_32_ldst_immpost   // LDR (immediate)
	encoding_LDRSW_64_ldst_immpost // LDRSW (immediate)
	encoding_STR_S_ldst_immpost    // STR (immediate, SIMD&FP)
	encoding_LDR_S_ldst_immpost    // LDR (immediate, SIMD&FP)
	encoding_STR_64_ldst_immpost   // STR (immediate)
	encoding_LDR_64_ldst_immpost   // LDR (immediate)
	encoding_STR_D_ldst_immpost    // STR (immediate, SIMD&FP)
	encoding_LDR_D_ldst_immpost    // LDR (immediate, SIMD&FP)

	// iclass_ldst_unpriv
	encoding_STTRB_32_ldst_unpriv  // STTRB
	encoding_LDTRB_32_ldst_unpriv  // LDTRB
	encoding_LDTRSB_64_ldst_unpriv // LDTRSB
	encoding_LDTRSB_32_ldst_unpriv // LDTRSB
	encoding_STTRH_32_ldst_unpriv  // STTRH
	encoding_LDTRH_32_ldst_unpriv  // LDTRH
	encoding_LDTRSH_64_ldst_unpriv // LDTRSH
	encoding_LDTRSH_32_ldst_unpriv // LDTRSH
	encoding_STTR_32_ldst_unpriv   // STTR
	encoding_LDTR_32_ldst_unpriv   // LDTR
	encoding_LDTRSW_64_ldst_unpriv // LDTRSW
	encoding_STTR_64_ldst_unpriv   // STTR
	encoding_LDTR_64_ldst_unpriv   // LDTR

	// iclass_ldst_immpre
	encoding_STRB_32_ldst_immpre  // STRB (immediate)
	encoding_LDRB_32_ldst_immpre  // LDRB (immediate)
	encoding_LDRSB_64_ldst_immpre // LDRSB (immediate)
	encoding_LDRSB_32_ldst_immpre // LDRSB (immediate)
	encoding_STR_B_ldst_immpre    // STR (immediate, SIMD&FP)
	encoding_LDR_B_ldst_immpre    // LDR (immediate, SIMD&FP)
	encoding_STR_Q_ldst_immpre    // STR (immediate, SIMD&FP)
	encoding_LDR_Q_ldst_immpre    // LDR (immediate, SIMD&FP)
	encoding_STRH_32_ldst_immpre  // STRH (immediate)
	encoding_LDRH_32_ldst_immpre  // LDRH (immediate)
	encoding_LDRSH_64_ldst_immpre // LDRSH (immediate)
	encoding_LDRSH_32_ldst_immpre // LDRSH (immediate)
	encoding_STR_H_ldst_immpre    // STR (immediate, SIMD&FP)
	encoding_LDR_H_ldst_immpre    // LDR (immediate, SIMD&FP)
	encoding_STR_32_ldst_immpre   // STR (immediate)
	encoding_LDR_32_ldst_immpre   // LDR (immediate)
	encoding_LDRSW_64_ldst_immpre // LDRSW (immediate)
	encoding_STR_S_ldst_immpre    // STR (immediate, SIMD&FP)
	encoding_LDR_S_ldst_immpre    // LDR (immediate, SIMD&FP)
	encoding_STR_64_ldst_immpre   // STR (immediate)
	encoding_LDR_64_ldst_immpre   // LDR (immediate)
	encoding_STR_D_ldst_immpre    // STR (immediate, SIMD&FP)
	encoding_LDR_D_ldst_immpre    // LDR (immediate, SIMD&FP)

	// iclass_memop
	encoding_LDADDB_32_memop    // LDADDB, LDADDAB, LDADDALB, LDADDLB
	encoding_LDCLRB_32_memop    // LDCLRB, LDCLRAB, LDCLRALB, LDCLRLB
	encoding_LDEORB_32_memop    // LDEORB, LDEORAB, LDEORALB, LDEORLB
	encoding_LDSETB_32_memop    // LDSETB, LDSETAB, LDSETALB, LDSETLB
	encoding_LDSMAXB_32_memop   // LDSMAXB, LDSMAXAB, LDSMAXALB, LDSMAXLB
	encoding_LDSMINB_32_memop   // LDSMINB, LDSMINAB, LDSMINALB, LDSMINLB
	encoding_LDUMAXB_32_memop   // LDUMAXB, LDUMAXAB, LDUMAXALB, LDUMAXLB
	encoding_LDUMINB_32_memop   // LDUMINB, LDUMINAB, LDUMINALB, LDUMINLB
	encoding_SWPB_32_memop      // SWPB, SWPAB, SWPALB, SWPLB
	encoding_LDADDLB_32_memop   // LDADDB, LDADDAB, LDADDALB, LDADDLB
	encoding_LDCLRLB_32_memop   // LDCLRB, LDCLRAB, LDCLRALB, LDCLRLB
	encoding_LDEORLB_32_memop   // LDEORB, LDEORAB, LDEORALB, LDEORLB
	encoding_LDSETLB_32_memop   // LDSETB, LDSETAB, LDSETALB, LDSETLB
	encoding_LDSMAXLB_32_memop  // LDSMAXB, LDSMAXAB, LDSMAXALB, LDSMAXLB
	encoding_LDSMINLB_32_memop  // LDSMINB, LDSMINAB, LDSMINALB, LDSMINLB
	encoding_LDUMAXLB_32_memop  // LDUMAXB, LDUMAXAB, LDUMAXALB, LDUMAXLB
	encoding_LDUMINLB_32_memop  // LDUMINB, LDUMINAB, LDUMINALB, LDUMINLB
	encoding_SWPLB_32_memop     // SWPB, SWPAB, SWPALB, SWPLB
	encoding_LDADDAB_32_memop   // LDADDB, LDADDAB, LDADDALB, LDADDLB
	encoding_LDCLRAB_32_memop   // LDCLRB, LDCLRAB, LDCLRALB, LDCLRLB
	encoding_LDEORAB_32_memop   // LDEORB, LDEORAB, LDEORALB, LDEORLB
	encoding_LDSETAB_32_memop   // LDSETB, LDSETAB, LDSETALB, LDSETLB
	encoding_LDSMAXAB_32_memop  // LDSMAXB, LDSMAXAB, LDSMAXALB, LDSMAXLB
	encoding_LDSMINAB_32_memop  // LDSMINB, LDSMINAB, LDSMINALB, LDSMINLB
	encoding_LDUMAXAB_32_memop  // LDUMAXB, LDUMAXAB, LDUMAXALB, LDUMAXLB
	encoding_LDUMINAB_32_memop  // LDUMINB, LDUMINAB, LDUMINALB, LDUMINLB
	encoding_SWPAB_32_memop     // SWPB, SWPAB, SWPALB, SWPLB
	encoding_LDAPRB_32L_memop   // LDAPRB
	encoding_LDADDALB_32_memop  // LDADDB, LDADDAB, LDADDALB, LDADDLB
	encoding_LDCLRALB_32_memop  // LDCLRB, LDCLRAB, LDCLRALB, LDCLRLB
	encoding_LDEORALB_32_memop  // LDEORB, LDEORAB, LDEORALB, LDEORLB
	encoding_LDSETALB_32_memop  // LDSETB, LDSETAB, LDSETALB, LDSETLB
	encoding_LDSMAXALB_32_memop // LDSMAXB, LDSMAXAB, LDSMAXALB, LDSMAXLB
	encoding_LDSMINALB_32_memop // LDSMINB, LDSMINAB, LDSMINALB, LDSMINLB
	encoding_LDUMAXALB_32_memop // LDUMAXB, LDUMAXAB, LDUMAXALB, LDUMAXLB
	encoding_LDUMINALB_32_memop // LDUMINB, LDUMINAB, LDUMINALB, LDUMINLB
	encoding_SWPALB_32_memop    // SWPB, SWPAB, SWPALB, SWPLB
	encoding_LDADDH_32_memop    // LDADDH, LDADDAH, LDADDALH, LDADDLH
	encoding_LDCLRH_32_memop    // LDCLRH, LDCLRAH, LDCLRALH, LDCLRLH
	encoding_LDEORH_32_memop    // LDEORH, LDEORAH, LDEORALH, LDEORLH
	encoding_LDSETH_32_memop    // LDSETH, LDSETAH, LDSETALH, LDSETLH
	encoding_LDSMAXH_32_memop   // LDSMAXH, LDSMAXAH, LDSMAXALH, LDSMAXLH
	encoding_LDSMINH_32_memop   // LDSMINH, LDSMINAH, LDSMINALH, LDSMINLH
	encoding_LDUMAXH_32_memop   // LDUMAXH, LDUMAXAH, LDUMAXALH, LDUMAXLH
	encoding_LDUMINH_32_memop   // LDUMINH, LDUMINAH, LDUMINALH, LDUMINLH
	encoding_SWPH_32_memop      // SWPH, SWPAH, SWPALH, SWPLH
	encoding_LDADDLH_32_memop   // LDADDH, LDADDAH, LDADDALH, LDADDLH
	encoding_LDCLRLH_32_memop   // LDCLRH, LDCLRAH, LDCLRALH, LDCLRLH
	encoding_LDEORLH_32_memop   // LDEORH, LDEORAH, LDEORALH, LDEORLH
	encoding_LDSETLH_32_memop   // LDSETH, LDSETAH, LDSETALH, LDSETLH
	encoding_LDSMAXLH_32_memop  // LDSMAXH, LDSMAXAH, LDSMAXALH, LDSMAXLH
	encoding_LDSMINLH_32_memop  // LDSMINH, LDSMINAH, LDSMINALH, LDSMINLH
	encoding_LDUMAXLH_32_memop  // LDUMAXH, LDUMAXAH, LDUMAXALH, LDUMAXLH
	encoding_LDUMINLH_32_memop  // LDUMINH, LDUMINAH, LDUMINALH, LDUMINLH
	encoding_SWPLH_32_memop     // SWPH, SWPAH, SWPALH, SWPLH
	encoding_LDADDAH_32_memop   // LDADDH, LDADDAH, LDADDALH, LDADDLH
	encoding_LDCLRAH_32_memop   // LDCLRH, LDCLRAH, LDCLRALH, LDCLRLH
	encoding_LDEORAH_32_memop   // LDEORH, LDEORAH, LDEORALH, LDEORLH
	encoding_LDSETAH_32_memop   // LDSETH, LDSETAH, LDSETALH, LDSETLH
	encoding_LDSMAXAH_32_memop  // LDSMAXH, LDSMAXAH, LDSMAXALH, LDSMAXLH
	encoding_LDSMINAH_32_memop  // LDSMINH, LDSMINAH, LDSMINALH, LDSMINLH
	encoding_LDUMAXAH_32_memop  // LDUMAXH, LDUMAXAH, LDUMAXALH, LDUMAXLH
	encoding_LDUMINAH_32_memop  // LDUMINH, LDUMINAH, LDUMINALH, LDUMINLH
	encoding_SWPAH_32_memop     // SWPH, SWPAH, SWPALH, SWPLH
	encoding_LDAPRH_32L_memop   // LDAPRH
	encoding_LDADDALH_32_memop  // LDADDH, LDADDAH, LDADDALH, LDADDLH
	encoding_LDCLRALH_32_memop  // LDCLRH, LDCLRAH, LDCLRALH, LDCLRLH
	encoding_LDEORALH_32_memop  // LDEORH, LDEORAH, LDEORALH, LDEORLH
	encoding_LDSETALH_32_memop  // LDSETH, LDSETAH, LDSETALH, LDSETLH
	encoding_LDSMAXALH_32_memop // LDSMAXH, LDSMAXAH, LDSMAXALH, LDSMAXLH
	encoding_LDSMINALH_32_memop // LDSMINH, LDSMINAH, LDSMINALH, LDSMINLH
	encoding_LDUMAXALH_32_memop // LDUMAXH, LDUMAXAH, LDUMAXALH, LDUMAXLH
	encoding_LDUMINALH_32_memop // LDUMINH, LDUMINAH, LDUMINALH, LDUMINLH
	encoding_SWPALH_32_memop    // SWPH, SWPAH, SWPALH, SWPLH
	encoding_LDADD_32_memop     // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLR_32_memop     // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEOR_32_memop     // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSET_32_memop     // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAX_32_memop    // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMIN_32_memop    // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAX_32_memop    // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMIN_32_memop    // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWP_32_memop       // SWP, SWPA, SWPAL, SWPL
	encoding_LDADDL_32_memop    // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLRL_32_memop    // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEORL_32_memop    // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSETL_32_memop    // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAXL_32_memop   // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMINL_32_memop   // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAXL_32_memop   // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMINL_32_memop   // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWPL_32_memop      // SWP, SWPA, SWPAL, SWPL
	encoding_LDADDA_32_memop    // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLRA_32_memop    // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEORA_32_memop    // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSETA_32_memop    // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAXA_32_memop   // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMINA_32_memop   // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAXA_32_memop   // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMINA_32_memop   // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWPA_32_memop      // SWP, SWPA, SWPAL, SWPL
	encoding_LDAPR_32L_memop    // LDAPR
	encoding_LDADDAL_32_memop   // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLRAL_32_memop   // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEORAL_32_memop   // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSETAL_32_memop   // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAXAL_32_memop  // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMINAL_32_memop  // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAXAL_32_memop  // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMINAL_32_memop  // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWPAL_32_memop     // SWP, SWPA, SWPAL, SWPL
	encoding_LDADD_64_memop     // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLR_64_memop     // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEOR_64_memop     // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSET_64_memop     // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAX_64_memop    // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMIN_64_memop    // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAX_64_memop    // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMIN_64_memop    // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWP_64_memop       // SWP, SWPA, SWPAL, SWPL
	encoding_ST64BV0_64_memop   // ST64BV0
	encoding_ST64BV_64_memop    // ST64BV
	encoding_ST64B_64L_memop    // ST64B
	encoding_LD64B_64L_memop    // LD64B
	encoding_LDADDL_64_memop    // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLRL_64_memop    // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEORL_64_memop    // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSETL_64_memop    // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAXL_64_memop   // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMINL_64_memop   // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAXL_64_memop   // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMINL_64_memop   // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWPL_64_memop      // SWP, SWPA, SWPAL, SWPL
	encoding_LDADDA_64_memop    // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLRA_64_memop    // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEORA_64_memop    // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSETA_64_memop    // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAXA_64_memop   // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMINA_64_memop   // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAXA_64_memop   // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMINA_64_memop   // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWPA_64_memop      // SWP, SWPA, SWPAL, SWPL
	encoding_LDAPR_64L_memop    // LDAPR
	encoding_LDADDAL_64_memop   // LDADD, LDADDA, LDADDAL, LDADDL
	encoding_LDCLRAL_64_memop   // LDCLR, LDCLRA, LDCLRAL, LDCLRL
	encoding_LDEORAL_64_memop   // LDEOR, LDEORA, LDEORAL, LDEORL
	encoding_LDSETAL_64_memop   // LDSET, LDSETA, LDSETAL, LDSETL
	encoding_LDSMAXAL_64_memop  // LDSMAX, LDSMAXA, LDSMAXAL, LDSMAXL
	encoding_LDSMINAL_64_memop  // LDSMIN, LDSMINA, LDSMINAL, LDSMINL
	encoding_LDUMAXAL_64_memop  // LDUMAX, LDUMAXA, LDUMAXAL, LDUMAXL
	encoding_LDUMINAL_64_memop  // LDUMIN, LDUMINA, LDUMINAL, LDUMINL
	encoding_SWPAL_64_memop     // SWP, SWPA, SWPAL, SWPL

	// iclass_ldst_regoff
	encoding_STRB_32B_ldst_regoff   // STRB (register)
	encoding_STRB_32BL_ldst_regoff  // STRB (register)
	encoding_LDRB_32B_ldst_regoff   // LDRB (register)
	encoding_LDRB_32BL_ldst_regoff  // LDRB (register)
	encoding_LDRSB_64B_ldst_regoff  // LDRSB (register)
	encoding_LDRSB_64BL_ldst_regoff // LDRSB (register)
	encoding_LDRSB_32B_ldst_regoff  // LDRSB (register)
	encoding_LDRSB_32BL_ldst_regoff // LDRSB (register)
	encoding_STR_B_ldst_regoff      // STR (register, SIMD&FP)
	encoding_STR_BL_ldst_regoff     // STR (register, SIMD&FP)
	encoding_LDR_B_ldst_regoff      // LDR (register, SIMD&FP)
	encoding_LDR_BL_ldst_regoff     // LDR (register, SIMD&FP)
	encoding_STR_Q_ldst_regoff      // STR (register, SIMD&FP)
	encoding_LDR_Q_ldst_regoff      // LDR (register, SIMD&FP)
	encoding_STRH_32_ldst_regoff    // STRH (register)
	encoding_LDRH_32_ldst_regoff    // LDRH (register)
	encoding_LDRSH_64_ldst_regoff   // LDRSH (register)
	encoding_LDRSH_32_ldst_regoff   // LDRSH (register)
	encoding_STR_H_ldst_regoff      // STR (register, SIMD&FP)
	encoding_LDR_H_ldst_regoff      // LDR (register, SIMD&FP)
	encoding_STR_32_ldst_regoff     // STR (register)
	encoding_LDR_32_ldst_regoff     // LDR (register)
	encoding_LDRSW_64_ldst_regoff   // LDRSW (register)
	encoding_STR_S_ldst_regoff      // STR (register, SIMD&FP)
	encoding_LDR_S_ldst_regoff      // LDR (register, SIMD&FP)
	encoding_STR_64_ldst_regoff     // STR (register)
	encoding_LDR_64_ldst_regoff     // LDR (register)
	encoding_PRFM_P_ldst_regoff     // PRFM (register)
	encoding_STR_D_ldst_regoff      // STR (register, SIMD&FP)
	encoding_LDR_D_ldst_regoff      // LDR (register, SIMD&FP)

	// iclass_ldst_pac
	encoding_LDRAA_64_ldst_pac  // LDRAA, LDRAB
	encoding_LDRAA_64W_ldst_pac // LDRAA, LDRAB
	encoding_LDRAB_64_ldst_pac  // LDRAA, LDRAB
	encoding_LDRAB_64W_ldst_pac // LDRAA, LDRAB

	// iclass_ldst_pos
	encoding_STRB_32_ldst_pos  // STRB (immediate)
	encoding_LDRB_32_ldst_pos  // LDRB (immediate)
	encoding_LDRSB_64_ldst_pos // LDRSB (immediate)
	encoding_LDRSB_32_ldst_pos // LDRSB (immediate)
	encoding_STR_B_ldst_pos    // STR (immediate, SIMD&FP)
	encoding_LDR_B_ldst_pos    // LDR (immediate, SIMD&FP)
	encoding_STR_Q_ldst_pos    // STR (immediate, SIMD&FP)
	encoding_LDR_Q_ldst_pos    // LDR (immediate, SIMD&FP)
	encoding_STRH_32_ldst_pos  // STRH (immediate)
	encoding_LDRH_32_ldst_pos  // LDRH (immediate)
	encoding_LDRSH_64_ldst_pos // LDRSH (immediate)
	encoding_LDRSH_32_ldst_pos // LDRSH (immediate)
	encoding_STR_H_ldst_pos    // STR (immediate, SIMD&FP)
	encoding_LDR_H_ldst_pos    // LDR (immediate, SIMD&FP)
	encoding_STR_32_ldst_pos   // STR (immediate)
	encoding_LDR_32_ldst_pos   // LDR (immediate)
	encoding_LDRSW_64_ldst_pos // LDRSW (immediate)
	encoding_STR_S_ldst_pos    // STR (immediate, SIMD&FP)
	encoding_LDR_S_ldst_pos    // LDR (immediate, SIMD&FP)
	encoding_STR_64_ldst_pos   // STR (immediate)
	encoding_LDR_64_ldst_pos   // LDR (immediate)
	encoding_PRFM_P_ldst_pos   // PRFM (immediate)
	encoding_STR_D_ldst_pos    // STR (immediate, SIMD&FP)
	encoding_LDR_D_ldst_pos    // LDR (immediate, SIMD&FP)

	// iclass_dp_2src
	encoding_UDIV_32_dp_2src     // UDIV
	encoding_SDIV_32_dp_2src     // SDIV
	encoding_LSLV_32_dp_2src     // LSLV
	encoding_LSRV_32_dp_2src     // LSRV
	encoding_ASRV_32_dp_2src     // ASRV
	encoding_RORV_32_dp_2src     // RORV
	encoding_CRC32B_32C_dp_2src  // CRC32B, CRC32H, CRC32W, CRC32X
	encoding_CRC32H_32C_dp_2src  // CRC32B, CRC32H, CRC32W, CRC32X
	encoding_CRC32W_32C_dp_2src  // CRC32B, CRC32H, CRC32W, CRC32X
	encoding_CRC32CB_32C_dp_2src // CRC32CB, CRC32CH, CRC32CW, CRC32CX
	encoding_CRC32CH_32C_dp_2src // CRC32CB, CRC32CH, CRC32CW, CRC32CX
	encoding_CRC32CW_32C_dp_2src // CRC32CB, CRC32CH, CRC32CW, CRC32CX
	encoding_SUBP_64S_dp_2src    // SUBP
	encoding_UDIV_64_dp_2src     // UDIV
	encoding_SDIV_64_dp_2src     // SDIV
	encoding_IRG_64I_dp_2src     // IRG
	encoding_GMI_64G_dp_2src     // GMI
	encoding_LSLV_64_dp_2src     // LSLV
	encoding_LSRV_64_dp_2src     // LSRV
	encoding_ASRV_64_dp_2src     // ASRV
	encoding_RORV_64_dp_2src     // RORV
	encoding_PACGA_64P_dp_2src   // PACGA
	encoding_CRC32X_64C_dp_2src  // CRC32B, CRC32H, CRC32W, CRC32X
	encoding_CRC32CX_64C_dp_2src // CRC32CB, CRC32CH, CRC32CW, CRC32CX
	encoding_SUBPS_64S_dp_2src   // SUBPS

	// iclass_dp_1src
	encoding_RBIT_32_dp_1src    // RBIT
	encoding_REV16_32_dp_1src   // REV16
	encoding_REV_32_dp_1src     // REV
	encoding_CLZ_32_dp_1src     // CLZ
	encoding_CLS_32_dp_1src     // CLS
	encoding_RBIT_64_dp_1src    // RBIT
	encoding_REV16_64_dp_1src   // REV16
	encoding_REV32_64_dp_1src   // REV32
	encoding_REV_64_dp_1src     // REV
	encoding_CLZ_64_dp_1src     // CLZ
	encoding_CLS_64_dp_1src     // CLS
	encoding_PACIA_64P_dp_1src  // PACIA, PACIA1716, PACIASP, PACIAZ, PACIZA
	encoding_PACIB_64P_dp_1src  // PACIB, PACIB1716, PACIBSP, PACIBZ, PACIZB
	encoding_PACDA_64P_dp_1src  // PACDA, PACDZA
	encoding_PACDB_64P_dp_1src  // PACDB, PACDZB
	encoding_AUTIA_64P_dp_1src  // AUTIA, AUTIA1716, AUTIASP, AUTIAZ, AUTIZA
	encoding_AUTIB_64P_dp_1src  // AUTIB, AUTIB1716, AUTIBSP, AUTIBZ, AUTIZB
	encoding_AUTDA_64P_dp_1src  // AUTDA, AUTDZA
	encoding_AUTDB_64P_dp_1src  // AUTDB, AUTDZB
	encoding_PACIZA_64Z_dp_1src // PACIA, PACIA1716, PACIASP, PACIAZ, PACIZA
	encoding_PACIZB_64Z_dp_1src // PACIB, PACIB1716, PACIBSP, PACIBZ, PACIZB
	encoding_PACDZA_64Z_dp_1src // PACDA, PACDZA
	encoding_PACDZB_64Z_dp_1src // PACDB, PACDZB
	encoding_AUTIZA_64Z_dp_1src // AUTIA, AUTIA1716, AUTIASP, AUTIAZ, AUTIZA
	encoding_AUTIZB_64Z_dp_1src // AUTIB, AUTIB1716, AUTIBSP, AUTIBZ, AUTIZB
	encoding_AUTDZA_64Z_dp_1src // AUTDA, AUTDZA
	encoding_AUTDZB_64Z_dp_1src // AUTDB, AUTDZB
	encoding_XPACI_64Z_dp_1src  // XPACD, XPACI, XPACLRI
	encoding_XPACD_64Z_dp_1src  // XPACD, XPACI, XPACLRI

	// iclass_log_shift
	encoding_AND_32_log_shift  // AND (shifted register)
	encoding_BIC_32_log_shift  // BIC (shifted register)
	encoding_ORR_32_log_shift  // ORR (shifted register)
	encoding_ORN_32_log_shift  // ORN (shifted register)
	encoding_EOR_32_log_shift  // EOR (shifted register)
	encoding_EON_32_log_shift  // EON (shifted register)
	encoding_ANDS_32_log_shift // ANDS (shifted register)
	encoding_BICS_32_log_shift // BICS (shifted register)
	encoding_AND_64_log_shift  // AND (shifted register)
	encoding_BIC_64_log_shift  // BIC (shifted register)
	encoding_ORR_64_log_shift  // ORR (shifted register)
	encoding_ORN_64_log_shift  // ORN (shifted register)
	encoding_EOR_64_log_shift  // EOR (shifted register)
	encoding_EON_64_log_shift  // EON (shifted register)
	encoding_ANDS_64_log_shift // ANDS (shifted register)
	encoding_BICS_64_log_shift // BICS (shifted register)

	// iclass_addsub_shift
	encoding_ADD_32_addsub_shift  // ADD (shifted register)
	encoding_ADDS_32_addsub_shift // ADDS (shifted register)
	encoding_SUB_32_addsub_shift  // SUB (shifted register)
	encoding_SUBS_32_addsub_shift // SUBS (shifted register)
	encoding_ADD_64_addsub_shift  // ADD (shifted register)
	encoding_ADDS_64_addsub_shift // ADDS (shifted register)
	encoding_SUB_64_addsub_shift  // SUB (shifted register)
	encoding_SUBS_64_addsub_shift // SUBS (shifted register)

	// iclass_addsub_ext
	encoding_ADD_32_addsub_ext   // ADD (extended register)
	encoding_ADDS_32S_addsub_ext // ADDS (extended register)
	encoding_SUB_32_addsub_ext   // SUB (extended register)
	encoding_SUBS_32S_addsub_ext // SUBS (extended register)
	encoding_ADD_64_addsub_ext   // ADD (extended register)
	encoding_ADDS_64S_addsub_ext // ADDS (extended register)
	encoding_SUB_64_addsub_ext   // SUB (extended register)
	encoding_SUBS_64S_addsub_ext // SUBS (extended register)

	// iclass_addsub_carry
	encoding_ADC_32_addsub_carry  // ADC
	encoding_ADCS_32_addsub_carry // ADCS
	encoding_SBC_32_addsub_carry  // SBC
	encoding_SBCS_32_addsub_carry // SBCS
	encoding_ADC_64_addsub_carry  // ADC
	encoding_ADCS_64_addsub_carry // ADCS
	encoding_SBC_64_addsub_carry  // SBC
	encoding_SBCS_64_addsub_carry // SBCS

	// iclass_rmif
	encoding_RMIF_only_rmif // RMIF

	// iclass_setf
	encoding_SETF8_only_setf  // SETF8, SETF16
	encoding_SETF16_only_setf // SETF8, SETF16

	// iclass_condcmp_reg
	encoding_CCMN_32_condcmp_reg // CCMN (register)
	encoding_CCMP_32_condcmp_reg // CCMP (register)
	encoding_CCMN_64_condcmp_reg // CCMN (register)
	encoding_CCMP_64_condcmp_reg // CCMP (register)

	// iclass_condcmp_imm
	encoding_CCMN_32_condcmp_imm // CCMN (immediate)
	encoding_CCMP_32_condcmp_imm // CCMP (immediate)
	encoding_CCMN_64_condcmp_imm // CCMN (immediate)
	encoding_CCMP_64_condcmp_imm // CCMP (immediate)

	// iclass_condsel
	encoding_CSEL_32_condsel  // CSEL
	encoding_CSINC_32_condsel // CSINC
	encoding_CSINV_32_condsel // CSINV
	encoding_CSNEG_32_condsel // CSNEG
	encoding_CSEL_64_condsel  // CSEL
	encoding_CSINC_64_condsel // CSINC
	encoding_CSINV_64_condsel // CSINV
	encoding_CSNEG_64_condsel // CSNEG

	// iclass_dp_3src
	encoding_MADD_32A_dp_3src    // MADD
	encoding_MSUB_32A_dp_3src    // MSUB
	encoding_MADD_64A_dp_3src    // MADD
	encoding_MSUB_64A_dp_3src    // MSUB
	encoding_SMADDL_64WA_dp_3src // SMADDL
	encoding_SMSUBL_64WA_dp_3src // SMSUBL
	encoding_SMULH_64_dp_3src    // SMULH
	encoding_UMADDL_64WA_dp_3src // UMADDL
	encoding_UMSUBL_64WA_dp_3src // UMSUBL
	encoding_UMULH_64_dp_3src    // UMULH

	// iclass_cryptoaes
	encoding_AESE_B_cryptoaes   // AESE
	encoding_AESD_B_cryptoaes   // AESD
	encoding_AESMC_B_cryptoaes  // AESMC
	encoding_AESIMC_B_cryptoaes // AESIMC

	// iclass_cryptosha3
	encoding_SHA1C_QSV_cryptosha3     // SHA1C
	encoding_SHA1P_QSV_cryptosha3     // SHA1P
	encoding_SHA1M_QSV_cryptosha3     // SHA1M
	encoding_SHA1SU0_VVV_cryptosha3   // SHA1SU0
	encoding_SHA256H_QQV_cryptosha3   // SHA256H
	encoding_SHA256H2_QQV_cryptosha3  // SHA256H2
	encoding_SHA256SU1_VVV_cryptosha3 // SHA256SU1

	// iclass_cryptosha2
	encoding_SHA1H_SS_cryptosha2     // SHA1H
	encoding_SHA1SU1_VV_cryptosha2   // SHA1SU1
	encoding_SHA256SU0_VV_cryptosha2 // SHA256SU0

	// iclass_asisdone
	encoding_DUP_asisdone_only // DUP (element)

	// iclass_asisdsamefp16
	encoding_FMULX_asisdsamefp16_only   // FMULX
	encoding_FCMEQ_asisdsamefp16_only   // FCMEQ (register)
	encoding_FRECPS_asisdsamefp16_only  // FRECPS
	encoding_FRSQRTS_asisdsamefp16_only // FRSQRTS
	encoding_FCMGE_asisdsamefp16_only   // FCMGE (register)
	encoding_FACGE_asisdsamefp16_only   // FACGE
	encoding_FABD_asisdsamefp16_only    // FABD
	encoding_FCMGT_asisdsamefp16_only   // FCMGT (register)
	encoding_FACGT_asisdsamefp16_only   // FACGT

	// iclass_asisdmiscfp16
	encoding_FCVTNS_asisdmiscfp16_R  // FCVTNS (vector)
	encoding_FCVTMS_asisdmiscfp16_R  // FCVTMS (vector)
	encoding_FCVTAS_asisdmiscfp16_R  // FCVTAS (vector)
	encoding_SCVTF_asisdmiscfp16_R   // SCVTF (vector, integer)
	encoding_FCMGT_asisdmiscfp16_FZ  // FCMGT (zero)
	encoding_FCMEQ_asisdmiscfp16_FZ  // FCMEQ (zero)
	encoding_FCMLT_asisdmiscfp16_FZ  // FCMLT (zero)
	encoding_FCVTPS_asisdmiscfp16_R  // FCVTPS (vector)
	encoding_FCVTZS_asisdmiscfp16_R  // FCVTZS (vector, integer)
	encoding_FRECPE_asisdmiscfp16_R  // FRECPE
	encoding_FRECPX_asisdmiscfp16_R  // FRECPX
	encoding_FCVTNU_asisdmiscfp16_R  // FCVTNU (vector)
	encoding_FCVTMU_asisdmiscfp16_R  // FCVTMU (vector)
	encoding_FCVTAU_asisdmiscfp16_R  // FCVTAU (vector)
	encoding_UCVTF_asisdmiscfp16_R   // UCVTF (vector, integer)
	encoding_FCMGE_asisdmiscfp16_FZ  // FCMGE (zero)
	encoding_FCMLE_asisdmiscfp16_FZ  // FCMLE (zero)
	encoding_FCVTPU_asisdmiscfp16_R  // FCVTPU (vector)
	encoding_FCVTZU_asisdmiscfp16_R  // FCVTZU (vector, integer)
	encoding_FRSQRTE_asisdmiscfp16_R // FRSQRTE

	// iclass_asisdsame2
	encoding_SQRDMLAH_asisdsame2_only // SQRDMLAH (vector)
	encoding_SQRDMLSH_asisdsame2_only // SQRDMLSH (vector)

	// iclass_asisdmisc
	encoding_SUQADD_asisdmisc_R  // SUQADD
	encoding_SQABS_asisdmisc_R   // SQABS
	encoding_CMGT_asisdmisc_Z    // CMGT (zero)
	encoding_CMEQ_asisdmisc_Z    // CMEQ (zero)
	encoding_CMLT_asisdmisc_Z    // CMLT (zero)
	encoding_ABS_asisdmisc_R     // ABS
	encoding_SQXTN_asisdmisc_N   // SQXTN, SQXTN2
	encoding_FCVTNS_asisdmisc_R  // FCVTNS (vector)
	encoding_FCVTMS_asisdmisc_R  // FCVTMS (vector)
	encoding_FCVTAS_asisdmisc_R  // FCVTAS (vector)
	encoding_SCVTF_asisdmisc_R   // SCVTF (vector, integer)
	encoding_FCMGT_asisdmisc_FZ  // FCMGT (zero)
	encoding_FCMEQ_asisdmisc_FZ  // FCMEQ (zero)
	encoding_FCMLT_asisdmisc_FZ  // FCMLT (zero)
	encoding_FCVTPS_asisdmisc_R  // FCVTPS (vector)
	encoding_FCVTZS_asisdmisc_R  // FCVTZS (vector, integer)
	encoding_FRECPE_asisdmisc_R  // FRECPE
	encoding_FRECPX_asisdmisc_R  // FRECPX
	encoding_USQADD_asisdmisc_R  // USQADD
	encoding_SQNEG_asisdmisc_R   // SQNEG
	encoding_CMGE_asisdmisc_Z    // CMGE (zero)
	encoding_CMLE_asisdmisc_Z    // CMLE (zero)
	encoding_NEG_asisdmisc_R     // NEG (vector)
	encoding_SQXTUN_asisdmisc_N  // SQXTUN, SQXTUN2
	encoding_UQXTN_asisdmisc_N   // UQXTN, UQXTN2
	encoding_FCVTXN_asisdmisc_N  // FCVTXN, FCVTXN2
	encoding_FCVTNU_asisdmisc_R  // FCVTNU (vector)
	encoding_FCVTMU_asisdmisc_R  // FCVTMU (vector)
	encoding_FCVTAU_asisdmisc_R  // FCVTAU (vector)
	encoding_UCVTF_asisdmisc_R   // UCVTF (vector, integer)
	encoding_FCMGE_asisdmisc_FZ  // FCMGE (zero)
	encoding_FCMLE_asisdmisc_FZ  // FCMLE (zero)
	encoding_FCVTPU_asisdmisc_R  // FCVTPU (vector)
	encoding_FCVTZU_asisdmisc_R  // FCVTZU (vector, integer)
	encoding_FRSQRTE_asisdmisc_R // FRSQRTE

	// iclass_asisdpair
	encoding_ADDP_asisdpair_only       // ADDP (scalar)
	encoding_FMAXNMP_asisdpair_only_H  // FMAXNMP (scalar)
	encoding_FADDP_asisdpair_only_H    // FADDP (scalar)
	encoding_FMAXP_asisdpair_only_H    // FMAXP (scalar)
	encoding_FMINNMP_asisdpair_only_H  // FMINNMP (scalar)
	encoding_FMINP_asisdpair_only_H    // FMINP (scalar)
	encoding_FMAXNMP_asisdpair_only_SD // FMAXNMP (scalar)
	encoding_FADDP_asisdpair_only_SD   // FADDP (scalar)
	encoding_FMAXP_asisdpair_only_SD   // FMAXP (scalar)
	encoding_FMINNMP_asisdpair_only_SD // FMINNMP (scalar)
	encoding_FMINP_asisdpair_only_SD   // FMINP (scalar)

	// iclass_asisddiff
	encoding_SQDMLAL_asisddiff_only // SQDMLAL, SQDMLAL2 (vector)
	encoding_SQDMLSL_asisddiff_only // SQDMLSL, SQDMLSL2 (vector)
	encoding_SQDMULL_asisddiff_only // SQDMULL, SQDMULL2 (vector)

	// iclass_asisdsame
	encoding_SQADD_asisdsame_only    // SQADD
	encoding_SQSUB_asisdsame_only    // SQSUB
	encoding_CMGT_asisdsame_only     // CMGT (register)
	encoding_CMGE_asisdsame_only     // CMGE (register)
	encoding_SSHL_asisdsame_only     // SSHL
	encoding_SQSHL_asisdsame_only    // SQSHL (register)
	encoding_SRSHL_asisdsame_only    // SRSHL
	encoding_SQRSHL_asisdsame_only   // SQRSHL
	encoding_ADD_asisdsame_only      // ADD (vector)
	encoding_CMTST_asisdsame_only    // CMTST
	encoding_SQDMULH_asisdsame_only  // SQDMULH (vector)
	encoding_FMULX_asisdsame_only    // FMULX
	encoding_FCMEQ_asisdsame_only    // FCMEQ (register)
	encoding_FRECPS_asisdsame_only   // FRECPS
	encoding_FRSQRTS_asisdsame_only  // FRSQRTS
	encoding_UQADD_asisdsame_only    // UQADD
	encoding_UQSUB_asisdsame_only    // UQSUB
	encoding_CMHI_asisdsame_only     // CMHI (register)
	encoding_CMHS_asisdsame_only     // CMHS (register)
	encoding_USHL_asisdsame_only     // USHL
	encoding_UQSHL_asisdsame_only    // UQSHL (register)
	encoding_URSHL_asisdsame_only    // URSHL
	encoding_UQRSHL_asisdsame_only   // UQRSHL
	encoding_SUB_asisdsame_only      // SUB (vector)
	encoding_CMEQ_asisdsame_only     // CMEQ (register)
	encoding_SQRDMULH_asisdsame_only // SQRDMULH (vector)
	encoding_FCMGE_asisdsame_only    // FCMGE (register)
	encoding_FACGE_asisdsame_only    // FACGE
	encoding_FABD_asisdsame_only     // FABD
	encoding_FCMGT_asisdsame_only    // FCMGT (register)
	encoding_FACGT_asisdsame_only    // FACGT

	// iclass_asisdshf
	encoding_SSHR_asisdshf_R     // SSHR
	encoding_SSRA_asisdshf_R     // SSRA
	encoding_SRSHR_asisdshf_R    // SRSHR
	encoding_SRSRA_asisdshf_R    // SRSRA
	encoding_SHL_asisdshf_R      // SHL
	encoding_SQSHL_asisdshf_R    // SQSHL (immediate)
	encoding_SQSHRN_asisdshf_N   // SQSHRN, SQSHRN2
	encoding_SQRSHRN_asisdshf_N  // SQRSHRN, SQRSHRN2
	encoding_SCVTF_asisdshf_C    // SCVTF (vector, fixed-point)
	encoding_FCVTZS_asisdshf_C   // FCVTZS (vector, fixed-point)
	encoding_USHR_asisdshf_R     // USHR
	encoding_USRA_asisdshf_R     // USRA
	encoding_URSHR_asisdshf_R    // URSHR
	encoding_URSRA_asisdshf_R    // URSRA
	encoding_SRI_asisdshf_R      // SRI
	encoding_SLI_asisdshf_R      // SLI
	encoding_SQSHLU_asisdshf_R   // SQSHLU
	encoding_UQSHL_asisdshf_R    // UQSHL (immediate)
	encoding_SQSHRUN_asisdshf_N  // SQSHRUN, SQSHRUN2
	encoding_SQRSHRUN_asisdshf_N // SQRSHRUN, SQRSHRUN2
	encoding_UQSHRN_asisdshf_N   // UQSHRN, UQSHRN2
	encoding_UQRSHRN_asisdshf_N  // UQRSHRN, UQRSHRN2
	encoding_UCVTF_asisdshf_C    // UCVTF (vector, fixed-point)
	encoding_FCVTZU_asisdshf_C   // FCVTZU (vector, fixed-point)

	// iclass_asisdelem
	encoding_SQDMLAL_asisdelem_L  // SQDMLAL, SQDMLAL2 (by element)
	encoding_SQDMLSL_asisdelem_L  // SQDMLSL, SQDMLSL2 (by element)
	encoding_SQDMULL_asisdelem_L  // SQDMULL, SQDMULL2 (by element)
	encoding_SQDMULH_asisdelem_R  // SQDMULH (by element)
	encoding_SQRDMULH_asisdelem_R // SQRDMULH (by element)
	encoding_FMLA_asisdelem_RH_H  // FMLA (by element)
	encoding_FMLS_asisdelem_RH_H  // FMLS (by element)
	encoding_FMUL_asisdelem_RH_H  // FMUL (by element)
	encoding_FMLA_asisdelem_R_SD  // FMLA (by element)
	encoding_FMLS_asisdelem_R_SD  // FMLS (by element)
	encoding_FMUL_asisdelem_R_SD  // FMUL (by element)
	encoding_SQRDMLAH_asisdelem_R // SQRDMLAH (by element)
	encoding_SQRDMLSH_asisdelem_R // SQRDMLSH (by element)
	encoding_FMULX_asisdelem_RH_H // FMULX (by element)
	encoding_FMULX_asisdelem_R_SD // FMULX (by element)

	// iclass_asimdtbl
	encoding_TBL_asimdtbl_L1_1 // TBL
	encoding_TBX_asimdtbl_L1_1 // TBX
	encoding_TBL_asimdtbl_L2_2 // TBL
	encoding_TBX_asimdtbl_L2_2 // TBX
	encoding_TBL_asimdtbl_L3_3 // TBL
	encoding_TBX_asimdtbl_L3_3 // TBX
	encoding_TBL_asimdtbl_L4_4 // TBL
	encoding_TBX_asimdtbl_L4_4 // TBX

	// iclass_asimdperm
	encoding_UZP1_asimdperm_only // UZP1
	encoding_TRN1_asimdperm_only // TRN1
	encoding_ZIP1_asimdperm_only // ZIP1
	encoding_UZP2_asimdperm_only // UZP2
	encoding_TRN2_asimdperm_only // TRN2
	encoding_ZIP2_asimdperm_only // ZIP2

	// iclass_asimdext
	encoding_EXT_asimdext_only // EXT

	// iclass_asimdins
	encoding_DUP_asimdins_DV_v // DUP (element)
	encoding_DUP_asimdins_DR_r // DUP (general)
	encoding_SMOV_asimdins_W_w // SMOV
	encoding_UMOV_asimdins_W_w // UMOV
	encoding_INS_asimdins_IR_r // INS (general)
	encoding_SMOV_asimdins_X_x // SMOV
	encoding_UMOV_asimdins_X_x // UMOV
	encoding_INS_asimdins_IV_v // INS (element)

	// iclass_asimdsamefp16
	encoding_FMAXNM_asimdsamefp16_only  // FMAXNM (vector)
	encoding_FMLA_asimdsamefp16_only    // FMLA (vector)
	encoding_FADD_asimdsamefp16_only    // FADD (vector)
	encoding_FMULX_asimdsamefp16_only   // FMULX
	encoding_FCMEQ_asimdsamefp16_only   // FCMEQ (register)
	encoding_FMAX_asimdsamefp16_only    // FMAX (vector)
	encoding_FRECPS_asimdsamefp16_only  // FRECPS
	encoding_FMINNM_asimdsamefp16_only  // FMINNM (vector)
	encoding_FMLS_asimdsamefp16_only    // FMLS (vector)
	encoding_FSUB_asimdsamefp16_only    // FSUB (vector)
	encoding_FMIN_asimdsamefp16_only    // FMIN (vector)
	encoding_FRSQRTS_asimdsamefp16_only // FRSQRTS
	encoding_FMAXNMP_asimdsamefp16_only // FMAXNMP (vector)
	encoding_FADDP_asimdsamefp16_only   // FADDP (vector)
	encoding_FMUL_asimdsamefp16_only    // FMUL (vector)
	encoding_FCMGE_asimdsamefp16_only   // FCMGE (register)
	encoding_FACGE_asimdsamefp16_only   // FACGE
	encoding_FMAXP_asimdsamefp16_only   // FMAXP (vector)
	encoding_FDIV_asimdsamefp16_only    // FDIV (vector)
	encoding_FMINNMP_asimdsamefp16_only // FMINNMP (vector)
	encoding_FABD_asimdsamefp16_only    // FABD
	encoding_FCMGT_asimdsamefp16_only   // FCMGT (register)
	encoding_FACGT_asimdsamefp16_only   // FACGT
	encoding_FMINP_asimdsamefp16_only   // FMINP (vector)

	// iclass_asimdmiscfp16
	encoding_FRINTN_asimdmiscfp16_R  // FRINTN (vector)
	encoding_FRINTM_asimdmiscfp16_R  // FRINTM (vector)
	encoding_FCVTNS_asimdmiscfp16_R  // FCVTNS (vector)
	encoding_FCVTMS_asimdmiscfp16_R  // FCVTMS (vector)
	encoding_FCVTAS_asimdmiscfp16_R  // FCVTAS (vector)
	encoding_SCVTF_asimdmiscfp16_R   // SCVTF (vector, integer)
	encoding_FCMGT_asimdmiscfp16_FZ  // FCMGT (zero)
	encoding_FCMEQ_asimdmiscfp16_FZ  // FCMEQ (zero)
	encoding_FCMLT_asimdmiscfp16_FZ  // FCMLT (zero)
	encoding_FABS_asimdmiscfp16_R    // FABS (vector)
	encoding_FRINTP_asimdmiscfp16_R  // FRINTP (vector)
	encoding_FRINTZ_asimdmiscfp16_R  // FRINTZ (vector)
	encoding_FCVTPS_asimdmiscfp16_R  // FCVTPS (vector)
	encoding_FCVTZS_asimdmiscfp16_R  // FCVTZS (vector, integer)
	encoding_FRECPE_asimdmiscfp16_R  // FRECPE
	encoding_FRINTA_asimdmiscfp16_R  // FRINTA (vector)
	encoding_FRINTX_asimdmiscfp16_R  // FRINTX (vector)
	encoding_FCVTNU_asimdmiscfp16_R  // FCVTNU (vector)
	encoding_FCVTMU_asimdmiscfp16_R  // FCVTMU (vector)
	encoding_FCVTAU_asimdmiscfp16_R  // FCVTAU (vector)
	encoding_UCVTF_asimdmiscfp16_R   // UCVTF (vector, integer)
	encoding_FCMGE_asimdmiscfp16_FZ  // FCMGE (zero)
	encoding_FCMLE_asimdmiscfp16_FZ  // FCMLE (zero)
	encoding_FNEG_asimdmiscfp16_R    // FNEG (vector)
	encoding_FRINTI_asimdmiscfp16_R  // FRINTI (vector)
	encoding_FCVTPU_asimdmiscfp16_R  // FCVTPU (vector)
	encoding_FCVTZU_asimdmiscfp16_R  // FCVTZU (vector, integer)
	encoding_FRSQRTE_asimdmiscfp16_R // FRSQRTE
	encoding_FSQRT_asimdmiscfp16_R   // FSQRT (vector)

	// iclass_asimdsame2
	encoding_SDOT_asimdsame2_D        // SDOT (vector)
	encoding_USDOT_asimdsame2_D       // USDOT (vector)
	encoding_SQRDMLAH_asimdsame2_only // SQRDMLAH (vector)
	encoding_SQRDMLSH_asimdsame2_only // SQRDMLSH (vector)
	encoding_UDOT_asimdsame2_D        // UDOT (vector)
	encoding_FCMLA_asimdsame2_C       // FCMLA
	encoding_FCADD_asimdsame2_C       // FCADD
	encoding_BFDOT_asimdsame2_D       // BFDOT (vector)
	encoding_BFMLAL_asimdsame2_F_     // BFMLALB, BFMLALT (vector)
	encoding_SMMLA_asimdsame2_G       // SMMLA (vector)
	encoding_USMMLA_asimdsame2_G      // USMMLA (vector)
	encoding_BFMMLA_asimdsame2_E      // BFMMLA
	encoding_UMMLA_asimdsame2_G       // UMMLA (vector)

	// iclass_asimdmisc
	encoding_REV64_asimdmisc_R    // REV64
	encoding_REV16_asimdmisc_R    // REV16 (vector)
	encoding_SADDLP_asimdmisc_P   // SADDLP
	encoding_SUQADD_asimdmisc_R   // SUQADD
	encoding_CLS_asimdmisc_R      // CLS (vector)
	encoding_CNT_asimdmisc_R      // CNT
	encoding_SADALP_asimdmisc_P   // SADALP
	encoding_SQABS_asimdmisc_R    // SQABS
	encoding_CMGT_asimdmisc_Z     // CMGT (zero)
	encoding_CMEQ_asimdmisc_Z     // CMEQ (zero)
	encoding_CMLT_asimdmisc_Z     // CMLT (zero)
	encoding_ABS_asimdmisc_R      // ABS
	encoding_XTN_asimdmisc_N      // XTN, XTN2
	encoding_SQXTN_asimdmisc_N    // SQXTN, SQXTN2
	encoding_FCVTN_asimdmisc_N    // FCVTN, FCVTN2
	encoding_FCVTL_asimdmisc_L    // FCVTL, FCVTL2
	encoding_FRINTN_asimdmisc_R   // FRINTN (vector)
	encoding_FRINTM_asimdmisc_R   // FRINTM (vector)
	encoding_FCVTNS_asimdmisc_R   // FCVTNS (vector)
	encoding_FCVTMS_asimdmisc_R   // FCVTMS (vector)
	encoding_FCVTAS_asimdmisc_R   // FCVTAS (vector)
	encoding_SCVTF_asimdmisc_R    // SCVTF (vector, integer)
	encoding_FRINT32Z_asimdmisc_R // FRINT32Z (vector)
	encoding_FRINT64Z_asimdmisc_R // FRINT64Z (vector)
	encoding_FCMGT_asimdmisc_FZ   // FCMGT (zero)
	encoding_FCMEQ_asimdmisc_FZ   // FCMEQ (zero)
	encoding_FCMLT_asimdmisc_FZ   // FCMLT (zero)
	encoding_FABS_asimdmisc_R     // FABS (vector)
	encoding_FRINTP_asimdmisc_R   // FRINTP (vector)
	encoding_FRINTZ_asimdmisc_R   // FRINTZ (vector)
	encoding_FCVTPS_asimdmisc_R   // FCVTPS (vector)
	encoding_FCVTZS_asimdmisc_R   // FCVTZS (vector, integer)
	encoding_URECPE_asimdmisc_R   // URECPE
	encoding_FRECPE_asimdmisc_R   // FRECPE
	encoding_BFCVTN_asimdmisc_4S  // BFCVTN, BFCVTN2
	encoding_REV32_asimdmisc_R    // REV32 (vector)
	encoding_UADDLP_asimdmisc_P   // UADDLP
	encoding_USQADD_asimdmisc_R   // USQADD
	encoding_CLZ_asimdmisc_R      // CLZ (vector)
	encoding_UADALP_asimdmisc_P   // UADALP
	encoding_SQNEG_asimdmisc_R    // SQNEG
	encoding_CMGE_asimdmisc_Z     // CMGE (zero)
	encoding_CMLE_asimdmisc_Z     // CMLE (zero)
	encoding_NEG_asimdmisc_R      // NEG (vector)
	encoding_SQXTUN_asimdmisc_N   // SQXTUN, SQXTUN2
	encoding_SHLL_asimdmisc_S     // SHLL, SHLL2
	encoding_UQXTN_asimdmisc_N    // UQXTN, UQXTN2
	encoding_FCVTXN_asimdmisc_N   // FCVTXN, FCVTXN2
	encoding_FRINTA_asimdmisc_R   // FRINTA (vector)
	encoding_FRINTX_asimdmisc_R   // FRINTX (vector)
	encoding_FCVTNU_asimdmisc_R   // FCVTNU (vector)
	encoding_FCVTMU_asimdmisc_R   // FCVTMU (vector)
	encoding_FCVTAU_asimdmisc_R   // FCVTAU (vector)
	encoding_UCVTF_asimdmisc_R    // UCVTF (vector, integer)
	encoding_FRINT32X_asimdmisc_R // FRINT32X (vector)
	encoding_FRINT64X_asimdmisc_R // FRINT64X (vector)
	encoding_NOT_asimdmisc_R      // NOT
	encoding_RBIT_asimdmisc_R     // RBIT (vector)
	encoding_FCMGE_asimdmisc_FZ   // FCMGE (zero)
	encoding_FCMLE_asimdmisc_FZ   // FCMLE (zero)
	encoding_FNEG_asimdmisc_R     // FNEG (vector)
	encoding_FRINTI_asimdmisc_R   // FRINTI (vector)
	encoding_FCVTPU_asimdmisc_R   // FCVTPU (vector)
	encoding_FCVTZU_asimdmisc_R   // FCVTZU (vector, integer)
	encoding_URSQRTE_asimdmisc_R  // URSQRTE
	encoding_FRSQRTE_asimdmisc_R  // FRSQRTE
	encoding_FSQRT_asimdmisc_R    // FSQRT (vector)

	// iclass_asimdall
	encoding_SADDLV_asimdall_only     // SADDLV
	encoding_SMAXV_asimdall_only      // SMAXV
	encoding_SMINV_asimdall_only      // SMINV
	encoding_ADDV_asimdall_only       // ADDV
	encoding_FMAXNMV_asimdall_only_H  // FMAXNMV
	encoding_FMAXV_asimdall_only_H    // FMAXV
	encoding_FMINNMV_asimdall_only_H  // FMINNMV
	encoding_FMINV_asimdall_only_H    // FMINV
	encoding_UADDLV_asimdall_only     // UADDLV
	encoding_UMAXV_asimdall_only      // UMAXV
	encoding_UMINV_asimdall_only      // UMINV
	encoding_FMAXNMV_asimdall_only_SD // FMAXNMV
	encoding_FMAXV_asimdall_only_SD   // FMAXV
	encoding_FMINNMV_asimdall_only_SD // FMINNMV
	encoding_FMINV_asimdall_only_SD   // FMINV

	// iclass_asimddiff
	encoding_SADDL_asimddiff_L   // SADDL, SADDL2
	encoding_SADDW_asimddiff_W   // SADDW, SADDW2
	encoding_SSUBL_asimddiff_L   // SSUBL, SSUBL2
	encoding_SSUBW_asimddiff_W   // SSUBW, SSUBW2
	encoding_ADDHN_asimddiff_N   // ADDHN, ADDHN2
	encoding_SABAL_asimddiff_L   // SABAL, SABAL2
	encoding_SUBHN_asimddiff_N   // SUBHN, SUBHN2
	encoding_SABDL_asimddiff_L   // SABDL, SABDL2
	encoding_SMLAL_asimddiff_L   // SMLAL, SMLAL2 (vector)
	encoding_SQDMLAL_asimddiff_L // SQDMLAL, SQDMLAL2 (vector)
	encoding_SMLSL_asimddiff_L   // SMLSL, SMLSL2 (vector)
	encoding_SQDMLSL_asimddiff_L // SQDMLSL, SQDMLSL2 (vector)
	encoding_SMULL_asimddiff_L   // SMULL, SMULL2 (vector)
	encoding_SQDMULL_asimddiff_L // SQDMULL, SQDMULL2 (vector)
	encoding_PMULL_asimddiff_L   // PMULL, PMULL2
	encoding_UADDL_asimddiff_L   // UADDL, UADDL2
	encoding_UADDW_asimddiff_W   // UADDW, UADDW2
	encoding_USUBL_asimddiff_L   // USUBL, USUBL2
	encoding_USUBW_asimddiff_W   // USUBW, USUBW2
	encoding_RADDHN_asimddiff_N  // RADDHN, RADDHN2
	encoding_UABAL_asimddiff_L   // UABAL, UABAL2
	encoding_RSUBHN_asimddiff_N  // RSUBHN, RSUBHN2
	encoding_UABDL_asimddiff_L   // UABDL, UABDL2
	encoding_UMLAL_asimddiff_L   // UMLAL, UMLAL2 (vector)
	encoding_UMLSL_asimddiff_L   // UMLSL, UMLSL2 (vector)
	encoding_UMULL_asimddiff_L   // UMULL, UMULL2 (vector)

	// iclass_asimdsame
	encoding_SHADD_asimdsame_only    // SHADD
	encoding_SQADD_asimdsame_only    // SQADD
	encoding_SRHADD_asimdsame_only   // SRHADD
	encoding_SHSUB_asimdsame_only    // SHSUB
	encoding_SQSUB_asimdsame_only    // SQSUB
	encoding_CMGT_asimdsame_only     // CMGT (register)
	encoding_CMGE_asimdsame_only     // CMGE (register)
	encoding_SSHL_asimdsame_only     // SSHL
	encoding_SQSHL_asimdsame_only    // SQSHL (register)
	encoding_SRSHL_asimdsame_only    // SRSHL
	encoding_SQRSHL_asimdsame_only   // SQRSHL
	encoding_SMAX_asimdsame_only     // SMAX
	encoding_SMIN_asimdsame_only     // SMIN
	encoding_SABD_asimdsame_only     // SABD
	encoding_SABA_asimdsame_only     // SABA
	encoding_ADD_asimdsame_only      // ADD (vector)
	encoding_CMTST_asimdsame_only    // CMTST
	encoding_MLA_asimdsame_only      // MLA (vector)
	encoding_MUL_asimdsame_only      // MUL (vector)
	encoding_SMAXP_asimdsame_only    // SMAXP
	encoding_SMINP_asimdsame_only    // SMINP
	encoding_SQDMULH_asimdsame_only  // SQDMULH (vector)
	encoding_ADDP_asimdsame_only     // ADDP (vector)
	encoding_FMAXNM_asimdsame_only   // FMAXNM (vector)
	encoding_FMLA_asimdsame_only     // FMLA (vector)
	encoding_FADD_asimdsame_only     // FADD (vector)
	encoding_FMULX_asimdsame_only    // FMULX
	encoding_FCMEQ_asimdsame_only    // FCMEQ (register)
	encoding_FMAX_asimdsame_only     // FMAX (vector)
	encoding_FRECPS_asimdsame_only   // FRECPS
	encoding_AND_asimdsame_only      // AND (vector)
	encoding_FMLAL_asimdsame_F       // FMLAL, FMLAL2 (vector)
	encoding_BIC_asimdsame_only      // BIC (vector, register)
	encoding_FMINNM_asimdsame_only   // FMINNM (vector)
	encoding_FMLS_asimdsame_only     // FMLS (vector)
	encoding_FSUB_asimdsame_only     // FSUB (vector)
	encoding_FMIN_asimdsame_only     // FMIN (vector)
	encoding_FRSQRTS_asimdsame_only  // FRSQRTS
	encoding_ORR_asimdsame_only      // ORR (vector, register)
	encoding_FMLSL_asimdsame_F       // FMLSL, FMLSL2 (vector)
	encoding_ORN_asimdsame_only      // ORN (vector)
	encoding_UHADD_asimdsame_only    // UHADD
	encoding_UQADD_asimdsame_only    // UQADD
	encoding_URHADD_asimdsame_only   // URHADD
	encoding_UHSUB_asimdsame_only    // UHSUB
	encoding_UQSUB_asimdsame_only    // UQSUB
	encoding_CMHI_asimdsame_only     // CMHI (register)
	encoding_CMHS_asimdsame_only     // CMHS (register)
	encoding_USHL_asimdsame_only     // USHL
	encoding_UQSHL_asimdsame_only    // UQSHL (register)
	encoding_URSHL_asimdsame_only    // URSHL
	encoding_UQRSHL_asimdsame_only   // UQRSHL
	encoding_UMAX_asimdsame_only     // UMAX
	encoding_UMIN_asimdsame_only     // UMIN
	encoding_UABD_asimdsame_only     // UABD
	encoding_UABA_asimdsame_only     // UABA
	encoding_SUB_asimdsame_only      // SUB (vector)
	encoding_CMEQ_asimdsame_only     // CMEQ (register)
	encoding_MLS_asimdsame_only      // MLS (vector)
	encoding_PMUL_asimdsame_only     // PMUL
	encoding_UMAXP_asimdsame_only    // UMAXP
	encoding_UMINP_asimdsame_only    // UMINP
	encoding_SQRDMULH_asimdsame_only // SQRDMULH (vector)
	encoding_FMAXNMP_asimdsame_only  // FMAXNMP (vector)
	encoding_FADDP_asimdsame_only    // FADDP (vector)
	encoding_FMUL_asimdsame_only     // FMUL (vector)
	encoding_FCMGE_asimdsame_only    // FCMGE (register)
	encoding_FACGE_asimdsame_only    // FACGE
	encoding_FMAXP_asimdsame_only    // FMAXP (vector)
	encoding_FDIV_asimdsame_only     // FDIV (vector)
	encoding_EOR_asimdsame_only      // EOR (vector)
	encoding_FMLAL2_asimdsame_F      // FMLAL, FMLAL2 (vector)
	encoding_BSL_asimdsame_only      // BSL
	encoding_FMINNMP_asimdsame_only  // FMINNMP (vector)
	encoding_FABD_asimdsame_only     // FABD
	encoding_FCMGT_asimdsame_only    // FCMGT (register)
	encoding_FACGT_asimdsame_only    // FACGT
	encoding_FMINP_asimdsame_only    // FMINP (vector)
	encoding_BIT_asimdsame_only      // BIT
	encoding_FMLSL2_asimdsame_F      // FMLSL, FMLSL2 (vector)
	encoding_BIF_asimdsame_only      // BIF

	// iclass_asimdimm
	encoding_MOVI_asimdimm_L_sl // MOVI
	encoding_ORR_asimdimm_L_sl  // ORR (vector, immediate)
	encoding_MOVI_asimdimm_L_hl // MOVI
	encoding_ORR_asimdimm_L_hl  // ORR (vector, immediate)
	encoding_MOVI_asimdimm_M_sm // MOVI
	encoding_MOVI_asimdimm_N_b  // MOVI
	encoding_FMOV_asimdimm_S_s  // FMOV (vector, immediate)
	encoding_FMOV_asimdimm_H_h  // FMOV (vector, immediate)
	encoding_MVNI_asimdimm_L_sl // MVNI
	encoding_BIC_asimdimm_L_sl  // BIC (vector, immediate)
	encoding_MVNI_asimdimm_L_hl // MVNI
	encoding_BIC_asimdimm_L_hl  // BIC (vector, immediate)
	encoding_MVNI_asimdimm_M_sm // MVNI
	encoding_MOVI_asimdimm_D_ds // MOVI
	encoding_MOVI_asimdimm_D2_d // MOVI
	encoding_FMOV_asimdimm_D2_d // FMOV (vector, immediate)

	// iclass_asimdshf
	encoding_SSHR_asimdshf_R     // SSHR
	encoding_SSRA_asimdshf_R     // SSRA
	encoding_SRSHR_asimdshf_R    // SRSHR
	encoding_SRSRA_asimdshf_R    // SRSRA
	encoding_SHL_asimdshf_R      // SHL
	encoding_SQSHL_asimdshf_R    // SQSHL (immediate)
	encoding_SHRN_asimdshf_N     // SHRN, SHRN2
	encoding_RSHRN_asimdshf_N    // RSHRN, RSHRN2
	encoding_SQSHRN_asimdshf_N   // SQSHRN, SQSHRN2
	encoding_SQRSHRN_asimdshf_N  // SQRSHRN, SQRSHRN2
	encoding_SSHLL_asimdshf_L    // SSHLL, SSHLL2
	encoding_SCVTF_asimdshf_C    // SCVTF (vector, fixed-point)
	encoding_FCVTZS_asimdshf_C   // FCVTZS (vector, fixed-point)
	encoding_USHR_asimdshf_R     // USHR
	encoding_USRA_asimdshf_R     // USRA
	encoding_URSHR_asimdshf_R    // URSHR
	encoding_URSRA_asimdshf_R    // URSRA
	encoding_SRI_asimdshf_R      // SRI
	encoding_SLI_asimdshf_R      // SLI
	encoding_SQSHLU_asimdshf_R   // SQSHLU
	encoding_UQSHL_asimdshf_R    // UQSHL (immediate)
	encoding_SQSHRUN_asimdshf_N  // SQSHRUN, SQSHRUN2
	encoding_SQRSHRUN_asimdshf_N // SQRSHRUN, SQRSHRUN2
	encoding_UQSHRN_asimdshf_N   // UQSHRN, UQSHRN2
	encoding_UQRSHRN_asimdshf_N  // UQRSHRN, UQRSHRN2
	encoding_USHLL_asimdshf_L    // USHLL, USHLL2
	encoding_UCVTF_asimdshf_C    // UCVTF (vector, fixed-point)
	encoding_FCVTZU_asimdshf_C   // FCVTZU (vector, fixed-point)

	// iclass_asimdelem
	encoding_SMLAL_asimdelem_L    // SMLAL, SMLAL2 (by element)
	encoding_SQDMLAL_asimdelem_L  // SQDMLAL, SQDMLAL2 (by element)
	encoding_SMLSL_asimdelem_L    // SMLSL, SMLSL2 (by element)
	encoding_SQDMLSL_asimdelem_L  // SQDMLSL, SQDMLSL2 (by element)
	encoding_MUL_asimdelem_R      // MUL (by element)
	encoding_SMULL_asimdelem_L    // SMULL, SMULL2 (by element)
	encoding_SQDMULL_asimdelem_L  // SQDMULL, SQDMULL2 (by element)
	encoding_SQDMULH_asimdelem_R  // SQDMULH (by element)
	encoding_SQRDMULH_asimdelem_R // SQRDMULH (by element)
	encoding_SDOT_asimdelem_D     // SDOT (by element)
	encoding_FMLA_asimdelem_RH_H  // FMLA (by element)
	encoding_FMLS_asimdelem_RH_H  // FMLS (by element)
	encoding_FMUL_asimdelem_RH_H  // FMUL (by element)
	encoding_SUDOT_asimdelem_D    // SUDOT (by element)
	encoding_BFDOT_asimdelem_E    // BFDOT (by element)
	encoding_FMLA_asimdelem_R_SD  // FMLA (by element)
	encoding_FMLS_asimdelem_R_SD  // FMLS (by element)
	encoding_FMUL_asimdelem_R_SD  // FMUL (by element)
	encoding_FMLAL_asimdelem_LH   // FMLAL, FMLAL2 (by element)
	encoding_FMLSL_asimdelem_LH   // FMLSL, FMLSL2 (by element)
	encoding_USDOT_asimdelem_D    // USDOT (by element)
	encoding_BFMLAL_asimdelem_F   // BFMLALB, BFMLALT (by element)
	encoding_MLA_asimdelem_R      // MLA (by element)
	encoding_UMLAL_asimdelem_L    // UMLAL, UMLAL2 (by element)
	encoding_MLS_asimdelem_R      // MLS (by element)
	encoding_UMLSL_asimdelem_L    // UMLSL, UMLSL2 (by element)
	encoding_UMULL_asimdelem_L    // UMULL, UMULL2 (by element)
	encoding_SQRDMLAH_asimdelem_R // SQRDMLAH (by element)
	encoding_UDOT_asimdelem_D     // UDOT (by element)
	encoding_SQRDMLSH_asimdelem_R // SQRDMLSH (by element)
	encoding_FMULX_asimdelem_RH_H // FMULX (by element)
	encoding_FCMLA_asimdelem_C_H  // FCMLA (by element)
	encoding_FMULX_asimdelem_R_SD // FMULX (by element)
	encoding_FCMLA_asimdelem_C_S  // FCMLA (by element)
	encoding_FMLAL2_asimdelem_LH  // FMLAL, FMLAL2 (by element)
	encoding_FMLSL2_asimdelem_LH  // FMLSL, FMLSL2 (by element)

	// iclass_crypto3_imm2
	encoding_SM3TT1A_VVV4_crypto3_imm2 // SM3TT1A
	encoding_SM3TT1B_VVV4_crypto3_imm2 // SM3TT1B
	encoding_SM3TT2A_VVV4_crypto3_imm2 // SM3TT2A
	encoding_SM3TT2B_VVV_crypto3_imm2  // SM3TT2B

	// iclass_cryptosha512_3
	encoding_SHA512H_QQV_cryptosha512_3    // SHA512H
	encoding_SHA512H2_QQV_cryptosha512_3   // SHA512H2
	encoding_SHA512SU1_VVV2_cryptosha512_3 // SHA512SU1
	encoding_RAX1_VVV2_cryptosha512_3      // RAX1
	encoding_SM3PARTW1_VVV4_cryptosha512_3 // SM3PARTW1
	encoding_SM3PARTW2_VVV4_cryptosha512_3 // SM3PARTW2
	encoding_SM4EKEY_VVV4_cryptosha512_3   // SM4EKEY

	// iclass_crypto4
	encoding_EOR3_VVV16_crypto4  // EOR3
	encoding_BCAX_VVV16_crypto4  // BCAX
	encoding_SM3SS1_VVV4_crypto4 // SM3SS1

	// iclass_crypto3_imm6
	encoding_XAR_VVV2_crypto3_imm6 // XAR

	// iclass_cryptosha512_2
	encoding_SHA512SU0_VV2_cryptosha512_2 // SHA512SU0
	encoding_SM4E_VV4_cryptosha512_2      // SM4E

	// iclass_float2fix
	encoding_SCVTF_S32_float2fix  // SCVTF (scalar, fixed-point)
	encoding_UCVTF_S32_float2fix  // UCVTF (scalar, fixed-point)
	encoding_FCVTZS_32S_float2fix // FCVTZS (scalar, fixed-point)
	encoding_FCVTZU_32S_float2fix // FCVTZU (scalar, fixed-point)
	encoding_SCVTF_D32_float2fix  // SCVTF (scalar, fixed-point)
	encoding_UCVTF_D32_float2fix  // UCVTF (scalar, fixed-point)
	encoding_FCVTZS_32D_float2fix // FCVTZS (scalar, fixed-point)
	encoding_FCVTZU_32D_float2fix // FCVTZU (scalar, fixed-point)
	encoding_SCVTF_H32_float2fix  // SCVTF (scalar, fixed-point)
	encoding_UCVTF_H32_float2fix  // UCVTF (scalar, fixed-point)
	encoding_FCVTZS_32H_float2fix // FCVTZS (scalar, fixed-point)
	encoding_FCVTZU_32H_float2fix // FCVTZU (scalar, fixed-point)
	encoding_SCVTF_S64_float2fix  // SCVTF (scalar, fixed-point)
	encoding_UCVTF_S64_float2fix  // UCVTF (scalar, fixed-point)
	encoding_FCVTZS_64S_float2fix // FCVTZS (scalar, fixed-point)
	encoding_FCVTZU_64S_float2fix // FCVTZU (scalar, fixed-point)
	encoding_SCVTF_D64_float2fix  // SCVTF (scalar, fixed-point)
	encoding_UCVTF_D64_float2fix  // UCVTF (scalar, fixed-point)
	encoding_FCVTZS_64D_float2fix // FCVTZS (scalar, fixed-point)
	encoding_FCVTZU_64D_float2fix // FCVTZU (scalar, fixed-point)
	encoding_SCVTF_H64_float2fix  // SCVTF (scalar, fixed-point)
	encoding_UCVTF_H64_float2fix  // UCVTF (scalar, fixed-point)
	encoding_FCVTZS_64H_float2fix // FCVTZS (scalar, fixed-point)
	encoding_FCVTZU_64H_float2fix // FCVTZU (scalar, fixed-point)

	// iclass_float2int
	encoding_FCVTNS_32S_float2int  // FCVTNS (scalar)
	encoding_FCVTNU_32S_float2int  // FCVTNU (scalar)
	encoding_SCVTF_S32_float2int   // SCVTF (scalar, integer)
	encoding_UCVTF_S32_float2int   // UCVTF (scalar, integer)
	encoding_FCVTAS_32S_float2int  // FCVTAS (scalar)
	encoding_FCVTAU_32S_float2int  // FCVTAU (scalar)
	encoding_FMOV_32S_float2int    // FMOV (general)
	encoding_FMOV_S32_float2int    // FMOV (general)
	encoding_FCVTPS_32S_float2int  // FCVTPS (scalar)
	encoding_FCVTPU_32S_float2int  // FCVTPU (scalar)
	encoding_FCVTMS_32S_float2int  // FCVTMS (scalar)
	encoding_FCVTMU_32S_float2int  // FCVTMU (scalar)
	encoding_FCVTZS_32S_float2int  // FCVTZS (scalar, integer)
	encoding_FCVTZU_32S_float2int  // FCVTZU (scalar, integer)
	encoding_FCVTNS_32D_float2int  // FCVTNS (scalar)
	encoding_FCVTNU_32D_float2int  // FCVTNU (scalar)
	encoding_SCVTF_D32_float2int   // SCVTF (scalar, integer)
	encoding_UCVTF_D32_float2int   // UCVTF (scalar, integer)
	encoding_FCVTAS_32D_float2int  // FCVTAS (scalar)
	encoding_FCVTAU_32D_float2int  // FCVTAU (scalar)
	encoding_FCVTPS_32D_float2int  // FCVTPS (scalar)
	encoding_FCVTPU_32D_float2int  // FCVTPU (scalar)
	encoding_FCVTMS_32D_float2int  // FCVTMS (scalar)
	encoding_FCVTMU_32D_float2int  // FCVTMU (scalar)
	encoding_FCVTZS_32D_float2int  // FCVTZS (scalar, integer)
	encoding_FCVTZU_32D_float2int  // FCVTZU (scalar, integer)
	encoding_FJCVTZS_32D_float2int // FJCVTZS
	encoding_FCVTNS_32H_float2int  // FCVTNS (scalar)
	encoding_FCVTNU_32H_float2int  // FCVTNU (scalar)
	encoding_SCVTF_H32_float2int   // SCVTF (scalar, integer)
	encoding_UCVTF_H32_float2int   // UCVTF (scalar, integer)
	encoding_FCVTAS_32H_float2int  // FCVTAS (scalar)
	encoding_FCVTAU_32H_float2int  // FCVTAU (scalar)
	encoding_FMOV_32H_float2int    // FMOV (general)
	encoding_FMOV_H32_float2int    // FMOV (general)
	encoding_FCVTPS_32H_float2int  // FCVTPS (scalar)
	encoding_FCVTPU_32H_float2int  // FCVTPU (scalar)
	encoding_FCVTMS_32H_float2int  // FCVTMS (scalar)
	encoding_FCVTMU_32H_float2int  // FCVTMU (scalar)
	encoding_FCVTZS_32H_float2int  // FCVTZS (scalar, integer)
	encoding_FCVTZU_32H_float2int  // FCVTZU (scalar, integer)
	encoding_FCVTNS_64S_float2int  // FCVTNS (scalar)
	encoding_FCVTNU_64S_float2int  // FCVTNU (scalar)
	encoding_SCVTF_S64_float2int   // SCVTF (scalar, integer)
	encoding_UCVTF_S64_float2int   // UCVTF (scalar, integer)
	encoding_FCVTAS_64S_float2int  // FCVTAS (scalar)
	encoding_FCVTAU_64S_float2int  // FCVTAU (scalar)
	encoding_FCVTPS_64S_float2int  // FCVTPS (scalar)
	encoding_FCVTPU_64S_float2int  // FCVTPU (scalar)
	encoding_FCVTMS_64S_float2int  // FCVTMS (scalar)
	encoding_FCVTMU_64S_float2int  // FCVTMU (scalar)
	encoding_FCVTZS_64S_float2int  // FCVTZS (scalar, integer)
	encoding_FCVTZU_64S_float2int  // FCVTZU (scalar, integer)
	encoding_FCVTNS_64D_float2int  // FCVTNS (scalar)
	encoding_FCVTNU_64D_float2int  // FCVTNU (scalar)
	encoding_SCVTF_D64_float2int   // SCVTF (scalar, integer)
	encoding_UCVTF_D64_float2int   // UCVTF (scalar, integer)
	encoding_FCVTAS_64D_float2int  // FCVTAS (scalar)
	encoding_FCVTAU_64D_float2int  // FCVTAU (scalar)
	encoding_FMOV_64D_float2int    // FMOV (general)
	encoding_FMOV_D64_float2int    // FMOV (general)
	encoding_FCVTPS_64D_float2int  // FCVTPS (scalar)
	encoding_FCVTPU_64D_float2int  // FCVTPU (scalar)
	encoding_FCVTMS_64D_float2int  // FCVTMS (scalar)
	encoding_FCVTMU_64D_float2int  // FCVTMU (scalar)
	encoding_FCVTZS_64D_float2int  // FCVTZS (scalar, integer)
	encoding_FCVTZU_64D_float2int  // FCVTZU (scalar, integer)
	encoding_FMOV_64VX_float2int   // FMOV (general)
	encoding_FMOV_V64I_float2int   // FMOV (general)
	encoding_FCVTNS_64H_float2int  // FCVTNS (scalar)
	encoding_FCVTNU_64H_float2int  // FCVTNU (scalar)
	encoding_SCVTF_H64_float2int   // SCVTF (scalar, integer)
	encoding_UCVTF_H64_float2int   // UCVTF (scalar, integer)
	encoding_FCVTAS_64H_float2int  // FCVTAS (scalar)
	encoding_FCVTAU_64H_float2int  // FCVTAU (scalar)
	encoding_FMOV_64H_float2int    // FMOV (general)
	encoding_FMOV_H64_float2int    // FMOV (general)
	encoding_FCVTPS_64H_float2int  // FCVTPS (scalar)
	encoding_FCVTPU_64H_float2int  // FCVTPU (scalar)
	encoding_FCVTMS_64H_float2int  // FCVTMS (scalar)
	encoding_FCVTMU_64H_float2int  // FCVTMU (scalar)
	encoding_FCVTZS_64H_float2int  // FCVTZS (scalar, integer)
	encoding_FCVTZU_64H_float2int  // FCVTZU (scalar, integer)

	// iclass_floatdp1
	encoding_FMOV_S_floatdp1     // FMOV (register)
	encoding_FABS_S_floatdp1     // FABS (scalar)
	encoding_FNEG_S_floatdp1     // FNEG (scalar)
	encoding_FSQRT_S_floatdp1    // FSQRT (scalar)
	encoding_FCVT_DS_floatdp1    // FCVT
	encoding_FCVT_HS_floatdp1    // FCVT
	encoding_FRINTN_S_floatdp1   // FRINTN (scalar)
	encoding_FRINTP_S_floatdp1   // FRINTP (scalar)
	encoding_FRINTM_S_floatdp1   // FRINTM (scalar)
	encoding_FRINTZ_S_floatdp1   // FRINTZ (scalar)
	encoding_FRINTA_S_floatdp1   // FRINTA (scalar)
	encoding_FRINTX_S_floatdp1   // FRINTX (scalar)
	encoding_FRINTI_S_floatdp1   // FRINTI (scalar)
	encoding_FRINT32Z_S_floatdp1 // FRINT32Z (scalar)
	encoding_FRINT32X_S_floatdp1 // FRINT32X (scalar)
	encoding_FRINT64Z_S_floatdp1 // FRINT64Z (scalar)
	encoding_FRINT64X_S_floatdp1 // FRINT64X (scalar)
	encoding_FMOV_D_floatdp1     // FMOV (register)
	encoding_FABS_D_floatdp1     // FABS (scalar)
	encoding_FNEG_D_floatdp1     // FNEG (scalar)
	encoding_FSQRT_D_floatdp1    // FSQRT (scalar)
	encoding_FCVT_SD_floatdp1    // FCVT
	encoding_BFCVT_BS_floatdp1   // BFCVT
	encoding_FCVT_HD_floatdp1    // FCVT
	encoding_FRINTN_D_floatdp1   // FRINTN (scalar)
	encoding_FRINTP_D_floatdp1   // FRINTP (scalar)
	encoding_FRINTM_D_floatdp1   // FRINTM (scalar)
	encoding_FRINTZ_D_floatdp1   // FRINTZ (scalar)
	encoding_FRINTA_D_floatdp1   // FRINTA (scalar)
	encoding_FRINTX_D_floatdp1   // FRINTX (scalar)
	encoding_FRINTI_D_floatdp1   // FRINTI (scalar)
	encoding_FRINT32Z_D_floatdp1 // FRINT32Z (scalar)
	encoding_FRINT32X_D_floatdp1 // FRINT32X (scalar)
	encoding_FRINT64Z_D_floatdp1 // FRINT64Z (scalar)
	encoding_FRINT64X_D_floatdp1 // FRINT64X (scalar)
	encoding_FMOV_H_floatdp1     // FMOV (register)
	encoding_FABS_H_floatdp1     // FABS (scalar)
	encoding_FNEG_H_floatdp1     // FNEG (scalar)
	encoding_FSQRT_H_floatdp1    // FSQRT (scalar)
	encoding_FCVT_SH_floatdp1    // FCVT
	encoding_FCVT_DH_floatdp1    // FCVT
	encoding_FRINTN_H_floatdp1   // FRINTN (scalar)
	encoding_FRINTP_H_floatdp1   // FRINTP (scalar)
	encoding_FRINTM_H_floatdp1   // FRINTM (scalar)
	encoding_FRINTZ_H_floatdp1   // FRINTZ (scalar)
	encoding_FRINTA_H_floatdp1   // FRINTA (scalar)
	encoding_FRINTX_H_floatdp1   // FRINTX (scalar)
	encoding_FRINTI_H_floatdp1   // FRINTI (scalar)

	// iclass_floatcmp
	encoding_FCMP_S_floatcmp   // FCMP
	encoding_FCMP_SZ_floatcmp  // FCMP
	encoding_FCMPE_S_floatcmp  // FCMPE
	encoding_FCMPE_SZ_floatcmp // FCMPE
	encoding_FCMP_D_floatcmp   // FCMP
	encoding_FCMP_DZ_floatcmp  // FCMP
	encoding_FCMPE_D_floatcmp  // FCMPE
	encoding_FCMPE_DZ_floatcmp // FCMPE
	encoding_FCMP_H_floatcmp   // FCMP
	encoding_FCMP_HZ_floatcmp  // FCMP
	encoding_FCMPE_H_floatcmp  // FCMPE
	encoding_FCMPE_HZ_floatcmp // FCMPE

	// iclass_floatimm
	encoding_FMOV_S_floatimm // FMOV (scalar, immediate)
	encoding_FMOV_D_floatimm // FMOV (scalar, immediate)
	encoding_FMOV_H_floatimm // FMOV (scalar, immediate)

	// iclass_floatccmp
	encoding_FCCMP_S_floatccmp  // FCCMP
	encoding_FCCMPE_S_floatccmp // FCCMPE
	encoding_FCCMP_D_floatccmp  // FCCMP
	encoding_FCCMPE_D_floatccmp // FCCMPE
	encoding_FCCMP_H_floatccmp  // FCCMP
	encoding_FCCMPE_H_floatccmp // FCCMPE

	// iclass_floatdp2
	encoding_FMUL_S_floatdp2   // FMUL (scalar)
	encoding_FDIV_S_floatdp2   // FDIV (scalar)
	encoding_FADD_S_floatdp2   // FADD (scalar)
	encoding_FSUB_S_floatdp2   // FSUB (scalar)
	encoding_FMAX_S_floatdp2   // FMAX (scalar)
	encoding_FMIN_S_floatdp2   // FMIN (scalar)
	encoding_FMAXNM_S_floatdp2 // FMAXNM (scalar)
	encoding_FMINNM_S_floatdp2 // FMINNM (scalar)
	encoding_FNMUL_S_floatdp2  // FNMUL (scalar)
	encoding_FMUL_D_floatdp2   // FMUL (scalar)
	encoding_FDIV_D_floatdp2   // FDIV (scalar)
	encoding_FADD_D_floatdp2   // FADD (scalar)
	encoding_FSUB_D_floatdp2   // FSUB (scalar)
	encoding_FMAX_D_floatdp2   // FMAX (scalar)
	encoding_FMIN_D_floatdp2   // FMIN (scalar)
	encoding_FMAXNM_D_floatdp2 // FMAXNM (scalar)
	encoding_FMINNM_D_floatdp2 // FMINNM (scalar)
	encoding_FNMUL_D_floatdp2  // FNMUL (scalar)
	encoding_FMUL_H_floatdp2   // FMUL (scalar)
	encoding_FDIV_H_floatdp2   // FDIV (scalar)
	encoding_FADD_H_floatdp2   // FADD (scalar)
	encoding_FSUB_H_floatdp2   // FSUB (scalar)
	encoding_FMAX_H_floatdp2   // FMAX (scalar)
	encoding_FMIN_H_floatdp2   // FMIN (scalar)
	encoding_FMAXNM_H_floatdp2 // FMAXNM (scalar)
	encoding_FMINNM_H_floatdp2 // FMINNM (scalar)
	encoding_FNMUL_H_floatdp2  // FNMUL (scalar)

	// iclass_floatsel
	encoding_FCSEL_S_floatsel // FCSEL
	encoding_FCSEL_D_floatsel // FCSEL
	encoding_FCSEL_H_floatsel // FCSEL

	// iclass_floatdp3
	encoding_FMADD_S_floatdp3  // FMADD
	encoding_FMSUB_S_floatdp3  // FMSUB
	encoding_FNMADD_S_floatdp3 // FNMADD
	encoding_FNMSUB_S_floatdp3 // FNMSUB
	encoding_FMADD_D_floatdp3  // FMADD
	encoding_FMSUB_D_floatdp3  // FMSUB
	encoding_FNMADD_D_floatdp3 // FNMADD
	encoding_FNMSUB_D_floatdp3 // FNMSUB
	encoding_FMADD_H_floatdp3  // FMADD
	encoding_FMSUB_H_floatdp3  // FMSUB
	encoding_FNMADD_H_floatdp3 // FNMADD
	encoding_FNMSUB_H_floatdp3 // FNMSUB
)
