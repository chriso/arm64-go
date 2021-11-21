// This file was generated. DO NOT EDIT.
package arm64

import "errors"

var errOverflow = errors.New("overflow")

func encode_perm_undef(imm16 uint32) (ins uint32, err error) {
	switch {
	case imm16 > 0xffff:
		err = errOverflow
	default:
		// encoding_UDF_only_perm_undef
	}

	ins |= imm16

	return
}

func encode_mortlach_f32f32_prod(Zm, Pm, Pn, Zn, S, ZAda uint32) (ins uint32, err error) {
	switch {
	case Zm > 0x1f || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || S > 0x1 || ZAda > 0x3:
		err = errOverflow
	case S == 0x0:
		// encoding_fmopa_za_pp_zz_32
	case S == 0x1:
		// encoding_fmops_za_pp_zz_32
	default:
		err = errUnmatched
	}

	ins |= (Zm << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | (S << 4) | ZAda
	ins |= 0x80800000

	return
}

func encode_mortlach_b16f32_prod(Zm, Pm, Pn, Zn, S, ZAda uint32) (ins uint32, err error) {
	switch {
	case Zm > 0x1f || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || S > 0x1 || ZAda > 0x3:
		err = errOverflow
	case S == 0x0:
		// encoding_bfmopa_za32_pp_zz_
	case S == 0x1:
		// encoding_bfmops_za32_pp_zz_
	default:
		err = errUnmatched
	}

	ins |= (Zm << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | (S << 4) | ZAda
	ins |= 0x81800000

	return
}

func encode_mortlach_f16f32_prod(Zm, Pm, Pn, Zn, S, ZAda uint32) (ins uint32, err error) {
	switch {
	case Zm > 0x1f || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || S > 0x1 || ZAda > 0x3:
		err = errOverflow
	case S == 0x0:
		// encoding_fmopa_za32_pp_zz_16
	case S == 0x1:
		// encoding_fmops_za32_pp_zz_16
	default:
		err = errUnmatched
	}

	ins |= (Zm << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | (S << 4) | ZAda
	ins |= 0x81a00000

	return
}

func encode_mortlach_i8i32_prod(u0, u1, Zm, Pm, Pn, Zn, S, ZAda uint32) (ins uint32, err error) {
	switch {
	case u0 > 0x1 || u1 > 0x1 || Zm > 0x1f || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || S > 0x1 || ZAda > 0x3:
		err = errOverflow
	case u0 == 0x0 && u1 == 0x0 && S == 0x0:
		// encoding_smopa_za_pp_zz_32
	case u0 == 0x0 && u1 == 0x0 && S == 0x1:
		// encoding_smops_za_pp_zz_32
	case u0 == 0x0 && u1 == 0x1 && S == 0x0:
		// encoding_sumopa_za_pp_zz_32
	case u0 == 0x0 && u1 == 0x1 && S == 0x1:
		// encoding_sumops_za_pp_zz_32
	case u0 == 0x1 && u1 == 0x0 && S == 0x0:
		// encoding_usmopa_za_pp_zz_32
	case u0 == 0x1 && u1 == 0x0 && S == 0x1:
		// encoding_usmops_za_pp_zz_32
	case u0 == 0x1 && u1 == 0x1 && S == 0x0:
		// encoding_umopa_za_pp_zz_32
	case u0 == 0x1 && u1 == 0x1 && S == 0x1:
		// encoding_umops_za_pp_zz_32
	default:
		err = errUnmatched
	}

	ins |= (u0 << 24) | (u1 << 21) | (Zm << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | (S << 4) | ZAda
	ins |= 0xa0800000

	return
}

func encode_mortlach_f64f64_prod(Zm, Pm, Pn, Zn, S, ZAda uint32) (ins uint32, err error) {
	switch {
	case Zm > 0x1f || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || S > 0x1 || ZAda > 0x7:
		err = errOverflow
	case S == 0x0:
		// encoding_fmopa_za_pp_zz_64
	case S == 0x1:
		// encoding_fmops_za_pp_zz_64
	default:
		err = errUnmatched
	}

	ins |= (Zm << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | (S << 4) | ZAda
	ins |= 0x80c00000

	return
}

func encode_mortlach_i16i64_prod(u0, u1, Zm, Pm, Pn, Zn, S, ZAda uint32) (ins uint32, err error) {
	switch {
	case u0 > 0x1 || u1 > 0x1 || Zm > 0x1f || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || S > 0x1 || ZAda > 0x7:
		err = errOverflow
	case u0 == 0x0 && u1 == 0x0 && S == 0x0:
		// encoding_smopa_za_pp_zz_64
	case u0 == 0x0 && u1 == 0x0 && S == 0x1:
		// encoding_smops_za_pp_zz_64
	case u0 == 0x0 && u1 == 0x1 && S == 0x0:
		// encoding_sumopa_za_pp_zz_64
	case u0 == 0x0 && u1 == 0x1 && S == 0x1:
		// encoding_sumops_za_pp_zz_64
	case u0 == 0x1 && u1 == 0x0 && S == 0x0:
		// encoding_usmopa_za_pp_zz_64
	case u0 == 0x1 && u1 == 0x0 && S == 0x1:
		// encoding_usmops_za_pp_zz_64
	case u0 == 0x1 && u1 == 0x1 && S == 0x0:
		// encoding_umopa_za_pp_zz_64
	case u0 == 0x1 && u1 == 0x1 && S == 0x1:
		// encoding_umops_za_pp_zz_64
	default:
		err = errUnmatched
	}

	ins |= (u0 << 24) | (u1 << 21) | (Zm << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | (S << 4) | ZAda
	ins |= 0xa0c00000

	return
}

func encode_mortlach_insert_pred(size, Q, V, Rs, Pg, Zn, opc uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Q > 0x1 || V > 0x1 || Rs > 0x3 || Pg > 0x7 || Zn > 0x1f || opc > 0xf:
		err = errOverflow
	case (size&0x2) == 0x0 && Q == 0x1:
		err = errUnallocated
	case size == 0x0 && Q == 0x0:
		// encoding_mova_za_p_rz_b
	case size == 0x1 && Q == 0x0:
		// encoding_mova_za_p_rz_h
	case size == 0x2 && Q == 0x0:
		// encoding_mova_za_p_rz_w
	case size == 0x2 && Q == 0x1:
		err = errUnallocated
	case size == 0x3 && Q == 0x0:
		// encoding_mova_za_p_rz_d
	case size == 0x3 && Q == 0x1:
		// encoding_mova_za_p_rz_q
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Q << 16) | (V << 15) | (Rs << 13) | (Pg << 10) | (Zn << 5) | opc
	ins |= 0xc0000000

	return
}

func encode_mortlach_extract_pred(size, Q, V, Rs, Pg, opc, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Q > 0x1 || V > 0x1 || Rs > 0x3 || Pg > 0x7 || opc > 0xf || Zd > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && Q == 0x1:
		err = errUnallocated
	case size == 0x0 && Q == 0x0:
		// encoding_mova_z_p_rza_b
	case size == 0x1 && Q == 0x0:
		// encoding_mova_z_p_rza_h
	case size == 0x2 && Q == 0x0:
		// encoding_mova_z_p_rza_w
	case size == 0x2 && Q == 0x1:
		err = errUnallocated
	case size == 0x3 && Q == 0x0:
		// encoding_mova_z_p_rza_d
	case size == 0x3 && Q == 0x1:
		// encoding_mova_z_p_rza_q
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Q << 16) | (V << 15) | (Rs << 13) | (Pg << 10) | (opc << 5) | Zd
	ins |= 0xc0020000

	return
}

func encode_mortlach_zero(imm8 uint32) (ins uint32, err error) {
	switch {
	case imm8 > 0xff:
		err = errOverflow
	default:
		// encoding_zero_za_i_
	}

	ins |= imm8
	ins |= 0xc0080000

	return
}

func encode_mortlach_addhv(op, V, Pm, Pn, Zn, opc2 uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || V > 0x1 || Pm > 0x7 || Pn > 0x7 || Zn > 0x1f || opc2 > 0x7:
		err = errOverflow
	case op == 0x0 && (opc2&0x4) == 0x4:
		err = errUnallocated
	case op == 0x0 && V == 0x0 && (opc2&0x4) == 0x0:
		// encoding_addha_za_pp_z_32
	case op == 0x0 && V == 0x1 && (opc2&0x4) == 0x0:
		// encoding_addva_za_pp_z_32
	case op == 0x1 && V == 0x0:
		// encoding_addha_za_pp_z_64
	case op == 0x1 && V == 0x1:
		// encoding_addva_za_pp_z_64
	default:
		err = errUnmatched
	}

	ins |= (op << 22) | (V << 16) | (Pm << 13) | (Pn << 10) | (Zn << 5) | opc2
	ins |= 0xc0900000

	return
}

func encode_mortlach_contig_load(msz, Rm, V, Rs, Pg, Rn, opc uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || V > 0x1 || Rs > 0x3 || Pg > 0x7 || Rn > 0x1f || opc > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_ld1b_za_p_rrr_
	case msz == 0x1:
		// encoding_ld1h_za_p_rrr_
	case msz == 0x2:
		// encoding_ld1w_za_p_rrr_
	case msz == 0x3:
		// encoding_ld1d_za_p_rrr_
	default:
		err = errUnmatched
	}

	ins |= (msz << 22) | (Rm << 16) | (V << 15) | (Rs << 13) | (Pg << 10) | (Rn << 5) | opc
	ins |= 0xe0000000

	return
}

func encode_mortlach_contig_store(msz, Rm, V, Rs, Pg, Rn, opc uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || V > 0x1 || Rs > 0x3 || Pg > 0x7 || Rn > 0x1f || opc > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_za_p_rrr_
	case msz == 0x1:
		// encoding_st1h_za_p_rrr_
	case msz == 0x2:
		// encoding_st1w_za_p_rrr_
	case msz == 0x3:
		// encoding_st1d_za_p_rrr_
	default:
		err = errUnmatched
	}

	ins |= (msz << 22) | (Rm << 16) | (V << 15) | (Rs << 13) | (Pg << 10) | (Rn << 5) | opc
	ins |= 0xe0200000

	return
}

func encode_mortlach_ctxt_ldst(op, Rv, Rn, imm4 uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || Rv > 0x3 || Rn > 0x1f || imm4 > 0xf:
		err = errOverflow
	case op == 0x0:
		// encoding_ldr_za_ri_
	case op == 0x1:
		// encoding_str_za_ri_
	default:
		err = errUnmatched
	}

	ins |= (op << 21) | (Rv << 13) | (Rn << 5) | imm4
	ins |= 0xe1000000

	return
}

func encode_mortlach_contig_qload(Rm, V, Rs, Pg, Rn, ZAt uint32) (ins uint32, err error) {
	switch {
	case Rm > 0x1f || V > 0x1 || Rs > 0x3 || Pg > 0x7 || Rn > 0x1f || ZAt > 0xf:
		err = errOverflow
	default:
		// encoding_ld1q_za_p_rrr_
	}

	ins |= (Rm << 16) | (V << 15) | (Rs << 13) | (Pg << 10) | (Rn << 5) | ZAt
	ins |= 0xe1c00000

	return
}

func encode_mortlach_contig_qstore(Rm, V, Rs, Pg, Rn, ZAt uint32) (ins uint32, err error) {
	switch {
	case Rm > 0x1f || V > 0x1 || Rs > 0x3 || Pg > 0x7 || Rn > 0x1f || ZAt > 0xf:
		err = errOverflow
	default:
		// encoding_st1q_za_p_rrr_
	}

	ins |= (Rm << 16) | (V << 15) | (Rs << 13) | (Pg << 10) | (Rn << 5) | ZAt
	ins |= 0xe1e00000

	return
}

func encode_sve_int_mlas_vvv_pred(size, Zm, op, Pg, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || Pg > 0x7 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_mla_z_p_zzz_
	case op == 0x1:
		// encoding_mls_z_p_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 13) | (Pg << 10) | (Zn << 5) | Zda
	ins |= 0x4004000

	return
}

func encode_sve_int_mladdsub_vvv_pred(size, Zm, op, Pg, Za, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || Pg > 0x7 || Za > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_mad_z_p_zzz_
	case op == 0x1:
		// encoding_msb_z_p_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 13) | (Pg << 10) | (Za << 5) | Zdn
	ins |= 0x400c000

	return
}

func encode_sve_int_bin_pred_arit_0(size, opc, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_add_z_p_zz_
	case opc == 0x1:
		// encoding_sub_z_p_zz_
	case opc == 0x2:
		err = errUnallocated
	case opc == 0x3:
		// encoding_subr_z_p_zz_
	case (opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4000000

	return
}

func encode_sve_int_bin_pred_arit_1(size, opc, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0 && U == 0x0:
		// encoding_smax_z_p_zz_
	case opc == 0x0 && U == 0x1:
		// encoding_umax_z_p_zz_
	case opc == 0x1 && U == 0x0:
		// encoding_smin_z_p_zz_
	case opc == 0x1 && U == 0x1:
		// encoding_umin_z_p_zz_
	case opc == 0x2 && U == 0x0:
		// encoding_sabd_z_p_zz_
	case opc == 0x2 && U == 0x1:
		// encoding_uabd_z_p_zz_
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4080000

	return
}

func encode_sve_int_bin_pred_arit_2(size, H, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || H > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case H == 0x0 && U == 0x0:
		// encoding_mul_z_p_zz_
	case H == 0x0 && U == 0x1:
		err = errUnallocated
	case H == 0x1 && U == 0x0:
		// encoding_smulh_z_p_zz_
	case H == 0x1 && U == 0x1:
		// encoding_umulh_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (H << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4100000

	return
}

func encode_sve_int_bin_pred_div(size, R, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || R > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case R == 0x0 && U == 0x0:
		// encoding_sdiv_z_p_zz_
	case R == 0x0 && U == 0x1:
		// encoding_udiv_z_p_zz_
	case R == 0x1 && U == 0x0:
		// encoding_sdivr_z_p_zz_
	case R == 0x1 && U == 0x1:
		// encoding_udivr_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (R << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4140000

	return
}

func encode_sve_int_bin_pred_log(size, opc, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_orr_z_p_zz_
	case opc == 0x1:
		// encoding_eor_z_p_zz_
	case opc == 0x2:
		// encoding_and_z_p_zz_
	case opc == 0x3:
		// encoding_bic_z_p_zz_
	case (opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4180000

	return
}

func encode_sve_int_reduce_0(size, op, U, Pg, Zn, Vd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Vd > 0x1f:
		err = errOverflow
	case op == 0x0 && U == 0x0:
		// encoding_saddv_r_p_z_
	case op == 0x0 && U == 0x1:
		// encoding_uaddv_r_p_z_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 17) | (U << 16) | (Pg << 10) | (Zn << 5) | Vd
	ins |= 0x4002000

	return
}

func encode_sve_int_reduce_1(size, op, U, Pg, Zn, Vd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Vd > 0x1f:
		err = errOverflow
	case op == 0x0 && U == 0x0:
		// encoding_smaxv_r_p_z_
	case op == 0x0 && U == 0x1:
		// encoding_umaxv_r_p_z_
	case op == 0x1 && U == 0x0:
		// encoding_sminv_r_p_z_
	case op == 0x1 && U == 0x1:
		// encoding_uminv_r_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 17) | (U << 16) | (Pg << 10) | (Zn << 5) | Vd
	ins |= 0x4082000

	return
}

func encode_sve_int_movprfx_pred(size, opc, M, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || M > 0x1 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_movprfx_z_p_z_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 17) | (M << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x4102000

	return
}

func encode_sve_int_reduce_2(size, opc, Pg, Zn, Vd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || Pg > 0x7 || Zn > 0x1f || Vd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_orv_r_p_z_
	case opc == 0x1:
		// encoding_eorv_r_p_z_
	case opc == 0x2:
		// encoding_andv_r_p_z_
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Vd
	ins |= 0x4182000

	return
}

func encode_sve_int_bin_pred_shift_0(tszh, opc, L, U, Pg, tszl, imm3, Zdn uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x3 || opc > 0x3 || L > 0x1 || U > 0x1 || Pg > 0x7 || tszl > 0x3 || imm3 > 0x7 || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0 && L == 0x0 && U == 0x0:
		// encoding_asr_z_p_zi_
	case opc == 0x0 && L == 0x0 && U == 0x1:
		// encoding_lsr_z_p_zi_
	case opc == 0x0 && L == 0x1 && U == 0x0:
		err = errUnallocated
	case opc == 0x0 && L == 0x1 && U == 0x1:
		// encoding_lsl_z_p_zi_
	case opc == 0x1 && L == 0x0 && U == 0x0:
		// encoding_asrd_z_p_zi_
	case opc == 0x1 && L == 0x0 && U == 0x1:
		err = errUnallocated
	case opc == 0x1 && L == 0x1 && U == 0x0:
		// encoding_sqshl_z_p_zi_
	case opc == 0x1 && L == 0x1 && U == 0x1:
		// encoding_uqshl_z_p_zi_
	case opc == 0x2:
		err = errUnallocated
	case opc == 0x3 && L == 0x0 && U == 0x0:
		// encoding_srshr_z_p_zi_
	case opc == 0x3 && L == 0x0 && U == 0x1:
		// encoding_urshr_z_p_zi_
	case opc == 0x3 && L == 0x1 && U == 0x0:
		err = errUnallocated
	case opc == 0x3 && L == 0x1 && U == 0x1:
		// encoding_sqshlu_z_p_zi_
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (opc << 18) | (L << 17) | (U << 16) | (Pg << 10) | (tszl << 8) | (imm3 << 5) | Zdn
	ins |= 0x4008000

	return
}

func encode_sve_int_bin_pred_shift_1(size, R, L, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || R > 0x1 || L > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case L == 0x1 && U == 0x0:
		err = errUnallocated
	case R == 0x0 && L == 0x0 && U == 0x0:
		// encoding_asr_z_p_zz_
	case R == 0x0 && L == 0x0 && U == 0x1:
		// encoding_lsr_z_p_zz_
	case R == 0x0 && L == 0x1 && U == 0x1:
		// encoding_lsl_z_p_zz_
	case R == 0x1 && L == 0x0 && U == 0x0:
		// encoding_asrr_z_p_zz_
	case R == 0x1 && L == 0x0 && U == 0x1:
		// encoding_lsrr_z_p_zz_
	case R == 0x1 && L == 0x1 && U == 0x1:
		// encoding_lslr_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (R << 18) | (L << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4108000

	return
}

func encode_sve_int_bin_pred_shift_2(size, R, L, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || R > 0x1 || L > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case R == 0x0 && L == 0x0 && U == 0x0:
		// encoding_asr_z_p_zw_
	case R == 0x0 && L == 0x0 && U == 0x1:
		// encoding_lsr_z_p_zw_
	case R == 0x0 && L == 0x1 && U == 0x0:
		err = errUnallocated
	case R == 0x0 && L == 0x1 && U == 0x1:
		// encoding_lsl_z_p_zw_
	case R == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (R << 18) | (L << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4188000

	return
}

func encode_sve_int_un_pred_arit_0(size, opc, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_sxtb_z_p_z_
	case opc == 0x1:
		// encoding_uxtb_z_p_z_
	case opc == 0x2:
		// encoding_sxth_z_p_z_
	case opc == 0x3:
		// encoding_uxth_z_p_z_
	case opc == 0x4:
		// encoding_sxtw_z_p_z_
	case opc == 0x5:
		// encoding_uxtw_z_p_z_
	case opc == 0x6:
		// encoding_abs_z_p_z_
	case opc == 0x7:
		// encoding_neg_z_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x410a000

	return
}

func encode_sve_int_un_pred_arit_1(size, opc, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_cls_z_p_z_
	case opc == 0x1:
		// encoding_clz_z_p_z_
	case opc == 0x2:
		// encoding_cnt_z_p_z_
	case opc == 0x3:
		// encoding_cnot_z_p_z_
	case opc == 0x4:
		// encoding_fabs_z_p_z_
	case opc == 0x5:
		// encoding_fneg_z_p_z_
	case opc == 0x6:
		// encoding_not_z_p_z_
	case opc == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x418a000

	return
}

func encode_sve_int_bin_cons_arit_0(size, Zm, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_add_z_zz_
	case opc == 0x1:
		// encoding_sub_z_zz_
	case (opc & 0x6) == 0x2:
		err = errUnallocated
	case opc == 0x4:
		// encoding_sqadd_z_zz_
	case opc == 0x5:
		// encoding_uqadd_z_zz_
	case opc == 0x6:
		// encoding_sqsub_z_zz_
	case opc == 0x7:
		// encoding_uqsub_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x4200000

	return
}

func encode_sve_int_bin_cons_log(opc, Zm, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || Zm > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_and_z_zz_
	case opc == 0x1:
		// encoding_orr_z_zz_
	case opc == 0x2:
		// encoding_eor_z_zz_
	case opc == 0x3:
		// encoding_bic_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (Zm << 16) | (Zn << 5) | Zd
	ins |= 0x4203000

	return
}

func encode_sve_int_rotate_imm(tszh, tszl, imm3, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x3 || tszl > 0x3 || imm3 > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_xar_z_zzi_
	}

	ins |= (tszh << 22) | (tszl << 19) | (imm3 << 16) | (Zm << 5) | Zdn
	ins |= 0x4203400

	return
}

func encode_sve_int_tern_log(opc, Zm, o2, Zk, Zdn uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || Zm > 0x1f || o2 > 0x1 || Zk > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0 && o2 == 0x0:
		// encoding_eor3_z_zzz_
	case opc == 0x0 && o2 == 0x1:
		// encoding_bsl_z_zzz_
	case opc == 0x1 && o2 == 0x0:
		// encoding_bcax_z_zzz_
	case opc == 0x1 && o2 == 0x1:
		// encoding_bsl1n_z_zzz_
	case (opc&0x2) == 0x2 && o2 == 0x0:
		err = errUnallocated
	case opc == 0x2 && o2 == 0x1:
		// encoding_bsl2n_z_zzz_
	case opc == 0x3 && o2 == 0x1:
		// encoding_nbsl_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (Zm << 16) | (o2 << 10) | (Zk << 5) | Zdn
	ins |= 0x4203800

	return
}

func encode_sve_int_index_ii(size, imm5b, imm5, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm5b > 0x1f || imm5 > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_index_z_ii_
	}

	ins |= (size << 22) | (imm5b << 16) | (imm5 << 5) | Zd
	ins |= 0x4204000

	return
}

func encode_sve_int_index_ri(size, imm5, Rn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm5 > 0x1f || Rn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_index_z_ri_
	}

	ins |= (size << 22) | (imm5 << 16) | (Rn << 5) | Zd
	ins |= 0x4204400

	return
}

func encode_sve_int_index_ir(size, Rm, imm5, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rm > 0x1f || imm5 > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_index_z_ir_
	}

	ins |= (size << 22) | (Rm << 16) | (imm5 << 5) | Zd
	ins |= 0x4204800

	return
}

func encode_sve_int_index_rr(size, Rm, Rn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rm > 0x1f || Rn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_index_z_rr_
	}

	ins |= (size << 22) | (Rm << 16) | (Rn << 5) | Zd
	ins |= 0x4204c00

	return
}

func encode_sve_int_arith_vl(op, Rn, imm6, Rd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || Rn > 0x1f || imm6 > 0x3f || Rd > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_addvl_r_ri_
	case op == 0x1:
		// encoding_addpl_r_ri_
	default:
		err = errUnmatched
	}

	ins |= (op << 22) | (Rn << 16) | (imm6 << 5) | Rd
	ins |= 0x4205000

	return
}

func encode_sve_int_read_vl_a(op, opc2, imm6, Rd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || opc2 > 0x1f || imm6 > 0x3f || Rd > 0x1f:
		err = errOverflow
	case op == 0x0 && (opc2&0x10) == 0x0:
		err = errUnallocated
	case op == 0x0 && (opc2&0x18) == 0x10:
		err = errUnallocated
	case op == 0x0 && (opc2&0x1c) == 0x18:
		err = errUnallocated
	case op == 0x0 && (opc2&0x1e) == 0x1c:
		err = errUnallocated
	case op == 0x0 && opc2 == 0x1e:
		err = errUnallocated
	case op == 0x0 && opc2 == 0x1f:
		// encoding_rdvl_r_i_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 22) | (opc2 << 16) | (imm6 << 5) | Rd
	ins |= 0x4a05000

	return
}

func encode_sve_int_mul_b(size, Zm, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x3 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_mul_z_zz_
	case opc == 0x2:
		// encoding_smulh_z_zz_
	case opc == 0x3:
		// encoding_umulh_z_zz_
	case size == 0x0 && opc == 0x1:
		// encoding_pmul_z_zz_
	case size == 0x1 && opc == 0x1:
		err = errUnallocated
	case (size&0x2) == 0x2 && opc == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x4206000

	return
}

func encode_sve_int_sqdmulh(size, Zm, R, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || R > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case R == 0x0:
		// encoding_sqdmulh_z_zz_
	case R == 0x1:
		// encoding_sqrdmulh_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (R << 10) | (Zn << 5) | Zd
	ins |= 0x4207000

	return
}

func encode_sve_int_bin_cons_shift_a(size, Zm, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x3 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_asr_z_zw_
	case opc == 0x1:
		// encoding_lsr_z_zw_
	case opc == 0x2:
		err = errUnallocated
	case opc == 0x3:
		// encoding_lsl_z_zw_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x4208000

	return
}

func encode_sve_int_bin_cons_shift_b(tszh, tszl, imm3, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x3 || tszl > 0x3 || imm3 > 0x7 || opc > 0x3 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_asr_z_zi_
	case opc == 0x1:
		// encoding_lsr_z_zi_
	case opc == 0x2:
		err = errUnallocated
	case opc == 0x3:
		// encoding_lsl_z_zi_
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (tszl << 19) | (imm3 << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x4209000

	return
}

func encode_sve_int_bin_cons_misc_0_a(opc, Zm, msz, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || Zm > 0x1f || msz > 0x3 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_adr_z_az_d_s32_scaled
	case opc == 0x1:
		// encoding_adr_z_az_d_u32_scaled
	case (opc & 0x2) == 0x2:
		// encoding_adr_z_az_sd_same_scaled
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (Zm << 16) | (msz << 10) | (Zn << 5) | Zd
	ins |= 0x420a000

	return
}

func encode_sve_int_bin_cons_misc_0_b(size, Zm, op, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_ftssel_z_zz_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 10) | (Zn << 5) | Zd
	ins |= 0x420b000

	return
}

func encode_sve_int_bin_cons_misc_0_c(size, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fexpa_z_z_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x1e) == 0x2:
		err = errUnallocated
	case (opc & 0x1c) == 0x4:
		err = errUnallocated
	case (opc & 0x18) == 0x8:
		err = errUnallocated
	case (opc & 0x10) == 0x10:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Zn << 5) | Zd
	ins |= 0x420b800

	return
}

func encode_sve_int_bin_cons_misc_0_d(opc, opc2, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || opc2 > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0 && opc2 == 0x0:
		// encoding_movprfx_z_z_
	case opc == 0x0 && opc2 == 0x1:
		err = errUnallocated
	case opc == 0x0 && (opc2&0x1e) == 0x2:
		err = errUnallocated
	case opc == 0x0 && (opc2&0x1c) == 0x4:
		err = errUnallocated
	case opc == 0x0 && (opc2&0x18) == 0x8:
		err = errUnallocated
	case opc == 0x0 && (opc2&0x10) == 0x10:
		err = errUnallocated
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (opc2 << 16) | (Zn << 5) | Zd
	ins |= 0x420bc00

	return
}

func encode_sve_int_countvlv0(size, imm4, D, U, pattern, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm4 > 0xf || D > 0x1 || U > 0x1 || pattern > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case size == 0x0:
		err = errUnallocated
	case size == 0x1 && D == 0x0 && U == 0x0:
		// encoding_sqinch_z_zs_
	case size == 0x1 && D == 0x0 && U == 0x1:
		// encoding_uqinch_z_zs_
	case size == 0x1 && D == 0x1 && U == 0x0:
		// encoding_sqdech_z_zs_
	case size == 0x1 && D == 0x1 && U == 0x1:
		// encoding_uqdech_z_zs_
	case size == 0x2 && D == 0x0 && U == 0x0:
		// encoding_sqincw_z_zs_
	case size == 0x2 && D == 0x0 && U == 0x1:
		// encoding_uqincw_z_zs_
	case size == 0x2 && D == 0x1 && U == 0x0:
		// encoding_sqdecw_z_zs_
	case size == 0x2 && D == 0x1 && U == 0x1:
		// encoding_uqdecw_z_zs_
	case size == 0x3 && D == 0x0 && U == 0x0:
		// encoding_sqincd_z_zs_
	case size == 0x3 && D == 0x0 && U == 0x1:
		// encoding_uqincd_z_zs_
	case size == 0x3 && D == 0x1 && U == 0x0:
		// encoding_sqdecd_z_zs_
	case size == 0x3 && D == 0x1 && U == 0x1:
		// encoding_uqdecd_z_zs_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (imm4 << 16) | (D << 11) | (U << 10) | (pattern << 5) | Zdn
	ins |= 0x420c000

	return
}

func encode_sve_int_count(size, imm4, op, pattern, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm4 > 0xf || op > 0x1 || pattern > 0x1f || Rd > 0x1f:
		err = errOverflow
	case op == 0x1:
		err = errUnallocated
	case size == 0x0 && op == 0x0:
		// encoding_cntb_r_s_
	case size == 0x1 && op == 0x0:
		// encoding_cnth_r_s_
	case size == 0x2 && op == 0x0:
		// encoding_cntw_r_s_
	case size == 0x3 && op == 0x0:
		// encoding_cntd_r_s_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (imm4 << 16) | (op << 10) | (pattern << 5) | Rd
	ins |= 0x420e000

	return
}

func encode_sve_int_countvlv1(size, imm4, D, pattern, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm4 > 0xf || D > 0x1 || pattern > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case size == 0x0:
		err = errUnallocated
	case size == 0x1 && D == 0x0:
		// encoding_inch_z_zs_
	case size == 0x1 && D == 0x1:
		// encoding_dech_z_zs_
	case size == 0x2 && D == 0x0:
		// encoding_incw_z_zs_
	case size == 0x2 && D == 0x1:
		// encoding_decw_z_zs_
	case size == 0x3 && D == 0x0:
		// encoding_incd_z_zs_
	case size == 0x3 && D == 0x1:
		// encoding_decd_z_zs_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (imm4 << 16) | (D << 10) | (pattern << 5) | Zdn
	ins |= 0x430c000

	return
}

func encode_sve_int_pred_pattern_a(size, imm4, D, pattern, Rdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm4 > 0xf || D > 0x1 || pattern > 0x1f || Rdn > 0x1f:
		err = errOverflow
	case size == 0x0 && D == 0x0:
		// encoding_incb_r_rs_
	case size == 0x0 && D == 0x1:
		// encoding_decb_r_rs_
	case size == 0x1 && D == 0x0:
		// encoding_inch_r_rs_
	case size == 0x1 && D == 0x1:
		// encoding_dech_r_rs_
	case size == 0x2 && D == 0x0:
		// encoding_incw_r_rs_
	case size == 0x2 && D == 0x1:
		// encoding_decw_r_rs_
	case size == 0x3 && D == 0x0:
		// encoding_incd_r_rs_
	case size == 0x3 && D == 0x1:
		// encoding_decd_r_rs_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (imm4 << 16) | (D << 10) | (pattern << 5) | Rdn
	ins |= 0x430e000

	return
}

func encode_sve_int_pred_pattern_b(size, sf, imm4, D, U, pattern, Rdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || sf > 0x1 || imm4 > 0xf || D > 0x1 || U > 0x1 || pattern > 0x1f || Rdn > 0x1f:
		err = errOverflow
	case size == 0x0 && sf == 0x0 && D == 0x0 && U == 0x0:
		// encoding_sqincb_r_rs_sx
	case size == 0x0 && sf == 0x0 && D == 0x0 && U == 0x1:
		// encoding_uqincb_r_rs_uw
	case size == 0x0 && sf == 0x0 && D == 0x1 && U == 0x0:
		// encoding_sqdecb_r_rs_sx
	case size == 0x0 && sf == 0x0 && D == 0x1 && U == 0x1:
		// encoding_uqdecb_r_rs_uw
	case size == 0x0 && sf == 0x1 && D == 0x0 && U == 0x0:
		// encoding_sqincb_r_rs_x
	case size == 0x0 && sf == 0x1 && D == 0x0 && U == 0x1:
		// encoding_uqincb_r_rs_x
	case size == 0x0 && sf == 0x1 && D == 0x1 && U == 0x0:
		// encoding_sqdecb_r_rs_x
	case size == 0x0 && sf == 0x1 && D == 0x1 && U == 0x1:
		// encoding_uqdecb_r_rs_x
	case size == 0x1 && sf == 0x0 && D == 0x0 && U == 0x0:
		// encoding_sqinch_r_rs_sx
	case size == 0x1 && sf == 0x0 && D == 0x0 && U == 0x1:
		// encoding_uqinch_r_rs_uw
	case size == 0x1 && sf == 0x0 && D == 0x1 && U == 0x0:
		// encoding_sqdech_r_rs_sx
	case size == 0x1 && sf == 0x0 && D == 0x1 && U == 0x1:
		// encoding_uqdech_r_rs_uw
	case size == 0x1 && sf == 0x1 && D == 0x0 && U == 0x0:
		// encoding_sqinch_r_rs_x
	case size == 0x1 && sf == 0x1 && D == 0x0 && U == 0x1:
		// encoding_uqinch_r_rs_x
	case size == 0x1 && sf == 0x1 && D == 0x1 && U == 0x0:
		// encoding_sqdech_r_rs_x
	case size == 0x1 && sf == 0x1 && D == 0x1 && U == 0x1:
		// encoding_uqdech_r_rs_x
	case size == 0x2 && sf == 0x0 && D == 0x0 && U == 0x0:
		// encoding_sqincw_r_rs_sx
	case size == 0x2 && sf == 0x0 && D == 0x0 && U == 0x1:
		// encoding_uqincw_r_rs_uw
	case size == 0x2 && sf == 0x0 && D == 0x1 && U == 0x0:
		// encoding_sqdecw_r_rs_sx
	case size == 0x2 && sf == 0x0 && D == 0x1 && U == 0x1:
		// encoding_uqdecw_r_rs_uw
	case size == 0x2 && sf == 0x1 && D == 0x0 && U == 0x0:
		// encoding_sqincw_r_rs_x
	case size == 0x2 && sf == 0x1 && D == 0x0 && U == 0x1:
		// encoding_uqincw_r_rs_x
	case size == 0x2 && sf == 0x1 && D == 0x1 && U == 0x0:
		// encoding_sqdecw_r_rs_x
	case size == 0x2 && sf == 0x1 && D == 0x1 && U == 0x1:
		// encoding_uqdecw_r_rs_x
	case size == 0x3 && sf == 0x0 && D == 0x0 && U == 0x0:
		// encoding_sqincd_r_rs_sx
	case size == 0x3 && sf == 0x0 && D == 0x0 && U == 0x1:
		// encoding_uqincd_r_rs_uw
	case size == 0x3 && sf == 0x0 && D == 0x1 && U == 0x0:
		// encoding_sqdecd_r_rs_sx
	case size == 0x3 && sf == 0x0 && D == 0x1 && U == 0x1:
		// encoding_uqdecd_r_rs_uw
	case size == 0x3 && sf == 0x1 && D == 0x0 && U == 0x0:
		// encoding_sqincd_r_rs_x
	case size == 0x3 && sf == 0x1 && D == 0x0 && U == 0x1:
		// encoding_uqincd_r_rs_x
	case size == 0x3 && sf == 0x1 && D == 0x1 && U == 0x0:
		// encoding_sqdecd_r_rs_x
	case size == 0x3 && sf == 0x1 && D == 0x1 && U == 0x1:
		// encoding_uqdecd_r_rs_x
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (sf << 20) | (imm4 << 16) | (D << 11) | (U << 10) | (pattern << 5) | Rdn
	ins |= 0x420f000

	return
}

func encode_sve_int_dup_mask_imm(imm13, Zd uint32) (ins uint32, err error) {
	switch {
	case imm13 > 0x1fff || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_dupm_z_i_
	}

	ins |= (imm13 << 5) | Zd
	ins |= 0x5c00000

	return
}

func encode_sve_int_log_imm(opc, imm13, Zdn uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || imm13 > 0x1fff || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_orr_z_zi_
	case opc == 0x1:
		// encoding_eor_z_zi_
	case opc == 0x2:
		// encoding_and_z_zi_
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (imm13 << 5) | Zdn
	ins |= 0x5000000

	return
}

func encode_sve_int_dup_imm_pred(size, Pg, M, sh, imm8, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0xf || M > 0x1 || sh > 0x1 || imm8 > 0xff || Zd > 0x1f:
		err = errOverflow
	case M == 0x0:
		// encoding_cpy_z_o_i_
	case M == 0x1:
		// encoding_cpy_z_p_i_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Pg << 16) | (M << 14) | (sh << 13) | (imm8 << 5) | Zd
	ins |= 0x5100000

	return
}

func encode_sve_int_dup_fpimm_pred(size, Pg, imm8, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0xf || imm8 > 0xff || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_fcpy_z_p_i_
	}

	ins |= (size << 22) | (Pg << 16) | (imm8 << 5) | Zd
	ins |= 0x510c000

	return
}

func encode_sve_int_perm_dup_i(imm2, tsz, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case imm2 > 0x3 || tsz > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_dup_z_zi_
	}

	ins |= (imm2 << 22) | (tsz << 16) | (Zn << 5) | Zd
	ins |= 0x5202000

	return
}

func encode_sve_int_perm_tbl_3src(size, Zm, op, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_tbl_z_zz_2
	case op == 0x1:
		// encoding_tbx_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 10) | (Zn << 5) | Zd
	ins |= 0x5202800

	return
}

func encode_sve_int_perm_tbl(size, Zm, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_tbl_z_zz_1
	}

	ins |= (size << 22) | (Zm << 16) | (Zn << 5) | Zd
	ins |= 0x5203000

	return
}

func encode_sve_int_perm_dup_r(size, Rn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_dup_z_r_
	}

	ins |= (size << 22) | (Rn << 5) | Zd
	ins |= 0x5203800

	return
}

func encode_sve_int_perm_insrs(size, Rm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_insr_z_r_
	}

	ins |= (size << 22) | (Rm << 5) | Zdn
	ins |= 0x5243800

	return
}

func encode_sve_int_perm_unpk(size, U, H, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || U > 0x1 || H > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case U == 0x0 && H == 0x0:
		// encoding_sunpklo_z_z_
	case U == 0x0 && H == 0x1:
		// encoding_sunpkhi_z_z_
	case U == 0x1 && H == 0x0:
		// encoding_uunpklo_z_z_
	case U == 0x1 && H == 0x1:
		// encoding_uunpkhi_z_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (U << 17) | (H << 16) | (Zn << 5) | Zd
	ins |= 0x5303800

	return
}

func encode_sve_int_perm_insrv(size, Vm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Vm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_insr_z_v_
	}

	ins |= (size << 22) | (Vm << 5) | Zdn
	ins |= 0x5343800

	return
}

func encode_sve_int_perm_reverse_z(size, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_rev_z_z_
	}

	ins |= (size << 22) | (Zn << 5) | Zd
	ins |= 0x5383800

	return
}

func encode_sve_int_perm_punpk(H, Pn, Pd uint32) (ins uint32, err error) {
	switch {
	case H > 0x1 || Pn > 0xf || Pd > 0xf:
		err = errOverflow
	case H == 0x0:
		// encoding_punpklo_p_p_
	case H == 0x1:
		// encoding_punpkhi_p_p_
	default:
		err = errUnmatched
	}

	ins |= (H << 16) | (Pn << 5) | Pd
	ins |= 0x5304000

	return
}

func encode_sve_int_perm_bin_perm_pp(size, Pm, opc, H, Pn, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pm > 0xf || opc > 0x3 || H > 0x1 || Pn > 0xf || Pd > 0xf:
		err = errOverflow
	case opc == 0x0 && H == 0x0:
		// encoding_zip1_p_pp_
	case opc == 0x0 && H == 0x1:
		// encoding_zip2_p_pp_
	case opc == 0x1 && H == 0x0:
		// encoding_uzp1_p_pp_
	case opc == 0x1 && H == 0x1:
		// encoding_uzp2_p_pp_
	case opc == 0x2 && H == 0x0:
		// encoding_trn1_p_pp_
	case opc == 0x2 && H == 0x1:
		// encoding_trn2_p_pp_
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Pm << 16) | (opc << 11) | (H << 10) | (Pn << 5) | Pd
	ins |= 0x5204000

	return
}

func encode_sve_int_perm_reverse_p(size, Pn, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pn > 0xf || Pd > 0xf:
		err = errOverflow
	default:
		// encoding_rev_p_p_
	}

	ins |= (size << 22) | (Pn << 5) | Pd
	ins |= 0x5344000

	return
}

func encode_sve_int_perm_bin_perm_zz(size, Zm, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_zip1_z_zz_
	case opc == 0x1:
		// encoding_zip2_z_zz_
	case opc == 0x2:
		// encoding_uzp1_z_zz_
	case opc == 0x3:
		// encoding_uzp2_z_zz_
	case opc == 0x4:
		// encoding_trn1_z_zz_
	case opc == 0x5:
		// encoding_trn2_z_zz_
	case (opc & 0x6) == 0x6:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x5206000

	return
}

func encode_sve_int_perm_cpy_v(size, Pg, Vn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0x7 || Vn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_cpy_z_p_v_
	}

	ins |= (size << 22) | (Pg << 10) | (Vn << 5) | Zd
	ins |= 0x5208000

	return
}

func encode_sve_int_perm_compact(size, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_compact_z_p_z_
	}

	ins |= (size << 22) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x5218000

	return
}

func encode_sve_int_perm_last_r(size, B, Pg, Zn, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || B > 0x1 || Pg > 0x7 || Zn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case B == 0x0:
		// encoding_lasta_r_p_z_
	case B == 0x1:
		// encoding_lastb_r_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (B << 16) | (Pg << 10) | (Zn << 5) | Rd
	ins |= 0x520a000

	return
}

func encode_sve_int_perm_last_v(size, B, Pg, Zn, Vd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || B > 0x1 || Pg > 0x7 || Zn > 0x1f || Vd > 0x1f:
		err = errOverflow
	case B == 0x0:
		// encoding_lasta_v_p_z_
	case B == 0x1:
		// encoding_lastb_v_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (B << 16) | (Pg << 10) | (Zn << 5) | Vd
	ins |= 0x5228000

	return
}

func encode_sve_int_perm_rev(size, opc, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_revb_z_z_
	case opc == 0x1:
		// encoding_revh_z_z_
	case opc == 0x2:
		// encoding_revw_z_z_
	case opc == 0x3:
		// encoding_rbit_z_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x5248000

	return
}

func encode_sve_int_perm_cpy_r(size, Pg, Rn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0x7 || Rn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_cpy_z_p_r_
	}

	ins |= (size << 22) | (Pg << 10) | (Rn << 5) | Zd
	ins |= 0x528a000

	return
}

func encode_sve_int_perm_clast_zz(size, B, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || B > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case B == 0x0:
		// encoding_clasta_z_p_zz_
	case B == 0x1:
		// encoding_clastb_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (B << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x5288000

	return
}

func encode_sve_int_perm_clast_vz(size, B, Pg, Zm, Vdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || B > 0x1 || Pg > 0x7 || Zm > 0x1f || Vdn > 0x1f:
		err = errOverflow
	case B == 0x0:
		// encoding_clasta_v_p_z_
	case B == 0x1:
		// encoding_clastb_v_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (B << 16) | (Pg << 10) | (Zm << 5) | Vdn
	ins |= 0x52a8000

	return
}

func encode_sve_int_perm_splice(size, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_splice_z_p_zz_des
	}

	ins |= (size << 22) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x52c8000

	return
}

func encode_sve_intx_perm_splice(size, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_splice_z_p_zz_con
	}

	ins |= (size << 22) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x52d8000

	return
}

func encode_sve_int_perm_revd(size, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case size == 0x0:
		// encoding_revd_z_p_z_
	case size == 0x1:
		err = errUnallocated
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x52e8000

	return
}

func encode_sve_int_perm_clast_rz(size, B, Pg, Zm, Rdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || B > 0x1 || Pg > 0x7 || Zm > 0x1f || Rdn > 0x1f:
		err = errOverflow
	case B == 0x0:
		// encoding_clasta_r_p_z_
	case B == 0x1:
		// encoding_clastb_r_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (B << 16) | (Pg << 10) | (Zm << 5) | Rdn
	ins |= 0x530a000

	return
}

func encode_sve_int_sel_vvv(size, Zm, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || Pg > 0xf || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_sel_z_p_zz_
	}

	ins |= (size << 22) | (Zm << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x520c000

	return
}

func encode_sve_int_perm_extract_i(imm8h, imm8l, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case imm8h > 0x1f || imm8l > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_ext_z_zi_des
	}

	ins |= (imm8h << 16) | (imm8l << 10) | (Zm << 5) | Zdn
	ins |= 0x5200000

	return
}

func encode_sve_intx_perm_extract_i(imm8h, imm8l, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case imm8h > 0x1f || imm8l > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_ext_z_zi_con
	}

	ins |= (imm8h << 16) | (imm8l << 10) | (Zn << 5) | Zd
	ins |= 0x5600000

	return
}

func encode_sve_int_perm_bin_long_perm_zz(Zm, opc, H, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case Zm > 0x1f || opc > 0x3 || H > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0 && H == 0x0:
		// encoding_zip1_z_zz_q
	case opc == 0x0 && H == 0x1:
		// encoding_zip2_z_zz_q
	case opc == 0x1 && H == 0x0:
		// encoding_uzp1_z_zz_q
	case opc == 0x1 && H == 0x1:
		// encoding_uzp2_z_zz_q
	case opc == 0x2:
		err = errUnallocated
	case opc == 0x3 && H == 0x0:
		// encoding_trn1_z_zz_q
	case opc == 0x3 && H == 0x1:
		// encoding_trn2_z_zz_q
	default:
		err = errUnmatched
	}

	ins |= (Zm << 16) | (opc << 11) | (H << 10) | (Zn << 5) | Zd
	ins |= 0x5a00000

	return
}

func encode_sve_int_cmp_0(size, Zm, op, o2, Pg, Zn, ne, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || o2 > 0x1 || Pg > 0x7 || Zn > 0x1f || ne > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && o2 == 0x0 && ne == 0x0:
		// encoding_cmphs_p_p_zz_
	case op == 0x0 && o2 == 0x0 && ne == 0x1:
		// encoding_cmphi_p_p_zz_
	case op == 0x0 && o2 == 0x1 && ne == 0x0:
		// encoding_cmpeq_p_p_zw_
	case op == 0x0 && o2 == 0x1 && ne == 0x1:
		// encoding_cmpne_p_p_zw_
	case op == 0x1 && o2 == 0x0 && ne == 0x0:
		// encoding_cmpge_p_p_zz_
	case op == 0x1 && o2 == 0x0 && ne == 0x1:
		// encoding_cmpgt_p_p_zz_
	case op == 0x1 && o2 == 0x1 && ne == 0x0:
		// encoding_cmpeq_p_p_zz_
	case op == 0x1 && o2 == 0x1 && ne == 0x1:
		// encoding_cmpne_p_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 15) | (o2 << 13) | (Pg << 10) | (Zn << 5) | (ne << 4) | Pd
	ins |= 0x24000000

	return
}

func encode_sve_int_cmp_1(size, Zm, U, lt, Pg, Zn, ne, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || U > 0x1 || lt > 0x1 || Pg > 0x7 || Zn > 0x1f || ne > 0x1 || Pd > 0xf:
		err = errOverflow
	case U == 0x0 && lt == 0x0 && ne == 0x0:
		// encoding_cmpge_p_p_zw_
	case U == 0x0 && lt == 0x0 && ne == 0x1:
		// encoding_cmpgt_p_p_zw_
	case U == 0x0 && lt == 0x1 && ne == 0x0:
		// encoding_cmplt_p_p_zw_
	case U == 0x0 && lt == 0x1 && ne == 0x1:
		// encoding_cmple_p_p_zw_
	case U == 0x1 && lt == 0x0 && ne == 0x0:
		// encoding_cmphs_p_p_zw_
	case U == 0x1 && lt == 0x0 && ne == 0x1:
		// encoding_cmphi_p_p_zw_
	case U == 0x1 && lt == 0x1 && ne == 0x0:
		// encoding_cmplo_p_p_zw_
	case U == 0x1 && lt == 0x1 && ne == 0x1:
		// encoding_cmpls_p_p_zw_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (U << 15) | (lt << 13) | (Pg << 10) | (Zn << 5) | (ne << 4) | Pd
	ins |= 0x24004000

	return
}

func encode_sve_int_ucmp_vi(size, imm7, lt, Pg, Zn, ne, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm7 > 0x7f || lt > 0x1 || Pg > 0x7 || Zn > 0x1f || ne > 0x1 || Pd > 0xf:
		err = errOverflow
	case lt == 0x0 && ne == 0x0:
		// encoding_cmphs_p_p_zi_
	case lt == 0x0 && ne == 0x1:
		// encoding_cmphi_p_p_zi_
	case lt == 0x1 && ne == 0x0:
		// encoding_cmplo_p_p_zi_
	case lt == 0x1 && ne == 0x1:
		// encoding_cmpls_p_p_zi_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (imm7 << 14) | (lt << 13) | (Pg << 10) | (Zn << 5) | (ne << 4) | Pd
	ins |= 0x24200000

	return
}

func encode_sve_int_scmp_vi(size, imm5, op, o2, Pg, Zn, ne, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm5 > 0x1f || op > 0x1 || o2 > 0x1 || Pg > 0x7 || Zn > 0x1f || ne > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && o2 == 0x0 && ne == 0x0:
		// encoding_cmpge_p_p_zi_
	case op == 0x0 && o2 == 0x0 && ne == 0x1:
		// encoding_cmpgt_p_p_zi_
	case op == 0x0 && o2 == 0x1 && ne == 0x0:
		// encoding_cmplt_p_p_zi_
	case op == 0x0 && o2 == 0x1 && ne == 0x1:
		// encoding_cmple_p_p_zi_
	case op == 0x1 && o2 == 0x0 && ne == 0x0:
		// encoding_cmpeq_p_p_zi_
	case op == 0x1 && o2 == 0x0 && ne == 0x1:
		// encoding_cmpne_p_p_zi_
	case op == 0x1 && o2 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (imm5 << 16) | (op << 15) | (o2 << 13) | (Pg << 10) | (Zn << 5) | (ne << 4) | Pd
	ins |= 0x25000000

	return
}

func encode_sve_int_pred_log(op, S, Pm, Pg, o2, Pn, o3, Pd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pm > 0xf || Pg > 0xf || o2 > 0x1 || Pn > 0xf || o3 > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0 && o2 == 0x0 && o3 == 0x0:
		// encoding_and_p_p_pp_z
	case op == 0x0 && S == 0x0 && o2 == 0x0 && o3 == 0x1:
		// encoding_bic_p_p_pp_z
	case op == 0x0 && S == 0x0 && o2 == 0x1 && o3 == 0x0:
		// encoding_eor_p_p_pp_z
	case op == 0x0 && S == 0x0 && o2 == 0x1 && o3 == 0x1:
		// encoding_sel_p_p_pp_
	case op == 0x0 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_ands_p_p_pp_z
	case op == 0x0 && S == 0x1 && o2 == 0x0 && o3 == 0x1:
		// encoding_bics_p_p_pp_z
	case op == 0x0 && S == 0x1 && o2 == 0x1 && o3 == 0x0:
		// encoding_eors_p_p_pp_z
	case op == 0x0 && S == 0x1 && o2 == 0x1 && o3 == 0x1:
		err = errUnallocated
	case op == 0x1 && S == 0x0 && o2 == 0x0 && o3 == 0x0:
		// encoding_orr_p_p_pp_z
	case op == 0x1 && S == 0x0 && o2 == 0x0 && o3 == 0x1:
		// encoding_orn_p_p_pp_z
	case op == 0x1 && S == 0x0 && o2 == 0x1 && o3 == 0x0:
		// encoding_nor_p_p_pp_z
	case op == 0x1 && S == 0x0 && o2 == 0x1 && o3 == 0x1:
		// encoding_nand_p_p_pp_z
	case op == 0x1 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_orrs_p_p_pp_z
	case op == 0x1 && S == 0x1 && o2 == 0x0 && o3 == 0x1:
		// encoding_orns_p_p_pp_z
	case op == 0x1 && S == 0x1 && o2 == 0x1 && o3 == 0x0:
		// encoding_nors_p_p_pp_z
	case op == 0x1 && S == 0x1 && o2 == 0x1 && o3 == 0x1:
		// encoding_nands_p_p_pp_z
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | (Pm << 16) | (Pg << 10) | (o2 << 9) | (Pn << 5) | (o3 << 4) | Pd
	ins |= 0x25004000

	return
}

func encode_sve_int_brkp(op, S, Pm, Pg, Pn, B, Pd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pm > 0xf || Pg > 0xf || Pn > 0xf || B > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0 && B == 0x0:
		// encoding_brkpa_p_p_pp_
	case op == 0x0 && S == 0x0 && B == 0x1:
		// encoding_brkpb_p_p_pp_
	case op == 0x0 && S == 0x1 && B == 0x0:
		// encoding_brkpas_p_p_pp_
	case op == 0x0 && S == 0x1 && B == 0x1:
		// encoding_brkpbs_p_p_pp_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | (Pm << 16) | (Pg << 10) | (Pn << 5) | (B << 4) | Pd
	ins |= 0x2500c000

	return
}

func encode_sve_int_brkn(S, Pg, Pn, Pdm uint32) (ins uint32, err error) {
	switch {
	case S > 0x1 || Pg > 0xf || Pn > 0xf || Pdm > 0xf:
		err = errOverflow
	case S == 0x0:
		// encoding_brkn_p_p_pp_
	case S == 0x1:
		// encoding_brkns_p_p_pp_
	default:
		err = errUnmatched
	}

	ins |= (S << 22) | (Pg << 10) | (Pn << 5) | Pdm
	ins |= 0x25184000

	return
}

func encode_sve_int_break(B, S, Pg, Pn, M, Pd uint32) (ins uint32, err error) {
	switch {
	case B > 0x1 || S > 0x1 || Pg > 0xf || Pn > 0xf || M > 0x1 || Pd > 0xf:
		err = errOverflow
	case S == 0x1 && M == 0x1:
		err = errUnallocated
	case B == 0x0 && S == 0x0:
		// encoding_brka_p_p_p_
	case B == 0x0 && S == 0x1 && M == 0x0:
		// encoding_brkas_p_p_p_z
	case B == 0x1 && S == 0x0:
		// encoding_brkb_p_p_p_
	case B == 0x1 && S == 0x1 && M == 0x0:
		// encoding_brkbs_p_p_p_z
	default:
		err = errUnmatched
	}

	ins |= (B << 23) | (S << 22) | (Pg << 10) | (Pn << 5) | (M << 4) | Pd
	ins |= 0x25104000

	return
}

func encode_sve_int_ptest(op, S, Pg, Pn, opc2 uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pg > 0xf || Pn > 0xf || opc2 > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0:
		err = errUnallocated
	case op == 0x0 && S == 0x1 && opc2 == 0x0:
		// encoding_ptest_p_p_
	case op == 0x0 && S == 0x1 && opc2 == 0x1:
		err = errUnallocated
	case op == 0x0 && S == 0x1 && (opc2&0xe) == 0x2:
		err = errUnallocated
	case op == 0x0 && S == 0x1 && (opc2&0xc) == 0x4:
		err = errUnallocated
	case op == 0x0 && S == 0x1 && (opc2&0x8) == 0x8:
		err = errUnallocated
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | (Pg << 10) | (Pn << 5) | opc2
	ins |= 0x2510c000

	return
}

func encode_sve_int_pfirst(op, S, Pg, Pdn uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pg > 0xf || Pdn > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0:
		err = errUnallocated
	case op == 0x0 && S == 0x1:
		// encoding_pfirst_p_p_p_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | (Pg << 5) | Pdn
	ins |= 0x2518c000

	return
}

func encode_sve_int_pfalse(op, S, Pd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0:
		// encoding_pfalse_p_
	case op == 0x0 && S == 0x1:
		err = errUnallocated
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | Pd
	ins |= 0x2518e400

	return
}

func encode_sve_int_rdffr(op, S, Pg, Pd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pg > 0xf || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0:
		// encoding_rdffr_p_p_f_
	case op == 0x0 && S == 0x1:
		// encoding_rdffrs_p_p_f_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | (Pg << 5) | Pd
	ins |= 0x2518f000

	return
}

func encode_sve_int_pnext(size, Pg, Pdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Pg > 0xf || Pdn > 0xf:
		err = errOverflow
	default:
		// encoding_pnext_p_p_p_
	}

	ins |= (size << 22) | (Pg << 5) | Pdn
	ins |= 0x2519c400

	return
}

func encode_sve_int_rdffr_2(op, S, Pd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || S > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && S == 0x0:
		// encoding_rdffr_p_f_
	case op == 0x0 && S == 0x1:
		err = errUnallocated
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (S << 22) | Pd
	ins |= 0x2519f000

	return
}

func encode_sve_int_ptrue(size, S, pattern, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || S > 0x1 || pattern > 0x1f || Pd > 0xf:
		err = errOverflow
	case S == 0x0:
		// encoding_ptrue_p_s_
	case S == 0x1:
		// encoding_ptrues_p_s_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (S << 16) | (pattern << 5) | Pd
	ins |= 0x2518e000

	return
}

func encode_sve_int_while_rr(size, Rm, sf, U, lt, Rn, eq, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rm > 0x1f || sf > 0x1 || U > 0x1 || lt > 0x1 || Rn > 0x1f || eq > 0x1 || Pd > 0xf:
		err = errOverflow
	case U == 0x0 && lt == 0x0 && eq == 0x0:
		// encoding_whilege_p_p_rr_
	case U == 0x0 && lt == 0x0 && eq == 0x1:
		// encoding_whilegt_p_p_rr_
	case U == 0x0 && lt == 0x1 && eq == 0x0:
		// encoding_whilelt_p_p_rr_
	case U == 0x0 && lt == 0x1 && eq == 0x1:
		// encoding_whilele_p_p_rr_
	case U == 0x1 && lt == 0x0 && eq == 0x0:
		// encoding_whilehs_p_p_rr_
	case U == 0x1 && lt == 0x0 && eq == 0x1:
		// encoding_whilehi_p_p_rr_
	case U == 0x1 && lt == 0x1 && eq == 0x0:
		// encoding_whilelo_p_p_rr_
	case U == 0x1 && lt == 0x1 && eq == 0x1:
		// encoding_whilels_p_p_rr_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Rm << 16) | (sf << 12) | (U << 11) | (lt << 10) | (Rn << 5) | (eq << 4) | Pd
	ins |= 0x25200000

	return
}

func encode_sve_int_cterm(op, sz, Rm, Rn, ne uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || sz > 0x1 || Rm > 0x1f || Rn > 0x1f || ne > 0x1:
		err = errOverflow
	case op == 0x0:
		err = errUnallocated
	case op == 0x1 && ne == 0x0:
		// encoding_ctermeq_rr_
	case op == 0x1 && ne == 0x1:
		// encoding_ctermne_rr_
	default:
		err = errUnmatched
	}

	ins |= (op << 23) | (sz << 22) | (Rm << 16) | (Rn << 5) | (ne << 4)
	ins |= 0x25202000

	return
}

func encode_sve_int_whilenc(size, Rm, Rn, rw, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rm > 0x1f || Rn > 0x1f || rw > 0x1 || Pd > 0xf:
		err = errOverflow
	case rw == 0x0:
		// encoding_whilewr_p_rr_
	case rw == 0x1:
		// encoding_whilerw_p_rr_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Rm << 16) | (Rn << 5) | (rw << 4) | Pd
	ins |= 0x25203000

	return
}

func encode_sve_int_pred_dup(i1, tszh, tszl, Rv, Pn, S, Pm, Pd uint32) (ins uint32, err error) {
	switch {
	case i1 > 0x1 || tszh > 0x1 || tszl > 0x7 || Rv > 0x3 || Pn > 0xf || S > 0x1 || Pm > 0xf || Pd > 0xf:
		err = errOverflow
	case S == 0x0:
		// encoding_psel_p_ppi_
	case S == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (i1 << 23) | (tszh << 22) | (tszl << 18) | (Rv << 16) | (Pn << 10) | (S << 9) | (Pm << 5) | Pd
	ins |= 0x25204000

	return
}

func encode_sve_int_arith_imm0(size, opc, sh, imm8, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || sh > 0x1 || imm8 > 0xff || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_add_z_zi_
	case opc == 0x1:
		// encoding_sub_z_zi_
	case opc == 0x2:
		err = errUnallocated
	case opc == 0x3:
		// encoding_subr_z_zi_
	case opc == 0x4:
		// encoding_sqadd_z_zi_
	case opc == 0x5:
		// encoding_uqadd_z_zi_
	case opc == 0x6:
		// encoding_sqsub_z_zi_
	case opc == 0x7:
		// encoding_uqsub_z_zi_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (sh << 13) | (imm8 << 5) | Zdn
	ins |= 0x2520c000

	return
}

func encode_sve_int_arith_imm1(size, opc, o2, imm8, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || o2 > 0x1 || imm8 > 0xff || Zdn > 0x1f:
		err = errOverflow
	case (opc&0x4) == 0x0 && o2 == 0x1:
		err = errUnallocated
	case opc == 0x0 && o2 == 0x0:
		// encoding_smax_z_zi_
	case opc == 0x1 && o2 == 0x0:
		// encoding_umax_z_zi_
	case opc == 0x2 && o2 == 0x0:
		// encoding_smin_z_zi_
	case opc == 0x3 && o2 == 0x0:
		// encoding_umin_z_zi_
	case (opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (o2 << 13) | (imm8 << 5) | Zdn
	ins |= 0x2528c000

	return
}

func encode_sve_int_arith_imm2(size, opc, o2, imm8, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || o2 > 0x1 || imm8 > 0xff || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0 && o2 == 0x0:
		// encoding_mul_z_zi_
	case opc == 0x0 && o2 == 0x1:
		err = errUnallocated
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x6) == 0x2:
		err = errUnallocated
	case (opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (o2 << 13) | (imm8 << 5) | Zdn
	ins |= 0x2530c000

	return
}

func encode_sve_int_dup_imm(size, opc, sh, imm8, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || sh > 0x1 || imm8 > 0xff || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_dup_z_i_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 17) | (sh << 13) | (imm8 << 5) | Zd
	ins |= 0x2538c000

	return
}

func encode_sve_int_dup_fpimm(size, opc, o2, imm8, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || o2 > 0x1 || imm8 > 0xff || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0 && o2 == 0x0:
		// encoding_fdup_z_i_
	case opc == 0x0 && o2 == 0x1:
		err = errUnallocated
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 17) | (o2 << 13) | (imm8 << 5) | Zd
	ins |= 0x2539c000

	return
}

func encode_sve_int_pcount_pred(size, opc, Pg, Pn, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0xf || Pn > 0xf || Rd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_cntp_r_p_p_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x6) == 0x2:
		err = errUnallocated
	case (opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Pn << 5) | Rd
	ins |= 0x25208000

	return
}

func encode_sve_int_count_v_sat(size, D, U, opc, Pm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || D > 0x1 || U > 0x1 || opc > 0x3 || Pm > 0xf || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	case D == 0x0 && U == 0x0 && opc == 0x0:
		// encoding_sqincp_z_p_z_
	case D == 0x0 && U == 0x1 && opc == 0x0:
		// encoding_uqincp_z_p_z_
	case D == 0x1 && U == 0x0 && opc == 0x0:
		// encoding_sqdecp_z_p_z_
	case D == 0x1 && U == 0x1 && opc == 0x0:
		// encoding_uqdecp_z_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (D << 17) | (U << 16) | (opc << 9) | (Pm << 5) | Zdn
	ins |= 0x25288000

	return
}

func encode_sve_int_count_r_sat(size, D, U, sf, op, Pm, Rdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || D > 0x1 || U > 0x1 || sf > 0x1 || op > 0x1 || Pm > 0xf || Rdn > 0x1f:
		err = errOverflow
	case op == 0x1:
		err = errUnallocated
	case D == 0x0 && U == 0x0 && sf == 0x0 && op == 0x0:
		// encoding_sqincp_r_p_r_sx
	case D == 0x0 && U == 0x0 && sf == 0x1 && op == 0x0:
		// encoding_sqincp_r_p_r_x
	case D == 0x0 && U == 0x1 && sf == 0x0 && op == 0x0:
		// encoding_uqincp_r_p_r_uw
	case D == 0x0 && U == 0x1 && sf == 0x1 && op == 0x0:
		// encoding_uqincp_r_p_r_x
	case D == 0x1 && U == 0x0 && sf == 0x0 && op == 0x0:
		// encoding_sqdecp_r_p_r_sx
	case D == 0x1 && U == 0x0 && sf == 0x1 && op == 0x0:
		// encoding_sqdecp_r_p_r_x
	case D == 0x1 && U == 0x1 && sf == 0x0 && op == 0x0:
		// encoding_uqdecp_r_p_r_uw
	case D == 0x1 && U == 0x1 && sf == 0x1 && op == 0x0:
		// encoding_uqdecp_r_p_r_x
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (D << 17) | (U << 16) | (sf << 10) | (op << 9) | (Pm << 5) | Rdn
	ins |= 0x25288800

	return
}

func encode_sve_int_count_v(size, op, D, opc2, Pm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || D > 0x1 || opc2 > 0x3 || Pm > 0xf || Zdn > 0x1f:
		err = errOverflow
	case op == 0x0 && opc2 == 0x1:
		err = errUnallocated
	case op == 0x0 && (opc2&0x2) == 0x2:
		err = errUnallocated
	case op == 0x0 && D == 0x0 && opc2 == 0x0:
		// encoding_incp_z_p_z_
	case op == 0x0 && D == 0x1 && opc2 == 0x0:
		// encoding_decp_z_p_z_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 17) | (D << 16) | (opc2 << 9) | (Pm << 5) | Zdn
	ins |= 0x252c8000

	return
}

func encode_sve_int_count_r(size, op, D, opc2, Pm, Rdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || D > 0x1 || opc2 > 0x3 || Pm > 0xf || Rdn > 0x1f:
		err = errOverflow
	case op == 0x0 && opc2 == 0x1:
		err = errUnallocated
	case op == 0x0 && (opc2&0x2) == 0x2:
		err = errUnallocated
	case op == 0x0 && D == 0x0 && opc2 == 0x0:
		// encoding_incp_r_p_r_
	case op == 0x0 && D == 0x1 && opc2 == 0x0:
		// encoding_decp_r_p_r_
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 17) | (D << 16) | (opc2 << 9) | (Pm << 5) | Rdn
	ins |= 0x252c8800

	return
}

func encode_sve_int_wrffr(opc, Pn uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || Pn > 0xf:
		err = errOverflow
	case opc == 0x0:
		// encoding_wrffr_f_p_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (Pn << 5)
	ins |= 0x25289000

	return
}

func encode_sve_int_setffr(opc uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3:
		err = errOverflow
	case opc == 0x0:
		// encoding_setffr_f_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 22)
	ins |= 0x252c9000

	return
}

func encode_sve_intx_dot(size, Zm, U, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || U > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case U == 0x0:
		// encoding_sdot_z_zzz_
	case U == 0x1:
		// encoding_udot_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (U << 10) | (Zn << 5) | Zda
	ins |= 0x44000000

	return
}

func encode_sve_intx_qdmlalbt(size, Zm, S, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case S == 0x0:
		// encoding_sqdmlalbt_z_zzz_
	case S == 0x1:
		// encoding_sqdmlslbt_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 10) | (Zn << 5) | Zda
	ins |= 0x44000800

	return
}

func encode_sve_intx_cdot(size, Zm, rot, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || rot > 0x3 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	default:
		// encoding_cdot_z_zzz_
	}

	ins |= (size << 22) | (Zm << 16) | (rot << 10) | (Zn << 5) | Zda
	ins |= 0x44001000

	return
}

func encode_sve_intx_cmla(size, Zm, op, rot, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || rot > 0x3 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_cmla_z_zzz_
	case op == 0x1:
		// encoding_sqrdcmlah_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 12) | (rot << 10) | (Zn << 5) | Zda
	ins |= 0x44002000

	return
}

func encode_sve_intx_mlal_long(size, Zm, S, U, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || U > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case S == 0x0 && U == 0x0 && T == 0x0:
		// encoding_smlalb_z_zzz_
	case S == 0x0 && U == 0x0 && T == 0x1:
		// encoding_smlalt_z_zzz_
	case S == 0x0 && U == 0x1 && T == 0x0:
		// encoding_umlalb_z_zzz_
	case S == 0x0 && U == 0x1 && T == 0x1:
		// encoding_umlalt_z_zzz_
	case S == 0x1 && U == 0x0 && T == 0x0:
		// encoding_smlslb_z_zzz_
	case S == 0x1 && U == 0x0 && T == 0x1:
		// encoding_smlslt_z_zzz_
	case S == 0x1 && U == 0x1 && T == 0x0:
		// encoding_umlslb_z_zzz_
	case S == 0x1 && U == 0x1 && T == 0x1:
		// encoding_umlslt_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 12) | (U << 11) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x44004000

	return
}

func encode_sve_intx_qdmlal_long(size, Zm, S, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case S == 0x0 && T == 0x0:
		// encoding_sqdmlalb_z_zzz_
	case S == 0x0 && T == 0x1:
		// encoding_sqdmlalt_z_zzz_
	case S == 0x1 && T == 0x0:
		// encoding_sqdmlslb_z_zzz_
	case S == 0x1 && T == 0x1:
		// encoding_sqdmlslt_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 11) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x44006000

	return
}

func encode_sve_intx_qrdmlah(size, Zm, S, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case S == 0x0:
		// encoding_sqrdmlah_z_zzz_
	case S == 0x1:
		// encoding_sqrdmlsh_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 10) | (Zn << 5) | Zda
	ins |= 0x44007000

	return
}

func encode_sve_intx_mixed_dot(size, Zm, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2:
		// encoding_usdot_z_zzz_s
	case size == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (Zn << 5) | Zda
	ins |= 0x44007800

	return
}

func encode_sve_intx_accumulate_long_pairs(size, U, Pg, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case U == 0x0:
		// encoding_sadalp_z_p_z_
	case U == 0x1:
		// encoding_uadalp_z_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (U << 16) | (Pg << 10) | (Zn << 5) | Zda
	ins |= 0x4404a000

	return
}

func encode_sve_intx_pred_arith_unary(size, Q, opc, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Q > 0x1 || opc > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	case Q == 0x0 && opc == 0x0:
		// encoding_urecpe_z_p_z_
	case Q == 0x0 && opc == 0x1:
		// encoding_ursqrte_z_p_z_
	case Q == 0x1 && opc == 0x0:
		// encoding_sqabs_z_p_z_
	case Q == 0x1 && opc == 0x1:
		// encoding_sqneg_z_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Q << 19) | (opc << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x4400a000

	return
}

func encode_sve_intx_bin_pred_shift_sat_round(size, Q, R, N, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Q > 0x1 || R > 0x1 || N > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case Q == 0x0 && N == 0x0:
		err = errUnallocated
	case Q == 0x0 && R == 0x0 && N == 0x1 && U == 0x0:
		// encoding_srshl_z_p_zz_
	case Q == 0x0 && R == 0x0 && N == 0x1 && U == 0x1:
		// encoding_urshl_z_p_zz_
	case Q == 0x0 && R == 0x1 && N == 0x1 && U == 0x0:
		// encoding_srshlr_z_p_zz_
	case Q == 0x0 && R == 0x1 && N == 0x1 && U == 0x1:
		// encoding_urshlr_z_p_zz_
	case Q == 0x1 && R == 0x0 && N == 0x0 && U == 0x0:
		// encoding_sqshl_z_p_zz_
	case Q == 0x1 && R == 0x0 && N == 0x0 && U == 0x1:
		// encoding_uqshl_z_p_zz_
	case Q == 0x1 && R == 0x0 && N == 0x1 && U == 0x0:
		// encoding_sqrshl_z_p_zz_
	case Q == 0x1 && R == 0x0 && N == 0x1 && U == 0x1:
		// encoding_uqrshl_z_p_zz_
	case Q == 0x1 && R == 0x1 && N == 0x0 && U == 0x0:
		// encoding_sqshlr_z_p_zz_
	case Q == 0x1 && R == 0x1 && N == 0x0 && U == 0x1:
		// encoding_uqshlr_z_p_zz_
	case Q == 0x1 && R == 0x1 && N == 0x1 && U == 0x0:
		// encoding_sqrshlr_z_p_zz_
	case Q == 0x1 && R == 0x1 && N == 0x1 && U == 0x1:
		// encoding_uqrshlr_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Q << 19) | (R << 18) | (N << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x44008000

	return
}

func encode_sve_intx_pred_arith_binary(size, R, S, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || R > 0x1 || S > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case R == 0x0 && S == 0x0 && U == 0x0:
		// encoding_shadd_z_p_zz_
	case R == 0x0 && S == 0x0 && U == 0x1:
		// encoding_uhadd_z_p_zz_
	case R == 0x0 && S == 0x1 && U == 0x0:
		// encoding_shsub_z_p_zz_
	case R == 0x0 && S == 0x1 && U == 0x1:
		// encoding_uhsub_z_p_zz_
	case R == 0x1 && S == 0x0 && U == 0x0:
		// encoding_srhadd_z_p_zz_
	case R == 0x1 && S == 0x0 && U == 0x1:
		// encoding_urhadd_z_p_zz_
	case R == 0x1 && S == 0x1 && U == 0x0:
		// encoding_shsubr_z_p_zz_
	case R == 0x1 && S == 0x1 && U == 0x1:
		// encoding_uhsubr_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (R << 18) | (S << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x44108000

	return
}

func encode_sve_intx_arith_binary_pairs(size, opc, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0 && U == 0x0:
		err = errUnallocated
	case opc == 0x0 && U == 0x1:
		// encoding_addp_z_p_zz_
	case opc == 0x1:
		err = errUnallocated
	case opc == 0x2 && U == 0x0:
		// encoding_smaxp_z_p_zz_
	case opc == 0x2 && U == 0x1:
		// encoding_umaxp_z_p_zz_
	case opc == 0x3 && U == 0x0:
		// encoding_sminp_z_p_zz_
	case opc == 0x3 && U == 0x1:
		// encoding_uminp_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x4410a000

	return
}

func encode_sve_intx_pred_arith_binary_sat(size, op, S, U, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || S > 0x1 || U > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case op == 0x0 && S == 0x0 && U == 0x0:
		// encoding_sqadd_z_p_zz_
	case op == 0x0 && S == 0x0 && U == 0x1:
		// encoding_uqadd_z_p_zz_
	case op == 0x0 && S == 0x1 && U == 0x0:
		// encoding_sqsub_z_p_zz_
	case op == 0x0 && S == 0x1 && U == 0x1:
		// encoding_uqsub_z_p_zz_
	case op == 0x1 && S == 0x0 && U == 0x0:
		// encoding_suqadd_z_p_zz_
	case op == 0x1 && S == 0x0 && U == 0x1:
		// encoding_usqadd_z_p_zz_
	case op == 0x1 && S == 0x1 && U == 0x0:
		// encoding_sqsubr_z_p_zz_
	case op == 0x1 && S == 0x1 && U == 0x1:
		// encoding_uqsubr_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 18) | (S << 17) | (U << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x44188000

	return
}

func encode_sve_intx_clamp(size, Zm, U, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || U > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case U == 0x0:
		// encoding_sclamp_z_zz_
	case U == 0x1:
		// encoding_uclamp_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (U << 10) | (Zn << 5) | Zd
	ins |= 0x4400c000

	return
}

func encode_sve_intx_dot_by_indexed_elem(size, opc, U, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || U > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2 && U == 0x0:
		// encoding_sdot_z_zzzi_s
	case size == 0x2 && U == 0x1:
		// encoding_udot_z_zzzi_s
	case size == 0x3 && U == 0x0:
		// encoding_sdot_z_zzzi_d
	case size == 0x3 && U == 0x1:
		// encoding_udot_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (U << 10) | (Zn << 5) | Zda
	ins |= 0x44200000

	return
}

func encode_sve_intx_mla_by_indexed_elem(size, opc, S, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || S > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && S == 0x0:
		// encoding_mla_z_zzzi_h
	case (size&0x2) == 0x0 && S == 0x1:
		// encoding_mls_z_zzzi_h
	case size == 0x2 && S == 0x0:
		// encoding_mla_z_zzzi_s
	case size == 0x2 && S == 0x1:
		// encoding_mls_z_zzzi_s
	case size == 0x3 && S == 0x0:
		// encoding_mla_z_zzzi_d
	case size == 0x3 && S == 0x1:
		// encoding_mls_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (S << 10) | (Zn << 5) | Zda
	ins |= 0x44200800

	return
}

func encode_sve_intx_qrdmlah_by_indexed_elem(size, opc, S, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || S > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && S == 0x0:
		// encoding_sqrdmlah_z_zzzi_h
	case (size&0x2) == 0x0 && S == 0x1:
		// encoding_sqrdmlsh_z_zzzi_h
	case size == 0x2 && S == 0x0:
		// encoding_sqrdmlah_z_zzzi_s
	case size == 0x2 && S == 0x1:
		// encoding_sqrdmlsh_z_zzzi_s
	case size == 0x3 && S == 0x0:
		// encoding_sqrdmlah_z_zzzi_d
	case size == 0x3 && S == 0x1:
		// encoding_sqrdmlsh_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (S << 10) | (Zn << 5) | Zda
	ins |= 0x44201000

	return
}

func encode_sve_intx_mixed_dot_by_indexed_elem(size, opc, U, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || U > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2 && U == 0x0:
		// encoding_usdot_z_zzzi_s
	case size == 0x2 && U == 0x1:
		// encoding_sudot_z_zzzi_s
	case size == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (U << 10) | (Zn << 5) | Zda
	ins |= 0x44201800

	return
}

func encode_sve_intx_qdmla_long_by_indexed_elem(size, opc, S, il, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || S > 0x1 || il > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2 && S == 0x0 && T == 0x0:
		// encoding_sqdmlalb_z_zzzi_s
	case size == 0x2 && S == 0x0 && T == 0x1:
		// encoding_sqdmlalt_z_zzzi_s
	case size == 0x2 && S == 0x1 && T == 0x0:
		// encoding_sqdmlslb_z_zzzi_s
	case size == 0x2 && S == 0x1 && T == 0x1:
		// encoding_sqdmlslt_z_zzzi_s
	case size == 0x3 && S == 0x0 && T == 0x0:
		// encoding_sqdmlalb_z_zzzi_d
	case size == 0x3 && S == 0x0 && T == 0x1:
		// encoding_sqdmlalt_z_zzzi_d
	case size == 0x3 && S == 0x1 && T == 0x0:
		// encoding_sqdmlslb_z_zzzi_d
	case size == 0x3 && S == 0x1 && T == 0x1:
		// encoding_sqdmlslt_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (S << 12) | (il << 11) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x44202000

	return
}

func encode_sve_intx_cdot_by_indexed_elem(size, opc, rot, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || rot > 0x3 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2:
		// encoding_cdot_z_zzzi_s
	case size == 0x3:
		// encoding_cdot_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (rot << 10) | (Zn << 5) | Zda
	ins |= 0x44204000

	return
}

func encode_sve_intx_cmla_by_indexed_elem(size, opc, rot, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || rot > 0x3 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2:
		// encoding_cmla_z_zzzi_h
	case size == 0x3:
		// encoding_cmla_z_zzzi_s
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (rot << 10) | (Zn << 5) | Zda
	ins |= 0x44206000

	return
}

func encode_sve_intx_qrdcmla_by_indexed_elem(size, opc, rot, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || rot > 0x3 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2:
		// encoding_sqrdcmlah_z_zzzi_h
	case size == 0x3:
		// encoding_sqrdcmlah_z_zzzi_s
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (rot << 10) | (Zn << 5) | Zda
	ins |= 0x44207000

	return
}

func encode_sve_intx_mla_long_by_indexed_elem(size, opc, S, U, il, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || S > 0x1 || U > 0x1 || il > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2 && S == 0x0 && U == 0x0 && T == 0x0:
		// encoding_smlalb_z_zzzi_s
	case size == 0x2 && S == 0x0 && U == 0x0 && T == 0x1:
		// encoding_smlalt_z_zzzi_s
	case size == 0x2 && S == 0x0 && U == 0x1 && T == 0x0:
		// encoding_umlalb_z_zzzi_s
	case size == 0x2 && S == 0x0 && U == 0x1 && T == 0x1:
		// encoding_umlalt_z_zzzi_s
	case size == 0x2 && S == 0x1 && U == 0x0 && T == 0x0:
		// encoding_smlslb_z_zzzi_s
	case size == 0x2 && S == 0x1 && U == 0x0 && T == 0x1:
		// encoding_smlslt_z_zzzi_s
	case size == 0x2 && S == 0x1 && U == 0x1 && T == 0x0:
		// encoding_umlslb_z_zzzi_s
	case size == 0x2 && S == 0x1 && U == 0x1 && T == 0x1:
		// encoding_umlslt_z_zzzi_s
	case size == 0x3 && S == 0x0 && U == 0x0 && T == 0x0:
		// encoding_smlalb_z_zzzi_d
	case size == 0x3 && S == 0x0 && U == 0x0 && T == 0x1:
		// encoding_smlalt_z_zzzi_d
	case size == 0x3 && S == 0x0 && U == 0x1 && T == 0x0:
		// encoding_umlalb_z_zzzi_d
	case size == 0x3 && S == 0x0 && U == 0x1 && T == 0x1:
		// encoding_umlalt_z_zzzi_d
	case size == 0x3 && S == 0x1 && U == 0x0 && T == 0x0:
		// encoding_smlslb_z_zzzi_d
	case size == 0x3 && S == 0x1 && U == 0x0 && T == 0x1:
		// encoding_smlslt_z_zzzi_d
	case size == 0x3 && S == 0x1 && U == 0x1 && T == 0x0:
		// encoding_umlslb_z_zzzi_d
	case size == 0x3 && S == 0x1 && U == 0x1 && T == 0x1:
		// encoding_umlslt_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (S << 13) | (U << 12) | (il << 11) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x44208000

	return
}

func encode_sve_intx_mul_long_by_indexed_elem(size, opc, U, il, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || U > 0x1 || il > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2 && U == 0x0 && T == 0x0:
		// encoding_smullb_z_zzi_s
	case size == 0x2 && U == 0x0 && T == 0x1:
		// encoding_smullt_z_zzi_s
	case size == 0x2 && U == 0x1 && T == 0x0:
		// encoding_umullb_z_zzi_s
	case size == 0x2 && U == 0x1 && T == 0x1:
		// encoding_umullt_z_zzi_s
	case size == 0x3 && U == 0x0 && T == 0x0:
		// encoding_smullb_z_zzi_d
	case size == 0x3 && U == 0x0 && T == 0x1:
		// encoding_smullt_z_zzi_d
	case size == 0x3 && U == 0x1 && T == 0x0:
		// encoding_umullb_z_zzi_d
	case size == 0x3 && U == 0x1 && T == 0x1:
		// encoding_umullt_z_zzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (U << 12) | (il << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x4420c000

	return
}

func encode_sve_intx_qdmul_long_by_indexed_elem(size, opc, il, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || il > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2 && T == 0x0:
		// encoding_sqdmullb_z_zzi_s
	case size == 0x2 && T == 0x1:
		// encoding_sqdmullt_z_zzi_s
	case size == 0x3 && T == 0x0:
		// encoding_sqdmullb_z_zzi_d
	case size == 0x3 && T == 0x1:
		// encoding_sqdmullt_z_zzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (il << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x4420e000

	return
}

func encode_sve_intx_qdmulh_by_indexed_elem(size, opc, R, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || R > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && R == 0x0:
		// encoding_sqdmulh_z_zzi_h
	case (size&0x2) == 0x0 && R == 0x1:
		// encoding_sqrdmulh_z_zzi_h
	case size == 0x2 && R == 0x0:
		// encoding_sqdmulh_z_zzi_s
	case size == 0x2 && R == 0x1:
		// encoding_sqrdmulh_z_zzi_s
	case size == 0x3 && R == 0x0:
		// encoding_sqdmulh_z_zzi_d
	case size == 0x3 && R == 0x1:
		// encoding_sqrdmulh_z_zzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (R << 10) | (Zn << 5) | Zd
	ins |= 0x4420f000

	return
}

func encode_sve_intx_mul_by_indexed_elem(size, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		// encoding_mul_z_zzi_h
	case size == 0x2:
		// encoding_mul_z_zzi_s
	case size == 0x3:
		// encoding_mul_z_zzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Zn << 5) | Zd
	ins |= 0x4420f800

	return
}

func encode_sve_intx_cons_arith_long(size, Zm, op, S, U, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || S > 0x1 || U > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case op == 0x0 && S == 0x0 && U == 0x0 && T == 0x0:
		// encoding_saddlb_z_zz_
	case op == 0x0 && S == 0x0 && U == 0x0 && T == 0x1:
		// encoding_saddlt_z_zz_
	case op == 0x0 && S == 0x0 && U == 0x1 && T == 0x0:
		// encoding_uaddlb_z_zz_
	case op == 0x0 && S == 0x0 && U == 0x1 && T == 0x1:
		// encoding_uaddlt_z_zz_
	case op == 0x0 && S == 0x1 && U == 0x0 && T == 0x0:
		// encoding_ssublb_z_zz_
	case op == 0x0 && S == 0x1 && U == 0x0 && T == 0x1:
		// encoding_ssublt_z_zz_
	case op == 0x0 && S == 0x1 && U == 0x1 && T == 0x0:
		// encoding_usublb_z_zz_
	case op == 0x0 && S == 0x1 && U == 0x1 && T == 0x1:
		// encoding_usublt_z_zz_
	case op == 0x1 && S == 0x0:
		err = errUnallocated
	case op == 0x1 && S == 0x1 && U == 0x0 && T == 0x0:
		// encoding_sabdlb_z_zz_
	case op == 0x1 && S == 0x1 && U == 0x0 && T == 0x1:
		// encoding_sabdlt_z_zz_
	case op == 0x1 && S == 0x1 && U == 0x1 && T == 0x0:
		// encoding_uabdlb_z_zz_
	case op == 0x1 && S == 0x1 && U == 0x1 && T == 0x1:
		// encoding_uabdlt_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 13) | (S << 12) | (U << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x45000000

	return
}

func encode_sve_intx_cons_arith_wide(size, Zm, S, U, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || U > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case S == 0x0 && U == 0x0 && T == 0x0:
		// encoding_saddwb_z_zz_
	case S == 0x0 && U == 0x0 && T == 0x1:
		// encoding_saddwt_z_zz_
	case S == 0x0 && U == 0x1 && T == 0x0:
		// encoding_uaddwb_z_zz_
	case S == 0x0 && U == 0x1 && T == 0x1:
		// encoding_uaddwt_z_zz_
	case S == 0x1 && U == 0x0 && T == 0x0:
		// encoding_ssubwb_z_zz_
	case S == 0x1 && U == 0x0 && T == 0x1:
		// encoding_ssubwt_z_zz_
	case S == 0x1 && U == 0x1 && T == 0x0:
		// encoding_usubwb_z_zz_
	case S == 0x1 && U == 0x1 && T == 0x1:
		// encoding_usubwt_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 12) | (U << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x45004000

	return
}

func encode_sve_intx_cons_mul_long(size, Zm, op, U, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || U > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case op == 0x0 && U == 0x0 && T == 0x0:
		// encoding_sqdmullb_z_zz_
	case op == 0x0 && U == 0x0 && T == 0x1:
		// encoding_sqdmullt_z_zz_
	case op == 0x0 && U == 0x1 && T == 0x0:
		// encoding_pmullb_z_zz_
	case op == 0x0 && U == 0x1 && T == 0x1:
		// encoding_pmullt_z_zz_
	case op == 0x1 && U == 0x0 && T == 0x0:
		// encoding_smullb_z_zz_
	case op == 0x1 && U == 0x0 && T == 0x1:
		// encoding_smullt_z_zz_
	case op == 0x1 && U == 0x1 && T == 0x0:
		// encoding_umullb_z_zz_
	case op == 0x1 && U == 0x1 && T == 0x1:
		// encoding_umullt_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 12) | (U << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x45006000

	return
}

func encode_sve_intx_shift_long(tszh, tszl, imm3, U, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x1 || tszl > 0x3 || imm3 > 0x7 || U > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case U == 0x0 && T == 0x0:
		// encoding_sshllb_z_zi_
	case U == 0x0 && T == 0x1:
		// encoding_sshllt_z_zi_
	case U == 0x1 && T == 0x0:
		// encoding_ushllb_z_zi_
	case U == 0x1 && T == 0x1:
		// encoding_ushllt_z_zi_
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (tszl << 19) | (imm3 << 16) | (U << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x4500a000

	return
}

func encode_sve_intx_clong(size, Zm, S, tb, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || tb > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case S == 0x0 && tb == 0x0:
		// encoding_saddlbt_z_zz_
	case S == 0x0 && tb == 0x1:
		err = errUnallocated
	case S == 0x1 && tb == 0x0:
		// encoding_ssublbt_z_zz_
	case S == 0x1 && tb == 0x1:
		// encoding_ssubltb_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 11) | (tb << 10) | (Zn << 5) | Zd
	ins |= 0x45008000

	return
}

func encode_sve_intx_eorx(size, Zm, tb, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || tb > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case tb == 0x0:
		// encoding_eorbt_z_zz_
	case tb == 0x1:
		// encoding_eortb_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (tb << 10) | (Zn << 5) | Zd
	ins |= 0x45009000

	return
}

func encode_sve_intx_mmla(uns, Zm, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case uns > 0x3 || Zm > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case uns == 0x0:
		// encoding_smmla_z_zzz_
	case uns == 0x1:
		err = errUnallocated
	case uns == 0x2:
		// encoding_usmmla_z_zzz_
	case uns == 0x3:
		// encoding_ummla_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (uns << 22) | (Zm << 16) | (Zn << 5) | Zd
	ins |= 0x45009800

	return
}

func encode_sve_intx_perm_bit(size, Zm, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x3 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_bext_z_zz_
	case opc == 0x1:
		// encoding_bdep_z_zz_
	case opc == 0x2:
		// encoding_bgrp_z_zz_
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x4500b000

	return
}

func encode_sve_intx_cadd(size, op, rot, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || rot > 0x1 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_cadd_z_zz_
	case op == 0x1:
		// encoding_sqcadd_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 16) | (rot << 10) | (Zm << 5) | Zdn
	ins |= 0x4500d800

	return
}

func encode_sve_intx_aba_long(size, Zm, U, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || U > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case U == 0x0 && T == 0x0:
		// encoding_sabalb_z_zzz_
	case U == 0x0 && T == 0x1:
		// encoding_sabalt_z_zzz_
	case U == 0x1 && T == 0x0:
		// encoding_uabalb_z_zzz_
	case U == 0x1 && T == 0x1:
		// encoding_uabalt_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (U << 11) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x4500c000

	return
}

func encode_sve_intx_adc_long(size, Zm, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && T == 0x0:
		// encoding_adclb_z_zzz_
	case (size&0x2) == 0x0 && T == 0x1:
		// encoding_adclt_z_zzz_
	case (size&0x2) == 0x2 && T == 0x0:
		// encoding_sbclb_z_zzz_
	case (size&0x2) == 0x2 && T == 0x1:
		// encoding_sbclt_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x4500d000

	return
}

func encode_sve_intx_sra(tszh, tszl, imm3, R, U, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x3 || tszl > 0x3 || imm3 > 0x7 || R > 0x1 || U > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case R == 0x0 && U == 0x0:
		// encoding_ssra_z_zi_
	case R == 0x0 && U == 0x1:
		// encoding_usra_z_zi_
	case R == 0x1 && U == 0x0:
		// encoding_srsra_z_zi_
	case R == 0x1 && U == 0x1:
		// encoding_ursra_z_zi_
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (tszl << 19) | (imm3 << 16) | (R << 11) | (U << 10) | (Zn << 5) | Zda
	ins |= 0x4500e000

	return
}

func encode_sve_intx_shift_insert(tszh, tszl, imm3, op, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x3 || tszl > 0x3 || imm3 > 0x7 || op > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_sri_z_zzi_
	case op == 0x1:
		// encoding_sli_z_zzi_
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (tszl << 19) | (imm3 << 16) | (op << 10) | (Zn << 5) | Zd
	ins |= 0x4500f000

	return
}

func encode_sve_intx_aba(size, Zm, U, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || U > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case U == 0x0:
		// encoding_saba_z_zzz_
	case U == 0x1:
		// encoding_uaba_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (U << 10) | (Zn << 5) | Zda
	ins |= 0x4500f800

	return
}

func encode_sve_intx_extract_narrow(tszh, tszl, opc, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x1 || tszl > 0x3 || opc > 0x3 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0 && T == 0x0:
		// encoding_sqxtnb_z_zz_
	case opc == 0x0 && T == 0x1:
		// encoding_sqxtnt_z_zz_
	case opc == 0x1 && T == 0x0:
		// encoding_uqxtnb_z_zz_
	case opc == 0x1 && T == 0x1:
		// encoding_uqxtnt_z_zz_
	case opc == 0x2 && T == 0x0:
		// encoding_sqxtunb_z_zz_
	case opc == 0x2 && T == 0x1:
		// encoding_sqxtunt_z_zz_
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (tszl << 19) | (opc << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x45204000

	return
}

func encode_sve_intx_shift_narrow(tszh, tszl, imm3, op, U, R, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case tszh > 0x1 || tszl > 0x3 || imm3 > 0x7 || op > 0x1 || U > 0x1 || R > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case op == 0x0 && U == 0x0 && R == 0x0 && T == 0x0:
		// encoding_sqshrunb_z_zi_
	case op == 0x0 && U == 0x0 && R == 0x0 && T == 0x1:
		// encoding_sqshrunt_z_zi_
	case op == 0x0 && U == 0x0 && R == 0x1 && T == 0x0:
		// encoding_sqrshrunb_z_zi_
	case op == 0x0 && U == 0x0 && R == 0x1 && T == 0x1:
		// encoding_sqrshrunt_z_zi_
	case op == 0x0 && U == 0x1 && R == 0x0 && T == 0x0:
		// encoding_shrnb_z_zi_
	case op == 0x0 && U == 0x1 && R == 0x0 && T == 0x1:
		// encoding_shrnt_z_zi_
	case op == 0x0 && U == 0x1 && R == 0x1 && T == 0x0:
		// encoding_rshrnb_z_zi_
	case op == 0x0 && U == 0x1 && R == 0x1 && T == 0x1:
		// encoding_rshrnt_z_zi_
	case op == 0x1 && U == 0x0 && R == 0x0 && T == 0x0:
		// encoding_sqshrnb_z_zi_
	case op == 0x1 && U == 0x0 && R == 0x0 && T == 0x1:
		// encoding_sqshrnt_z_zi_
	case op == 0x1 && U == 0x0 && R == 0x1 && T == 0x0:
		// encoding_sqrshrnb_z_zi_
	case op == 0x1 && U == 0x0 && R == 0x1 && T == 0x1:
		// encoding_sqrshrnt_z_zi_
	case op == 0x1 && U == 0x1 && R == 0x0 && T == 0x0:
		// encoding_uqshrnb_z_zi_
	case op == 0x1 && U == 0x1 && R == 0x0 && T == 0x1:
		// encoding_uqshrnt_z_zi_
	case op == 0x1 && U == 0x1 && R == 0x1 && T == 0x0:
		// encoding_uqrshrnb_z_zi_
	case op == 0x1 && U == 0x1 && R == 0x1 && T == 0x1:
		// encoding_uqrshrnt_z_zi_
	default:
		err = errUnmatched
	}

	ins |= (tszh << 22) | (tszl << 19) | (imm3 << 16) | (op << 13) | (U << 12) | (R << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x45200000

	return
}

func encode_sve_intx_arith_narrow(size, Zm, S, R, T, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || S > 0x1 || R > 0x1 || T > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case S == 0x0 && R == 0x0 && T == 0x0:
		// encoding_addhnb_z_zz_
	case S == 0x0 && R == 0x0 && T == 0x1:
		// encoding_addhnt_z_zz_
	case S == 0x0 && R == 0x1 && T == 0x0:
		// encoding_raddhnb_z_zz_
	case S == 0x0 && R == 0x1 && T == 0x1:
		// encoding_raddhnt_z_zz_
	case S == 0x1 && R == 0x0 && T == 0x0:
		// encoding_subhnb_z_zz_
	case S == 0x1 && R == 0x0 && T == 0x1:
		// encoding_subhnt_z_zz_
	case S == 0x1 && R == 0x1 && T == 0x0:
		// encoding_rsubhnb_z_zz_
	case S == 0x1 && R == 0x1 && T == 0x1:
		// encoding_rsubhnt_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (S << 12) | (R << 11) | (T << 10) | (Zn << 5) | Zd
	ins |= 0x45206000

	return
}

func encode_sve_intx_match(size, Zm, Pg, Zn, op, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || Pg > 0x7 || Zn > 0x1f || op > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0:
		// encoding_match_p_p_zz_
	case op == 0x1:
		// encoding_nmatch_p_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (Pg << 10) | (Zn << 5) | (op << 4) | Pd
	ins |= 0x45208000

	return
}

func encode_sve_intx_histseg(size, Zm, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_histseg_z_zz_
	}

	ins |= (size << 22) | (Zm << 16) | (Zn << 5) | Zd
	ins |= 0x4520a000

	return
}

func encode_sve_intx_histcnt(size, Zm, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	default:
		// encoding_histcnt_z_p_zz_
	}

	ins |= (size << 22) | (Zm << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x4520c000

	return
}

func encode_sve_crypto_unary(size, op, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || Zdn > 0x1f:
		err = errOverflow
	case size == 0x0 && op == 0x0:
		// encoding_aesmc_z_z_
	case size == 0x0 && op == 0x1:
		// encoding_aesimc_z_z_
	case size == 0x1:
		err = errUnallocated
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 10) | Zdn
	ins |= 0x4520e000

	return
}

func encode_sve_crypto_binary_dest(size, op, o2, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || op > 0x1 || o2 > 0x1 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case size == 0x0 && op == 0x0 && o2 == 0x0:
		// encoding_aese_z_zz_
	case size == 0x0 && op == 0x0 && o2 == 0x1:
		// encoding_aesd_z_zz_
	case size == 0x0 && op == 0x1 && o2 == 0x0:
		// encoding_sm4e_z_zz_
	case size == 0x0 && op == 0x1 && o2 == 0x1:
		err = errUnallocated
	case size == 0x1:
		err = errUnallocated
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (op << 16) | (o2 << 10) | (Zm << 5) | Zdn
	ins |= 0x4522e000

	return
}

func encode_sve_crypto_binary_const(size, Zm, op, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case size == 0x0 && op == 0x0:
		// encoding_sm4ekey_z_zz_
	case size == 0x0 && op == 0x1:
		// encoding_rax1_z_zz_
	case size == 0x1:
		err = errUnallocated
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 10) | (Zn << 5) | Zd
	ins |= 0x4520f000

	return
}

func encode_sve_fp_fcmla(size, Zm, rot, Pg, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || rot > 0x3 || Pg > 0x7 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	default:
		// encoding_fcmla_z_p_zzz_
	}

	ins |= (size << 22) | (Zm << 16) | (rot << 13) | (Pg << 10) | (Zn << 5) | Zda
	ins |= 0x64000000

	return
}

func encode_sve_fp_fcadd(size, rot, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || rot > 0x1 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_fcadd_z_p_zz_
	}

	ins |= (size << 22) | (rot << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x64008000

	return
}

func encode_sve_fp_fcvt2(opc, opc2, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || opc2 > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (opc&0x1) == 0x0 && opc2 == 0x3:
		err = errUnallocated
	case opc == 0x0 && (opc2&0x2) == 0x0:
		err = errUnallocated
	case opc == 0x0 && opc2 == 0x2:
		// encoding_fcvtxnt_z_p_z_d2s
	case opc == 0x1:
		err = errUnallocated
	case opc == 0x2 && opc2 == 0x0:
		// encoding_fcvtnt_z_p_z_s2h
	case opc == 0x2 && opc2 == 0x1:
		// encoding_fcvtlt_z_p_z_h2s
	case opc == 0x2 && opc2 == 0x2:
		// encoding_bfcvtnt_z_p_z_s2bf
	case opc == 0x3 && (opc2&0x2) == 0x0:
		err = errUnallocated
	case opc == 0x3 && opc2 == 0x2:
		// encoding_fcvtnt_z_p_z_d2s
	case opc == 0x3 && opc2 == 0x3:
		// encoding_fcvtlt_z_p_z_s2d
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (opc2 << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x6408a000

	return
}

func encode_sve_fp_pairwise(size, opc, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_faddp_z_p_zz_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x6) == 0x2:
		err = errUnallocated
	case opc == 0x4:
		// encoding_fmaxnmp_z_p_zz_
	case opc == 0x5:
		// encoding_fminnmp_z_p_zz_
	case opc == 0x6:
		// encoding_fmaxp_z_p_zz_
	case opc == 0x7:
		// encoding_fminp_z_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x64108000

	return
}

func encode_sve_fp_fma_by_indexed_elem(size, opc, op, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || op > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && op == 0x0:
		// encoding_fmla_z_zzzi_h
	case (size&0x2) == 0x0 && op == 0x1:
		// encoding_fmls_z_zzzi_h
	case size == 0x2 && op == 0x0:
		// encoding_fmla_z_zzzi_s
	case size == 0x2 && op == 0x1:
		// encoding_fmls_z_zzzi_s
	case size == 0x3 && op == 0x0:
		// encoding_fmla_z_zzzi_d
	case size == 0x3 && op == 0x1:
		// encoding_fmls_z_zzzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (op << 10) | (Zn << 5) | Zda
	ins |= 0x64200000

	return
}

func encode_sve_fp_fcmla_by_indexed_elem(size, opc, rot, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || rot > 0x3 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		err = errUnallocated
	case size == 0x2:
		// encoding_fcmla_z_zzzi_h
	case size == 0x3:
		// encoding_fcmla_z_zzzi_s
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (rot << 10) | (Zn << 5) | Zda
	ins |= 0x64201000

	return
}

func encode_sve_fp_fmul_by_indexed_elem(size, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x1f || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (size & 0x2) == 0x0:
		// encoding_fmul_z_zzi_h
	case size == 0x2:
		// encoding_fmul_z_zzi_s
	case size == 0x3:
		// encoding_fmul_z_zzi_d
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Zn << 5) | Zd
	ins |= 0x64202000

	return
}

func encode_sve_fp_fdot_by_indexed_elem(op, i2, Zm, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || i2 > 0x3 || Zm > 0x7 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case op == 0x0:
		err = errUnallocated
	case op == 0x1:
		// encoding_bfdot_z_zzzi_
	default:
		err = errUnmatched
	}

	ins |= (op << 22) | (i2 << 19) | (Zm << 16) | (Zn << 5) | Zda
	ins |= 0x64204000

	return
}

func encode_sve_fp_fma_long_by_indexed_elem(o2, i3h, Zm, op, i3l, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case o2 > 0x1 || i3h > 0x3 || Zm > 0x7 || op > 0x1 || i3l > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case o2 == 0x0 && op == 0x0 && T == 0x0:
		// encoding_fmlalb_z_zzzi_s
	case o2 == 0x0 && op == 0x0 && T == 0x1:
		// encoding_fmlalt_z_zzzi_s
	case o2 == 0x0 && op == 0x1 && T == 0x0:
		// encoding_fmlslb_z_zzzi_s
	case o2 == 0x0 && op == 0x1 && T == 0x1:
		// encoding_fmlslt_z_zzzi_s
	case o2 == 0x1 && op == 0x0 && T == 0x0:
		// encoding_bfmlalb_z_zzzi_
	case o2 == 0x1 && op == 0x0 && T == 0x1:
		// encoding_bfmlalt_z_zzzi_
	case o2 == 0x1 && op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (o2 << 22) | (i3h << 19) | (Zm << 16) | (op << 13) | (i3l << 11) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x64a04000

	return
}

func encode_sve_fp_fdot(op, Zm, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || Zm > 0x1f || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case op == 0x0:
		err = errUnallocated
	case op == 0x1:
		// encoding_bfdot_z_zzz_
	default:
		err = errUnmatched
	}

	ins |= (op << 22) | (Zm << 16) | (Zn << 5) | Zda
	ins |= 0x64208000

	return
}

func encode_sve_fp_fma_long(o2, Zm, op, T, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case o2 > 0x1 || Zm > 0x1f || op > 0x1 || T > 0x1 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case o2 == 0x0 && op == 0x0 && T == 0x0:
		// encoding_fmlalb_z_zzz_
	case o2 == 0x0 && op == 0x0 && T == 0x1:
		// encoding_fmlalt_z_zzz_
	case o2 == 0x0 && op == 0x1 && T == 0x0:
		// encoding_fmlslb_z_zzz_
	case o2 == 0x0 && op == 0x1 && T == 0x1:
		// encoding_fmlslt_z_zzz_
	case o2 == 0x1 && op == 0x0 && T == 0x0:
		// encoding_bfmlalb_z_zzz_
	case o2 == 0x1 && op == 0x0 && T == 0x1:
		// encoding_bfmlalt_z_zzz_
	case o2 == 0x1 && op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (o2 << 22) | (Zm << 16) | (op << 13) | (T << 10) | (Zn << 5) | Zda
	ins |= 0x64a08000

	return
}

func encode_sve_fp_fmmla(opc, Zm, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || Zm > 0x1f || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case opc == 0x0:
		err = errUnallocated
	case opc == 0x1:
		// encoding_bfmmla_z_zzz_
	case opc == 0x2:
		// encoding_fmmla_z_zzz_s
	case opc == 0x3:
		// encoding_fmmla_z_zzz_d
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (Zm << 16) | (Zn << 5) | Zda
	ins |= 0x6420e400

	return
}

func encode_sve_fp_3op_p_pd(size, Zm, op, o2, Pg, Zn, o3, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || op > 0x1 || o2 > 0x1 || Pg > 0x7 || Zn > 0x1f || o3 > 0x1 || Pd > 0xf:
		err = errOverflow
	case op == 0x0 && o2 == 0x0 && o3 == 0x0:
		// encoding_fcmge_p_p_zz_
	case op == 0x0 && o2 == 0x0 && o3 == 0x1:
		// encoding_fcmgt_p_p_zz_
	case op == 0x0 && o2 == 0x1 && o3 == 0x0:
		// encoding_fcmeq_p_p_zz_
	case op == 0x0 && o2 == 0x1 && o3 == 0x1:
		// encoding_fcmne_p_p_zz_
	case op == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_fcmuo_p_p_zz_
	case op == 0x1 && o2 == 0x0 && o3 == 0x1:
		// encoding_facge_p_p_zz_
	case op == 0x1 && o2 == 0x1 && o3 == 0x0:
		err = errUnallocated
	case op == 0x1 && o2 == 0x1 && o3 == 0x1:
		// encoding_facgt_p_p_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (op << 15) | (o2 << 13) | (Pg << 10) | (Zn << 5) | (o3 << 4) | Pd
	ins |= 0x65004000

	return
}

func encode_sve_fp_3op_u_zd(size, Zm, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fadd_z_zz_
	case opc == 0x1:
		// encoding_fsub_z_zz_
	case opc == 0x2:
		// encoding_fmul_z_zz_
	case opc == 0x3:
		// encoding_ftsmul_z_zz_
	case (opc & 0x6) == 0x4:
		err = errUnallocated
	case opc == 0x6:
		// encoding_frecps_z_zz_
	case opc == 0x7:
		// encoding_frsqrts_z_zz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 10) | (Zn << 5) | Zd
	ins |= 0x65000000

	return
}

func encode_sve_fp_2op_p_zds(size, opc, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0xf || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fadd_z_p_zz_
	case opc == 0x1:
		// encoding_fsub_z_p_zz_
	case opc == 0x2:
		// encoding_fmul_z_p_zz_
	case opc == 0x3:
		// encoding_fsubr_z_p_zz_
	case opc == 0x4:
		// encoding_fmaxnm_z_p_zz_
	case opc == 0x5:
		// encoding_fminnm_z_p_zz_
	case opc == 0x6:
		// encoding_fmax_z_p_zz_
	case opc == 0x7:
		// encoding_fmin_z_p_zz_
	case opc == 0x8:
		// encoding_fabd_z_p_zz_
	case opc == 0x9:
		// encoding_fscale_z_p_zz_
	case opc == 0xa:
		// encoding_fmulx_z_p_zz_
	case opc == 0xb:
		err = errUnallocated
	case opc == 0xc:
		// encoding_fdivr_z_p_zz_
	case opc == 0xd:
		// encoding_fdiv_z_p_zz_
	case (opc & 0xe) == 0xe:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x65008000

	return
}

func encode_sve_fp_ftmad(size, imm3, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || imm3 > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	default:
		// encoding_ftmad_z_zzi_
	}

	ins |= (size << 22) | (imm3 << 16) | (Zm << 5) | Zdn
	ins |= 0x65108000

	return
}

func encode_sve_fp_2op_i_p_zds(size, opc, Pg, i1, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || i1 > 0x1 || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fadd_z_p_zs_
	case opc == 0x1:
		// encoding_fsub_z_p_zs_
	case opc == 0x2:
		// encoding_fmul_z_p_zs_
	case opc == 0x3:
		// encoding_fsubr_z_p_zs_
	case opc == 0x4:
		// encoding_fmaxnm_z_p_zs_
	case opc == 0x5:
		// encoding_fminnm_z_p_zs_
	case opc == 0x6:
		// encoding_fmax_z_p_zs_
	case opc == 0x7:
		// encoding_fmin_z_p_zs_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (i1 << 5) | Zdn
	ins |= 0x65188000

	return
}

func encode_sve_fp_2op_p_zd_a(size, opc, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_frintn_z_p_z_
	case opc == 0x1:
		// encoding_frintp_z_p_z_
	case opc == 0x2:
		// encoding_frintm_z_p_z_
	case opc == 0x3:
		// encoding_frintz_z_p_z_
	case opc == 0x4:
		// encoding_frinta_z_p_z_
	case opc == 0x5:
		err = errUnallocated
	case opc == 0x6:
		// encoding_frintx_z_p_z_
	case opc == 0x7:
		// encoding_frinti_z_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x6500a000

	return
}

func encode_sve_fp_2op_p_zd_b_0(opc, opc2, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || opc2 > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (opc&0x1) == 0x0 && opc2 == 0x3:
		err = errUnallocated
	case opc == 0x0 && (opc2&0x2) == 0x0:
		err = errUnallocated
	case opc == 0x0 && opc2 == 0x2:
		// encoding_fcvtx_z_p_z_d2s
	case opc == 0x1:
		err = errUnallocated
	case opc == 0x2 && opc2 == 0x0:
		// encoding_fcvt_z_p_z_s2h
	case opc == 0x2 && opc2 == 0x1:
		// encoding_fcvt_z_p_z_h2s
	case opc == 0x2 && opc2 == 0x2:
		// encoding_bfcvt_z_p_z_s2bf
	case opc == 0x3 && opc2 == 0x0:
		// encoding_fcvt_z_p_z_d2h
	case opc == 0x3 && opc2 == 0x1:
		// encoding_fcvt_z_p_z_h2d
	case opc == 0x3 && opc2 == 0x2:
		// encoding_fcvt_z_p_z_d2s
	case opc == 0x3 && opc2 == 0x3:
		// encoding_fcvt_z_p_z_s2d
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (opc2 << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x6508a000

	return
}

func encode_sve_fp_2op_p_zd_b_1(size, opc, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_frecpx_z_p_z_
	case opc == 0x1:
		// encoding_fsqrt_z_p_z_
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x650ca000

	return
}

func encode_sve_fp_2op_p_zd_c(opc, opc2, U, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || opc2 > 0x3 || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		err = errUnallocated
	case opc == 0x1 && opc2 == 0x0:
		err = errUnallocated
	case opc == 0x1 && opc2 == 0x1 && U == 0x0:
		// encoding_scvtf_z_p_z_h2fp16
	case opc == 0x1 && opc2 == 0x1 && U == 0x1:
		// encoding_ucvtf_z_p_z_h2fp16
	case opc == 0x1 && opc2 == 0x2 && U == 0x0:
		// encoding_scvtf_z_p_z_w2fp16
	case opc == 0x1 && opc2 == 0x2 && U == 0x1:
		// encoding_ucvtf_z_p_z_w2fp16
	case opc == 0x1 && opc2 == 0x3 && U == 0x0:
		// encoding_scvtf_z_p_z_x2fp16
	case opc == 0x1 && opc2 == 0x3 && U == 0x1:
		// encoding_ucvtf_z_p_z_x2fp16
	case opc == 0x2 && (opc2&0x2) == 0x0:
		err = errUnallocated
	case opc == 0x2 && opc2 == 0x2 && U == 0x0:
		// encoding_scvtf_z_p_z_w2s
	case opc == 0x2 && opc2 == 0x2 && U == 0x1:
		// encoding_ucvtf_z_p_z_w2s
	case opc == 0x2 && opc2 == 0x3:
		err = errUnallocated
	case opc == 0x3 && opc2 == 0x0 && U == 0x0:
		// encoding_scvtf_z_p_z_w2d
	case opc == 0x3 && opc2 == 0x0 && U == 0x1:
		// encoding_ucvtf_z_p_z_w2d
	case opc == 0x3 && opc2 == 0x1:
		err = errUnallocated
	case opc == 0x3 && opc2 == 0x2 && U == 0x0:
		// encoding_scvtf_z_p_z_x2s
	case opc == 0x3 && opc2 == 0x2 && U == 0x1:
		// encoding_ucvtf_z_p_z_x2s
	case opc == 0x3 && opc2 == 0x3 && U == 0x0:
		// encoding_scvtf_z_p_z_x2d
	case opc == 0x3 && opc2 == 0x3 && U == 0x1:
		// encoding_ucvtf_z_p_z_x2d
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (opc2 << 17) | (U << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x6510a000

	return
}

func encode_sve_fp_2op_p_zd_d(opc, opc2, U, Pg, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || opc2 > 0x3 || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case opc == 0x0 && U == 0x0:
		// encoding_flogb_z_p_z_
	case opc == 0x0 && U == 0x1:
		err = errUnallocated
	case opc == 0x1 && opc2 == 0x0:
		err = errUnallocated
	case opc == 0x1 && opc2 == 0x1 && U == 0x0:
		// encoding_fcvtzs_z_p_z_fp162h
	case opc == 0x1 && opc2 == 0x1 && U == 0x1:
		// encoding_fcvtzu_z_p_z_fp162h
	case opc == 0x1 && opc2 == 0x2 && U == 0x0:
		// encoding_fcvtzs_z_p_z_fp162w
	case opc == 0x1 && opc2 == 0x2 && U == 0x1:
		// encoding_fcvtzu_z_p_z_fp162w
	case opc == 0x1 && opc2 == 0x3 && U == 0x0:
		// encoding_fcvtzs_z_p_z_fp162x
	case opc == 0x1 && opc2 == 0x3 && U == 0x1:
		// encoding_fcvtzu_z_p_z_fp162x
	case opc == 0x2 && (opc2&0x2) == 0x0:
		err = errUnallocated
	case opc == 0x2 && opc2 == 0x2 && U == 0x0:
		// encoding_fcvtzs_z_p_z_s2w
	case opc == 0x2 && opc2 == 0x2 && U == 0x1:
		// encoding_fcvtzu_z_p_z_s2w
	case opc == 0x2 && opc2 == 0x3:
		err = errUnallocated
	case opc == 0x3 && opc2 == 0x0 && U == 0x0:
		// encoding_fcvtzs_z_p_z_d2w
	case opc == 0x3 && opc2 == 0x0 && U == 0x1:
		// encoding_fcvtzu_z_p_z_d2w
	case opc == 0x3 && opc2 == 0x1:
		err = errUnallocated
	case opc == 0x3 && opc2 == 0x2 && U == 0x0:
		// encoding_fcvtzs_z_p_z_s2x
	case opc == 0x3 && opc2 == 0x2 && U == 0x1:
		// encoding_fcvtzu_z_p_z_s2x
	case opc == 0x3 && opc2 == 0x3 && U == 0x0:
		// encoding_fcvtzs_z_p_z_d2x
	case opc == 0x3 && opc2 == 0x3 && U == 0x1:
		// encoding_fcvtzu_z_p_z_d2x
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (opc2 << 17) | (U << 16) | (Pg << 10) | (Zn << 5) | Zd
	ins |= 0x6518a000

	return
}

func encode_sve_fp_fast_red(size, opc, Pg, Zn, Vd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Pg > 0x7 || Zn > 0x1f || Vd > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_faddv_v_p_z_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x6) == 0x2:
		err = errUnallocated
	case opc == 0x4:
		// encoding_fmaxnmv_v_p_z_
	case opc == 0x5:
		// encoding_fminnmv_v_p_z_
	case opc == 0x6:
		// encoding_fmaxv_v_p_z_
	case opc == 0x7:
		// encoding_fminv_v_p_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zn << 5) | Vd
	ins |= 0x65002000

	return
}

func encode_sve_fp_2op_u_zd(size, opc, Zn, Zd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x7 || Zn > 0x1f || Zd > 0x1f:
		err = errOverflow
	case (opc & 0x4) == 0x0:
		err = errUnallocated
	case (opc & 0x6) == 0x4:
		err = errUnallocated
	case opc == 0x6:
		// encoding_frecpe_z_z_
	case opc == 0x7:
		// encoding_frsqrte_z_z_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Zn << 5) | Zd
	ins |= 0x65083000

	return
}

func encode_sve_fp_2op_p_pd(size, eq, lt, Pg, Zn, ne, Pd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || eq > 0x1 || lt > 0x1 || Pg > 0x7 || Zn > 0x1f || ne > 0x1 || Pd > 0xf:
		err = errOverflow
	case eq == 0x0 && lt == 0x0 && ne == 0x0:
		// encoding_fcmge_p_p_z0_
	case eq == 0x0 && lt == 0x0 && ne == 0x1:
		// encoding_fcmgt_p_p_z0_
	case eq == 0x0 && lt == 0x1 && ne == 0x0:
		// encoding_fcmlt_p_p_z0_
	case eq == 0x0 && lt == 0x1 && ne == 0x1:
		// encoding_fcmle_p_p_z0_
	case eq == 0x1 && ne == 0x1:
		err = errUnallocated
	case eq == 0x1 && lt == 0x0 && ne == 0x0:
		// encoding_fcmeq_p_p_z0_
	case eq == 0x1 && lt == 0x1 && ne == 0x0:
		// encoding_fcmne_p_p_z0_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (eq << 17) | (lt << 16) | (Pg << 10) | (Zn << 5) | (ne << 4) | Pd
	ins |= 0x65102000

	return
}

func encode_sve_fp_2op_p_vd(size, opc, Pg, Zm, Vdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || Pg > 0x7 || Zm > 0x1f || Vdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fadda_v_p_z_
	case opc == 0x1:
		err = errUnallocated
	case (opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opc << 16) | (Pg << 10) | (Zm << 5) | Vdn
	ins |= 0x65182000

	return
}

func encode_sve_fp_3op_p_zds_a(size, Zm, opc, Pg, Zn, Zda uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Zm > 0x1f || opc > 0x3 || Pg > 0x7 || Zn > 0x1f || Zda > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fmla_z_p_zzz_
	case opc == 0x1:
		// encoding_fmls_z_p_zzz_
	case opc == 0x2:
		// encoding_fnmla_z_p_zzz_
	case opc == 0x3:
		// encoding_fnmls_z_p_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Zm << 16) | (opc << 13) | (Pg << 10) | (Zn << 5) | Zda
	ins |= 0x65200000

	return
}

func encode_sve_fp_3op_p_zds_b(size, Za, opc, Pg, Zm, Zdn uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Za > 0x1f || opc > 0x3 || Pg > 0x7 || Zm > 0x1f || Zdn > 0x1f:
		err = errOverflow
	case opc == 0x0:
		// encoding_fmad_z_p_zzz_
	case opc == 0x1:
		// encoding_fmsb_z_p_zzz_
	case opc == 0x2:
		// encoding_fnmad_z_p_zzz_
	case opc == 0x3:
		// encoding_fnmsb_z_p_zzz_
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Za << 16) | (opc << 13) | (Pg << 10) | (Zm << 5) | Zdn
	ins |= 0x65208000

	return
}

func encode_sve_mem_32b_prfm_sv(xs, Zm, msz, Pg, Rn, prfop uint32) (ins uint32, err error) {
	switch {
	case xs > 0x1 || Zm > 0x1f || msz > 0x3 || Pg > 0x7 || Rn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_bz_s_x32_scaled
	case msz == 0x1:
		// encoding_prfh_i_p_bz_s_x32_scaled
	case msz == 0x2:
		// encoding_prfw_i_p_bz_s_x32_scaled
	case msz == 0x3:
		// encoding_prfd_i_p_bz_s_x32_scaled
	default:
		err = errUnmatched
	}

	ins |= (xs << 22) | (Zm << 16) | (msz << 13) | (Pg << 10) | (Rn << 5) | prfop
	ins |= 0x84200000

	return
}

func encode_sve_mem_32b_gld_sv_a(xs, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case xs > 0x1 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_bz_s_x32_scaled
	case U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_bz_s_x32_scaled
	case U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_bz_s_x32_scaled
	case U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_bz_s_x32_scaled
	default:
		err = errUnmatched
	}

	ins |= (xs << 22) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0x84a00000

	return
}

func encode_sve_mem_32b_gld_sv_b(xs, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case xs > 0x1 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case U == 0x0:
		err = errUnallocated
	case U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_bz_s_x32_scaled
	case U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_bz_s_x32_scaled
	default:
		err = errUnmatched
	}

	ins |= (xs << 22) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0x85200000

	return
}

func encode_sve_mem_32b_pfill(imm9h, imm9l, Rn, Pt uint32) (ins uint32, err error) {
	switch {
	case imm9h > 0x3f || imm9l > 0x7 || Rn > 0x1f || Pt > 0xf:
		err = errOverflow
	default:
		// encoding_ldr_p_bi_
	}

	ins |= (imm9h << 16) | (imm9l << 10) | (Rn << 5) | Pt
	ins |= 0x85800000

	return
}

func encode_sve_mem_32b_fill(imm9h, imm9l, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case imm9h > 0x3f || imm9l > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	default:
		// encoding_ldr_z_bi_
	}

	ins |= (imm9h << 16) | (imm9l << 10) | (Rn << 5) | Zt
	ins |= 0x85804000

	return
}

func encode_sve_mem_prfm_si(imm6, msz, Pg, Rn, prfop uint32) (ins uint32, err error) {
	switch {
	case imm6 > 0x3f || msz > 0x3 || Pg > 0x7 || Rn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_bi_s
	case msz == 0x1:
		// encoding_prfh_i_p_bi_s
	case msz == 0x2:
		// encoding_prfw_i_p_bi_s
	case msz == 0x3:
		// encoding_prfd_i_p_bi_s
	default:
		err = errUnmatched
	}

	ins |= (imm6 << 16) | (msz << 13) | (Pg << 10) | (Rn << 5) | prfop
	ins |= 0x85c00000

	return
}

func encode_sve_mem_32b_gld_vs(opc, xs, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || xs > 0x1 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case opc == 0x0 && U == 0x0 && ff == 0x0:
		// encoding_ld1sb_z_p_bz_s_x32_unscaled
	case opc == 0x0 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sb_z_p_bz_s_x32_unscaled
	case opc == 0x0 && U == 0x1 && ff == 0x0:
		// encoding_ld1b_z_p_bz_s_x32_unscaled
	case opc == 0x0 && U == 0x1 && ff == 0x1:
		// encoding_ldff1b_z_p_bz_s_x32_unscaled
	case opc == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_bz_s_x32_unscaled
	case opc == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_bz_s_x32_unscaled
	case opc == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_bz_s_x32_unscaled
	case opc == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_bz_s_x32_unscaled
	case opc == 0x2 && U == 0x0:
		err = errUnallocated
	case opc == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_bz_s_x32_unscaled
	case opc == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_bz_s_x32_unscaled
	default:
		err = errUnmatched
	}

	ins |= (opc << 23) | (xs << 22) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0x84000000

	return
}

func encode_sve_mem_32b_gldnt_vs(msz, Rm, U, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && U == 0x0:
		// encoding_ldnt1sb_z_p_ar_s_x32_unscaled
	case msz == 0x0 && U == 0x1:
		// encoding_ldnt1b_z_p_ar_s_x32_unscaled
	case msz == 0x1 && U == 0x0:
		// encoding_ldnt1sh_z_p_ar_s_x32_unscaled
	case msz == 0x1 && U == 0x1:
		// encoding_ldnt1h_z_p_ar_s_x32_unscaled
	case msz == 0x2 && U == 0x0:
		err = errUnallocated
	case msz == 0x2 && U == 0x1:
		// encoding_ldnt1w_z_p_ar_s_x32_unscaled
	case msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (U << 13) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0x84008000

	return
}

func encode_sve_mem_prfm_ss(msz, Rm, Pg, Rn, prfop uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_br_s
	case msz == 0x1:
		// encoding_prfh_i_p_br_s
	case msz == 0x2:
		// encoding_prfw_i_p_br_s
	case msz == 0x3:
		// encoding_prfd_i_p_br_s
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (Pg << 10) | (Rn << 5) | prfop
	ins |= 0x8400c000

	return
}

func encode_sve_mem_32b_prfm_vi(msz, imm5, Pg, Zn, prfop uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm5 > 0x1f || Pg > 0x7 || Zn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_ai_s
	case msz == 0x1:
		// encoding_prfh_i_p_ai_s
	case msz == 0x2:
		// encoding_prfw_i_p_ai_s
	case msz == 0x3:
		// encoding_prfd_i_p_ai_s
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm5 << 16) | (Pg << 10) | (Zn << 5) | prfop
	ins |= 0x8400e000

	return
}

func encode_sve_mem_32b_gld_vi(msz, imm5, U, ff, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm5 > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && U == 0x0 && ff == 0x0:
		// encoding_ld1sb_z_p_ai_s
	case msz == 0x0 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sb_z_p_ai_s
	case msz == 0x0 && U == 0x1 && ff == 0x0:
		// encoding_ld1b_z_p_ai_s
	case msz == 0x0 && U == 0x1 && ff == 0x1:
		// encoding_ldff1b_z_p_ai_s
	case msz == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_ai_s
	case msz == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_ai_s
	case msz == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_ai_s
	case msz == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_ai_s
	case msz == 0x2 && U == 0x0:
		err = errUnallocated
	case msz == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_ai_s
	case msz == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_ai_s
	case msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm5 << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0x84208000

	return
}

func encode_sve_mem_ld_dup(dtypeh, imm6, dtypel, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case dtypeh > 0x3 || imm6 > 0x3f || dtypel > 0x3 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case dtypeh == 0x0 && dtypel == 0x0:
		// encoding_ld1rb_z_p_bi_u8
	case dtypeh == 0x0 && dtypel == 0x1:
		// encoding_ld1rb_z_p_bi_u16
	case dtypeh == 0x0 && dtypel == 0x2:
		// encoding_ld1rb_z_p_bi_u32
	case dtypeh == 0x0 && dtypel == 0x3:
		// encoding_ld1rb_z_p_bi_u64
	case dtypeh == 0x1 && dtypel == 0x0:
		// encoding_ld1rsw_z_p_bi_s64
	case dtypeh == 0x1 && dtypel == 0x1:
		// encoding_ld1rh_z_p_bi_u16
	case dtypeh == 0x1 && dtypel == 0x2:
		// encoding_ld1rh_z_p_bi_u32
	case dtypeh == 0x1 && dtypel == 0x3:
		// encoding_ld1rh_z_p_bi_u64
	case dtypeh == 0x2 && dtypel == 0x0:
		// encoding_ld1rsh_z_p_bi_s64
	case dtypeh == 0x2 && dtypel == 0x1:
		// encoding_ld1rsh_z_p_bi_s32
	case dtypeh == 0x2 && dtypel == 0x2:
		// encoding_ld1rw_z_p_bi_u32
	case dtypeh == 0x2 && dtypel == 0x3:
		// encoding_ld1rw_z_p_bi_u64
	case dtypeh == 0x3 && dtypel == 0x0:
		// encoding_ld1rsb_z_p_bi_s64
	case dtypeh == 0x3 && dtypel == 0x1:
		// encoding_ld1rsb_z_p_bi_s32
	case dtypeh == 0x3 && dtypel == 0x2:
		// encoding_ld1rsb_z_p_bi_s16
	case dtypeh == 0x3 && dtypel == 0x3:
		// encoding_ld1rd_z_p_bi_u64
	default:
		err = errUnmatched
	}

	ins |= (dtypeh << 23) | (imm6 << 16) | (dtypel << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0x84408000

	return
}

func encode_sve_mem_cldnt_si(msz, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_ldnt1b_z_p_bi_contiguous
	case msz == 0x1:
		// encoding_ldnt1h_z_p_bi_contiguous
	case msz == 0x2:
		// encoding_ldnt1w_z_p_bi_contiguous
	case msz == 0x3:
		// encoding_ldnt1d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa400e000

	return
}

func encode_sve_mem_cldnt_ss(msz, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_ldnt1b_z_p_br_contiguous
	case msz == 0x1:
		// encoding_ldnt1h_z_p_br_contiguous
	case msz == 0x2:
		// encoding_ldnt1w_z_p_br_contiguous
	case msz == 0x3:
		// encoding_ldnt1d_z_p_br_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa400c000

	return
}

func encode_sve_mem_eld_si(msz, opc, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || opc > 0x3 || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && opc == 0x1:
		// encoding_ld2b_z_p_bi_contiguous
	case msz == 0x0 && opc == 0x2:
		// encoding_ld3b_z_p_bi_contiguous
	case msz == 0x0 && opc == 0x3:
		// encoding_ld4b_z_p_bi_contiguous
	case msz == 0x1 && opc == 0x1:
		// encoding_ld2h_z_p_bi_contiguous
	case msz == 0x1 && opc == 0x2:
		// encoding_ld3h_z_p_bi_contiguous
	case msz == 0x1 && opc == 0x3:
		// encoding_ld4h_z_p_bi_contiguous
	case msz == 0x2 && opc == 0x1:
		// encoding_ld2w_z_p_bi_contiguous
	case msz == 0x2 && opc == 0x2:
		// encoding_ld3w_z_p_bi_contiguous
	case msz == 0x2 && opc == 0x3:
		// encoding_ld4w_z_p_bi_contiguous
	case msz == 0x3 && opc == 0x1:
		// encoding_ld2d_z_p_bi_contiguous
	case msz == 0x3 && opc == 0x2:
		// encoding_ld3d_z_p_bi_contiguous
	case msz == 0x3 && opc == 0x3:
		// encoding_ld4d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (opc << 21) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa400e000

	return
}

func encode_sve_mem_eld_ss(msz, opc, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || opc > 0x3 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && opc == 0x1:
		// encoding_ld2b_z_p_br_contiguous
	case msz == 0x0 && opc == 0x2:
		// encoding_ld3b_z_p_br_contiguous
	case msz == 0x0 && opc == 0x3:
		// encoding_ld4b_z_p_br_contiguous
	case msz == 0x1 && opc == 0x1:
		// encoding_ld2h_z_p_br_contiguous
	case msz == 0x1 && opc == 0x2:
		// encoding_ld3h_z_p_br_contiguous
	case msz == 0x1 && opc == 0x3:
		// encoding_ld4h_z_p_br_contiguous
	case msz == 0x2 && opc == 0x1:
		// encoding_ld2w_z_p_br_contiguous
	case msz == 0x2 && opc == 0x2:
		// encoding_ld3w_z_p_br_contiguous
	case msz == 0x2 && opc == 0x3:
		// encoding_ld4w_z_p_br_contiguous
	case msz == 0x3 && opc == 0x1:
		// encoding_ld2d_z_p_br_contiguous
	case msz == 0x3 && opc == 0x2:
		// encoding_ld3d_z_p_br_contiguous
	case msz == 0x3 && opc == 0x3:
		// encoding_ld4d_z_p_br_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (opc << 21) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa400c000

	return
}

func encode_sve_mem_ldqr_si(msz, ssz, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || ssz > 0x3 || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case (ssz & 0x2) == 0x2:
		err = errUnallocated
	case msz == 0x0 && ssz == 0x0:
		// encoding_ld1rqb_z_p_bi_u8
	case msz == 0x0 && ssz == 0x1:
		// encoding_ld1rob_z_p_bi_u8
	case msz == 0x1 && ssz == 0x0:
		// encoding_ld1rqh_z_p_bi_u16
	case msz == 0x1 && ssz == 0x1:
		// encoding_ld1roh_z_p_bi_u16
	case msz == 0x2 && ssz == 0x0:
		// encoding_ld1rqw_z_p_bi_u32
	case msz == 0x2 && ssz == 0x1:
		// encoding_ld1row_z_p_bi_u32
	case msz == 0x3 && ssz == 0x0:
		// encoding_ld1rqd_z_p_bi_u64
	case msz == 0x3 && ssz == 0x1:
		// encoding_ld1rod_z_p_bi_u64
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (ssz << 21) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa4002000

	return
}

func encode_sve_mem_cld_si(dtype, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case dtype > 0xf || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case dtype == 0x0:
		// encoding_ld1b_z_p_bi_u8
	case dtype == 0x1:
		// encoding_ld1b_z_p_bi_u16
	case dtype == 0x2:
		// encoding_ld1b_z_p_bi_u32
	case dtype == 0x3:
		// encoding_ld1b_z_p_bi_u64
	case dtype == 0x4:
		// encoding_ld1sw_z_p_bi_s64
	case dtype == 0x5:
		// encoding_ld1h_z_p_bi_u16
	case dtype == 0x6:
		// encoding_ld1h_z_p_bi_u32
	case dtype == 0x7:
		// encoding_ld1h_z_p_bi_u64
	case dtype == 0x8:
		// encoding_ld1sh_z_p_bi_s64
	case dtype == 0x9:
		// encoding_ld1sh_z_p_bi_s32
	case dtype == 0xa:
		// encoding_ld1w_z_p_bi_u32
	case dtype == 0xb:
		// encoding_ld1w_z_p_bi_u64
	case dtype == 0xc:
		// encoding_ld1sb_z_p_bi_s64
	case dtype == 0xd:
		// encoding_ld1sb_z_p_bi_s32
	case dtype == 0xe:
		// encoding_ld1sb_z_p_bi_s16
	case dtype == 0xf:
		// encoding_ld1d_z_p_bi_u64
	default:
		err = errUnmatched
	}

	ins |= (dtype << 21) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa400a000

	return
}

func encode_sve_mem_cldnf_si(dtype, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case dtype > 0xf || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case dtype == 0x0:
		// encoding_ldnf1b_z_p_bi_u8
	case dtype == 0x1:
		// encoding_ldnf1b_z_p_bi_u16
	case dtype == 0x2:
		// encoding_ldnf1b_z_p_bi_u32
	case dtype == 0x3:
		// encoding_ldnf1b_z_p_bi_u64
	case dtype == 0x4:
		// encoding_ldnf1sw_z_p_bi_s64
	case dtype == 0x5:
		// encoding_ldnf1h_z_p_bi_u16
	case dtype == 0x6:
		// encoding_ldnf1h_z_p_bi_u32
	case dtype == 0x7:
		// encoding_ldnf1h_z_p_bi_u64
	case dtype == 0x8:
		// encoding_ldnf1sh_z_p_bi_s64
	case dtype == 0x9:
		// encoding_ldnf1sh_z_p_bi_s32
	case dtype == 0xa:
		// encoding_ldnf1w_z_p_bi_u32
	case dtype == 0xb:
		// encoding_ldnf1w_z_p_bi_u64
	case dtype == 0xc:
		// encoding_ldnf1sb_z_p_bi_s64
	case dtype == 0xd:
		// encoding_ldnf1sb_z_p_bi_s32
	case dtype == 0xe:
		// encoding_ldnf1sb_z_p_bi_s16
	case dtype == 0xf:
		// encoding_ldnf1d_z_p_bi_u64
	default:
		err = errUnmatched
	}

	ins |= (dtype << 21) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa410a000

	return
}

func encode_sve_mem_ldqr_ss(msz, ssz, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || ssz > 0x3 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case (ssz & 0x2) == 0x2:
		err = errUnallocated
	case msz == 0x0 && ssz == 0x0:
		// encoding_ld1rqb_z_p_br_contiguous
	case msz == 0x0 && ssz == 0x1:
		// encoding_ld1rob_z_p_br_contiguous
	case msz == 0x1 && ssz == 0x0:
		// encoding_ld1rqh_z_p_br_contiguous
	case msz == 0x1 && ssz == 0x1:
		// encoding_ld1roh_z_p_br_contiguous
	case msz == 0x2 && ssz == 0x0:
		// encoding_ld1rqw_z_p_br_contiguous
	case msz == 0x2 && ssz == 0x1:
		// encoding_ld1row_z_p_br_contiguous
	case msz == 0x3 && ssz == 0x0:
		// encoding_ld1rqd_z_p_br_contiguous
	case msz == 0x3 && ssz == 0x1:
		// encoding_ld1rod_z_p_br_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (ssz << 21) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa4000000

	return
}

func encode_sve_mem_cld_ss(dtype, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case dtype > 0xf || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case dtype == 0x0:
		// encoding_ld1b_z_p_br_u8
	case dtype == 0x1:
		// encoding_ld1b_z_p_br_u16
	case dtype == 0x2:
		// encoding_ld1b_z_p_br_u32
	case dtype == 0x3:
		// encoding_ld1b_z_p_br_u64
	case dtype == 0x4:
		// encoding_ld1sw_z_p_br_s64
	case dtype == 0x5:
		// encoding_ld1h_z_p_br_u16
	case dtype == 0x6:
		// encoding_ld1h_z_p_br_u32
	case dtype == 0x7:
		// encoding_ld1h_z_p_br_u64
	case dtype == 0x8:
		// encoding_ld1sh_z_p_br_s64
	case dtype == 0x9:
		// encoding_ld1sh_z_p_br_s32
	case dtype == 0xa:
		// encoding_ld1w_z_p_br_u32
	case dtype == 0xb:
		// encoding_ld1w_z_p_br_u64
	case dtype == 0xc:
		// encoding_ld1sb_z_p_br_s64
	case dtype == 0xd:
		// encoding_ld1sb_z_p_br_s32
	case dtype == 0xe:
		// encoding_ld1sb_z_p_br_s16
	case dtype == 0xf:
		// encoding_ld1d_z_p_br_u64
	default:
		err = errUnmatched
	}

	ins |= (dtype << 21) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa4004000

	return
}

func encode_sve_mem_cldff_ss(dtype, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case dtype > 0xf || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case dtype == 0x0:
		// encoding_ldff1b_z_p_br_u8
	case dtype == 0x1:
		// encoding_ldff1b_z_p_br_u16
	case dtype == 0x2:
		// encoding_ldff1b_z_p_br_u32
	case dtype == 0x3:
		// encoding_ldff1b_z_p_br_u64
	case dtype == 0x4:
		// encoding_ldff1sw_z_p_br_s64
	case dtype == 0x5:
		// encoding_ldff1h_z_p_br_u16
	case dtype == 0x6:
		// encoding_ldff1h_z_p_br_u32
	case dtype == 0x7:
		// encoding_ldff1h_z_p_br_u64
	case dtype == 0x8:
		// encoding_ldff1sh_z_p_br_s64
	case dtype == 0x9:
		// encoding_ldff1sh_z_p_br_s32
	case dtype == 0xa:
		// encoding_ldff1w_z_p_br_u32
	case dtype == 0xb:
		// encoding_ldff1w_z_p_br_u64
	case dtype == 0xc:
		// encoding_ldff1sb_z_p_br_s64
	case dtype == 0xd:
		// encoding_ldff1sb_z_p_br_s32
	case dtype == 0xe:
		// encoding_ldff1sb_z_p_br_s16
	case dtype == 0xf:
		// encoding_ldff1d_z_p_br_u64
	default:
		err = errUnmatched
	}

	ins |= (dtype << 21) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xa4006000

	return
}

func encode_sve_mem_64b_prfm_sv2(Zm, msz, Pg, Rn, prfop uint32) (ins uint32, err error) {
	switch {
	case Zm > 0x1f || msz > 0x3 || Pg > 0x7 || Rn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_bz_d_64_scaled
	case msz == 0x1:
		// encoding_prfh_i_p_bz_d_64_scaled
	case msz == 0x2:
		// encoding_prfw_i_p_bz_d_64_scaled
	case msz == 0x3:
		// encoding_prfd_i_p_bz_d_64_scaled
	default:
		err = errUnmatched
	}

	ins |= (Zm << 16) | (msz << 13) | (Pg << 10) | (Rn << 5) | prfop
	ins |= 0xc4608000

	return
}

func encode_sve_mem_64b_prfm_sv(xs, Zm, msz, Pg, Rn, prfop uint32) (ins uint32, err error) {
	switch {
	case xs > 0x1 || Zm > 0x1f || msz > 0x3 || Pg > 0x7 || Rn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_bz_d_x32_scaled
	case msz == 0x1:
		// encoding_prfh_i_p_bz_d_x32_scaled
	case msz == 0x2:
		// encoding_prfw_i_p_bz_d_x32_scaled
	case msz == 0x3:
		// encoding_prfd_i_p_bz_d_x32_scaled
	default:
		err = errUnmatched
	}

	ins |= (xs << 22) | (Zm << 16) | (msz << 13) | (Pg << 10) | (Rn << 5) | prfop
	ins |= 0xc4200000

	return
}

func encode_sve_mem_64b_gld_sv2(opc, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case opc == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_bz_d_64_scaled
	case opc == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_bz_d_64_scaled
	case opc == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_bz_d_64_scaled
	case opc == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_bz_d_64_scaled
	case opc == 0x2 && U == 0x0 && ff == 0x0:
		// encoding_ld1sw_z_p_bz_d_64_scaled
	case opc == 0x2 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sw_z_p_bz_d_64_scaled
	case opc == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_bz_d_64_scaled
	case opc == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_bz_d_64_scaled
	case opc == 0x3 && U == 0x0:
		err = errUnallocated
	case opc == 0x3 && U == 0x1 && ff == 0x0:
		// encoding_ld1d_z_p_bz_d_64_scaled
	case opc == 0x3 && U == 0x1 && ff == 0x1:
		// encoding_ldff1d_z_p_bz_d_64_scaled
	default:
		err = errUnmatched
	}

	ins |= (opc << 23) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xc4608000

	return
}

func encode_sve_mem_64b_gld_sv(opc, xs, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || xs > 0x1 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case opc == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_bz_d_x32_scaled
	case opc == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_bz_d_x32_scaled
	case opc == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_bz_d_x32_scaled
	case opc == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_bz_d_x32_scaled
	case opc == 0x2 && U == 0x0 && ff == 0x0:
		// encoding_ld1sw_z_p_bz_d_x32_scaled
	case opc == 0x2 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sw_z_p_bz_d_x32_scaled
	case opc == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_bz_d_x32_scaled
	case opc == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_bz_d_x32_scaled
	case opc == 0x3 && U == 0x0:
		err = errUnallocated
	case opc == 0x3 && U == 0x1 && ff == 0x0:
		// encoding_ld1d_z_p_bz_d_x32_scaled
	case opc == 0x3 && U == 0x1 && ff == 0x1:
		// encoding_ldff1d_z_p_bz_d_x32_scaled
	default:
		err = errUnmatched
	}

	ins |= (opc << 23) | (xs << 22) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xc4200000

	return
}

func encode_sve_mem_64b_prfm_vi(msz, imm5, Pg, Zn, prfop uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm5 > 0x1f || Pg > 0x7 || Zn > 0x1f || prfop > 0xf:
		err = errOverflow
	case msz == 0x0:
		// encoding_prfb_i_p_ai_d
	case msz == 0x1:
		// encoding_prfh_i_p_ai_d
	case msz == 0x2:
		// encoding_prfw_i_p_ai_d
	case msz == 0x3:
		// encoding_prfd_i_p_ai_d
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm5 << 16) | (Pg << 10) | (Zn << 5) | prfop
	ins |= 0xc400e000

	return
}

func encode_sve_mem_64b_gldnt_vs(msz, Rm, U, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || U > 0x1 || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && U == 0x0:
		// encoding_ldnt1sb_z_p_ar_d_64_unscaled
	case msz == 0x0 && U == 0x1:
		// encoding_ldnt1b_z_p_ar_d_64_unscaled
	case msz == 0x1 && U == 0x0:
		// encoding_ldnt1sh_z_p_ar_d_64_unscaled
	case msz == 0x1 && U == 0x1:
		// encoding_ldnt1h_z_p_ar_d_64_unscaled
	case msz == 0x2 && U == 0x0:
		// encoding_ldnt1sw_z_p_ar_d_64_unscaled
	case msz == 0x2 && U == 0x1:
		// encoding_ldnt1w_z_p_ar_d_64_unscaled
	case msz == 0x3 && U == 0x0:
		err = errUnallocated
	case msz == 0x3 && U == 0x1:
		// encoding_ldnt1d_z_p_ar_d_64_unscaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (U << 14) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0xc4008000

	return
}

func encode_sve_mem_64b_gld_vi(msz, imm5, U, ff, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm5 > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && U == 0x0 && ff == 0x0:
		// encoding_ld1sb_z_p_ai_d
	case msz == 0x0 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sb_z_p_ai_d
	case msz == 0x0 && U == 0x1 && ff == 0x0:
		// encoding_ld1b_z_p_ai_d
	case msz == 0x0 && U == 0x1 && ff == 0x1:
		// encoding_ldff1b_z_p_ai_d
	case msz == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_ai_d
	case msz == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_ai_d
	case msz == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_ai_d
	case msz == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_ai_d
	case msz == 0x2 && U == 0x0 && ff == 0x0:
		// encoding_ld1sw_z_p_ai_d
	case msz == 0x2 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sw_z_p_ai_d
	case msz == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_ai_d
	case msz == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_ai_d
	case msz == 0x3 && U == 0x0:
		err = errUnallocated
	case msz == 0x3 && U == 0x1 && ff == 0x0:
		// encoding_ld1d_z_p_ai_d
	case msz == 0x3 && U == 0x1 && ff == 0x1:
		// encoding_ldff1d_z_p_ai_d
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm5 << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0xc4208000

	return
}

func encode_sve_mem_64b_gld_vs2(msz, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && U == 0x0 && ff == 0x0:
		// encoding_ld1sb_z_p_bz_d_64_unscaled
	case msz == 0x0 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sb_z_p_bz_d_64_unscaled
	case msz == 0x0 && U == 0x1 && ff == 0x0:
		// encoding_ld1b_z_p_bz_d_64_unscaled
	case msz == 0x0 && U == 0x1 && ff == 0x1:
		// encoding_ldff1b_z_p_bz_d_64_unscaled
	case msz == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_bz_d_64_unscaled
	case msz == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_bz_d_64_unscaled
	case msz == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_bz_d_64_unscaled
	case msz == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_bz_d_64_unscaled
	case msz == 0x2 && U == 0x0 && ff == 0x0:
		// encoding_ld1sw_z_p_bz_d_64_unscaled
	case msz == 0x2 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sw_z_p_bz_d_64_unscaled
	case msz == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_bz_d_64_unscaled
	case msz == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_bz_d_64_unscaled
	case msz == 0x3 && U == 0x0:
		err = errUnallocated
	case msz == 0x3 && U == 0x1 && ff == 0x0:
		// encoding_ld1d_z_p_bz_d_64_unscaled
	case msz == 0x3 && U == 0x1 && ff == 0x1:
		// encoding_ldff1d_z_p_bz_d_64_unscaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xc4408000

	return
}

func encode_sve_mem_64b_gld_vs(msz, xs, Zm, U, ff, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || xs > 0x1 || Zm > 0x1f || U > 0x1 || ff > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && U == 0x0 && ff == 0x0:
		// encoding_ld1sb_z_p_bz_d_x32_unscaled
	case msz == 0x0 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sb_z_p_bz_d_x32_unscaled
	case msz == 0x0 && U == 0x1 && ff == 0x0:
		// encoding_ld1b_z_p_bz_d_x32_unscaled
	case msz == 0x0 && U == 0x1 && ff == 0x1:
		// encoding_ldff1b_z_p_bz_d_x32_unscaled
	case msz == 0x1 && U == 0x0 && ff == 0x0:
		// encoding_ld1sh_z_p_bz_d_x32_unscaled
	case msz == 0x1 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sh_z_p_bz_d_x32_unscaled
	case msz == 0x1 && U == 0x1 && ff == 0x0:
		// encoding_ld1h_z_p_bz_d_x32_unscaled
	case msz == 0x1 && U == 0x1 && ff == 0x1:
		// encoding_ldff1h_z_p_bz_d_x32_unscaled
	case msz == 0x2 && U == 0x0 && ff == 0x0:
		// encoding_ld1sw_z_p_bz_d_x32_unscaled
	case msz == 0x2 && U == 0x0 && ff == 0x1:
		// encoding_ldff1sw_z_p_bz_d_x32_unscaled
	case msz == 0x2 && U == 0x1 && ff == 0x0:
		// encoding_ld1w_z_p_bz_d_x32_unscaled
	case msz == 0x2 && U == 0x1 && ff == 0x1:
		// encoding_ldff1w_z_p_bz_d_x32_unscaled
	case msz == 0x3 && U == 0x0:
		err = errUnallocated
	case msz == 0x3 && U == 0x1 && ff == 0x0:
		// encoding_ld1d_z_p_bz_d_x32_unscaled
	case msz == 0x3 && U == 0x1 && ff == 0x1:
		// encoding_ldff1d_z_p_bz_d_x32_unscaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (xs << 22) | (Zm << 16) | (U << 14) | (ff << 13) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xc4000000

	return
}

func encode_sve_mem_pspill(imm9h, imm9l, Rn, Pt uint32) (ins uint32, err error) {
	switch {
	case imm9h > 0x3f || imm9l > 0x7 || Rn > 0x1f || Pt > 0xf:
		err = errOverflow
	default:
		// encoding_str_p_bi_
	}

	ins |= (imm9h << 16) | (imm9l << 10) | (Rn << 5) | Pt
	ins |= 0xe5800000

	return
}

func encode_sve_mem_spill(imm9h, imm9l, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case imm9h > 0x3f || imm9l > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	default:
		// encoding_str_z_bi_
	}

	ins |= (imm9h << 16) | (imm9l << 10) | (Rn << 5) | Zt
	ins |= 0xe5804000

	return
}

func encode_sve_mem_cst_ss(opc, o2, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x7 || o2 > 0x1 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case (opc & 0x6) == 0x0:
		// encoding_st1b_z_p_br_
	case (opc & 0x6) == 0x2:
		// encoding_st1h_z_p_br_
	case (opc & 0x6) == 0x4:
		// encoding_st1w_z_p_br_
	case opc == 0x7 && o2 == 0x0:
		err = errUnallocated
	case opc == 0x7 && o2 == 0x1:
		// encoding_st1d_z_p_br_
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (o2 << 21) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4004000

	return
}

func encode_sve_mem_sstnt_64b_vs(msz, Rm, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_stnt1b_z_p_ar_d_64_unscaled
	case msz == 0x1:
		// encoding_stnt1h_z_p_ar_d_64_unscaled
	case msz == 0x2:
		// encoding_stnt1w_z_p_ar_d_64_unscaled
	case msz == 0x3:
		// encoding_stnt1d_z_p_ar_d_64_unscaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0xe4002000

	return
}

func encode_sve_mem_cstnt_ss(msz, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_stnt1b_z_p_br_contiguous
	case msz == 0x1:
		// encoding_stnt1h_z_p_br_contiguous
	case msz == 0x2:
		// encoding_stnt1w_z_p_br_contiguous
	case msz == 0x3:
		// encoding_stnt1d_z_p_br_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4006000

	return
}

func encode_sve_mem_sstnt_32b_vs(msz, Rm, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Rm > 0x1f || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_stnt1b_z_p_ar_s_x32_unscaled
	case msz == 0x1:
		// encoding_stnt1h_z_p_ar_s_x32_unscaled
	case msz == 0x2:
		// encoding_stnt1w_z_p_ar_s_x32_unscaled
	case msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Rm << 16) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0xe4402000

	return
}

func encode_sve_mem_est_ss(msz, opc, Rm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || opc > 0x3 || Rm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && opc == 0x1:
		// encoding_st2b_z_p_br_contiguous
	case msz == 0x0 && opc == 0x2:
		// encoding_st3b_z_p_br_contiguous
	case msz == 0x0 && opc == 0x3:
		// encoding_st4b_z_p_br_contiguous
	case msz == 0x1 && opc == 0x1:
		// encoding_st2h_z_p_br_contiguous
	case msz == 0x1 && opc == 0x2:
		// encoding_st3h_z_p_br_contiguous
	case msz == 0x1 && opc == 0x3:
		// encoding_st4h_z_p_br_contiguous
	case msz == 0x2 && opc == 0x1:
		// encoding_st2w_z_p_br_contiguous
	case msz == 0x2 && opc == 0x2:
		// encoding_st3w_z_p_br_contiguous
	case msz == 0x2 && opc == 0x3:
		// encoding_st4w_z_p_br_contiguous
	case msz == 0x3 && opc == 0x1:
		// encoding_st2d_z_p_br_contiguous
	case msz == 0x3 && opc == 0x2:
		// encoding_st3d_z_p_br_contiguous
	case msz == 0x3 && opc == 0x3:
		// encoding_st4d_z_p_br_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (opc << 21) | (Rm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4006000

	return
}

func encode_sve_mem_sst_vs_a(msz, Zm, xs, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || xs > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_z_p_bz_d_x32_unscaled
	case msz == 0x1:
		// encoding_st1h_z_p_bz_d_x32_unscaled
	case msz == 0x2:
		// encoding_st1w_z_p_bz_d_x32_unscaled
	case msz == 0x3:
		// encoding_st1d_z_p_bz_d_x32_unscaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (xs << 14) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4008000

	return
}

func encode_sve_mem_sst_sv_a(msz, Zm, xs, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || xs > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		err = errUnallocated
	case msz == 0x1:
		// encoding_st1h_z_p_bz_d_x32_scaled
	case msz == 0x2:
		// encoding_st1w_z_p_bz_d_x32_scaled
	case msz == 0x3:
		// encoding_st1d_z_p_bz_d_x32_scaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (xs << 14) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4208000

	return
}

func encode_sve_mem_sst_vs_b(msz, Zm, xs, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || xs > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_z_p_bz_s_x32_unscaled
	case msz == 0x1:
		// encoding_st1h_z_p_bz_s_x32_unscaled
	case msz == 0x2:
		// encoding_st1w_z_p_bz_s_x32_unscaled
	case msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (xs << 14) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4408000

	return
}

func encode_sve_mem_sst_sv_b(msz, Zm, xs, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || xs > 0x1 || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		err = errUnallocated
	case msz == 0x1:
		// encoding_st1h_z_p_bz_s_x32_scaled
	case msz == 0x2:
		// encoding_st1w_z_p_bz_s_x32_scaled
	case msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (xs << 14) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe4608000

	return
}

func encode_sve_mem_sst_vs2(msz, Zm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_z_p_bz_d_64_unscaled
	case msz == 0x1:
		// encoding_st1h_z_p_bz_d_64_unscaled
	case msz == 0x2:
		// encoding_st1w_z_p_bz_d_64_unscaled
	case msz == 0x3:
		// encoding_st1d_z_p_bz_d_64_unscaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe400a000

	return
}

func encode_sve_mem_sst_sv2(msz, Zm, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || Zm > 0x1f || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		err = errUnallocated
	case msz == 0x1:
		// encoding_st1h_z_p_bz_d_64_scaled
	case msz == 0x2:
		// encoding_st1w_z_p_bz_d_64_scaled
	case msz == 0x3:
		// encoding_st1d_z_p_bz_d_64_scaled
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (Zm << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe420a000

	return
}

func encode_sve_mem_sst_vi_a(msz, imm5, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm5 > 0x1f || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_z_p_ai_d
	case msz == 0x1:
		// encoding_st1h_z_p_ai_d
	case msz == 0x2:
		// encoding_st1w_z_p_ai_d
	case msz == 0x3:
		// encoding_st1d_z_p_ai_d
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm5 << 16) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0xe440a000

	return
}

func encode_sve_mem_sst_vi_b(msz, imm5, Pg, Zn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm5 > 0x1f || Pg > 0x7 || Zn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_z_p_ai_s
	case msz == 0x1:
		// encoding_st1h_z_p_ai_s
	case msz == 0x2:
		// encoding_st1w_z_p_ai_s
	case msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm5 << 16) | (Pg << 10) | (Zn << 5) | Zt
	ins |= 0xe460a000

	return
}

func encode_sve_mem_cstnt_si(msz, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_stnt1b_z_p_bi_contiguous
	case msz == 0x1:
		// encoding_stnt1h_z_p_bi_contiguous
	case msz == 0x2:
		// encoding_stnt1w_z_p_bi_contiguous
	case msz == 0x3:
		// encoding_stnt1d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe410e000

	return
}

func encode_sve_mem_est_si(msz, opc, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || opc > 0x3 || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0 && opc == 0x1:
		// encoding_st2b_z_p_bi_contiguous
	case msz == 0x0 && opc == 0x2:
		// encoding_st3b_z_p_bi_contiguous
	case msz == 0x0 && opc == 0x3:
		// encoding_st4b_z_p_bi_contiguous
	case msz == 0x1 && opc == 0x1:
		// encoding_st2h_z_p_bi_contiguous
	case msz == 0x1 && opc == 0x2:
		// encoding_st3h_z_p_bi_contiguous
	case msz == 0x1 && opc == 0x3:
		// encoding_st4h_z_p_bi_contiguous
	case msz == 0x2 && opc == 0x1:
		// encoding_st2w_z_p_bi_contiguous
	case msz == 0x2 && opc == 0x2:
		// encoding_st3w_z_p_bi_contiguous
	case msz == 0x2 && opc == 0x3:
		// encoding_st4w_z_p_bi_contiguous
	case msz == 0x3 && opc == 0x1:
		// encoding_st2d_z_p_bi_contiguous
	case msz == 0x3 && opc == 0x2:
		// encoding_st3d_z_p_bi_contiguous
	case msz == 0x3 && opc == 0x3:
		// encoding_st4d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (opc << 21) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe410e000

	return
}

func encode_sve_mem_cst_si(msz, size, imm4, Pg, Rn, Zt uint32) (ins uint32, err error) {
	switch {
	case msz > 0x3 || size > 0x3 || imm4 > 0xf || Pg > 0x7 || Rn > 0x1f || Zt > 0x1f:
		err = errOverflow
	case msz == 0x0:
		// encoding_st1b_z_p_bi_
	case msz == 0x1:
		// encoding_st1h_z_p_bi_
	case msz == 0x2:
		// encoding_st1w_z_p_bi_
	case msz == 0x3:
		// encoding_st1d_z_p_bi_
	default:
		err = errUnmatched
	}

	ins |= (msz << 23) | (size << 21) | (imm4 << 16) | (Pg << 10) | (Rn << 5) | Zt
	ins |= 0xe400e000

	return
}

func encode_pcreladdr(op, immlo, immhi, Rd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || immlo > 0x3 || immhi > 0x7ffff || Rd > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_ADR_only_pcreladdr
	case op == 0x1:
		// encoding_ADRP_only_pcreladdr
	default:
		err = errUnmatched
	}

	ins |= (op << 31) | (immlo << 29) | (immhi << 5) | Rd
	ins |= 0x10000000

	return
}

func encode_addsub_imm(sf, op, S, sh, imm12, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || sh > 0x1 || imm12 > 0xfff || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case sf == 0x0 && op == 0x0 && S == 0x0:
		// encoding_ADD_32_addsub_imm
	case sf == 0x0 && op == 0x0 && S == 0x1:
		// encoding_ADDS_32S_addsub_imm
	case sf == 0x0 && op == 0x1 && S == 0x0:
		// encoding_SUB_32_addsub_imm
	case sf == 0x0 && op == 0x1 && S == 0x1:
		// encoding_SUBS_32S_addsub_imm
	case sf == 0x1 && op == 0x0 && S == 0x0:
		// encoding_ADD_64_addsub_imm
	case sf == 0x1 && op == 0x0 && S == 0x1:
		// encoding_ADDS_64S_addsub_imm
	case sf == 0x1 && op == 0x1 && S == 0x0:
		// encoding_SUB_64_addsub_imm
	case sf == 0x1 && op == 0x1 && S == 0x1:
		// encoding_SUBS_64S_addsub_imm
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (sh << 22) | (imm12 << 10) | (Rn << 5) | Rd
	ins |= 0x11000000

	return
}

func encode_addsub_immtags(sf, op, S, o2, uimm6, op3, uimm4, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || o2 > 0x1 || uimm6 > 0x3f || op3 > 0x3 || uimm4 > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case o2 == 0x1:
		err = errUnallocated
	case sf == 0x0 && o2 == 0x0:
		err = errUnallocated
	case sf == 0x1 && S == 0x1 && o2 == 0x0:
		err = errUnallocated
	case sf == 0x1 && op == 0x0 && S == 0x0 && o2 == 0x0:
		// encoding_ADDG_64_addsub_immtags
	case sf == 0x1 && op == 0x1 && S == 0x0 && o2 == 0x0:
		// encoding_SUBG_64_addsub_immtags
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (o2 << 22) | (uimm6 << 16) | (op3 << 14) | (uimm4 << 10) | (Rn << 5) | Rd
	ins |= 0x11800000

	return
}

func encode_log_imm(sf, opc, N, immr, imms, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || opc > 0x3 || N > 0x1 || immr > 0x3f || imms > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case sf == 0x0 && N == 0x1:
		err = errUnallocated
	case sf == 0x0 && opc == 0x0 && N == 0x0:
		// encoding_AND_32_log_imm
	case sf == 0x0 && opc == 0x1 && N == 0x0:
		// encoding_ORR_32_log_imm
	case sf == 0x0 && opc == 0x2 && N == 0x0:
		// encoding_EOR_32_log_imm
	case sf == 0x0 && opc == 0x3 && N == 0x0:
		// encoding_ANDS_32S_log_imm
	case sf == 0x1 && opc == 0x0:
		// encoding_AND_64_log_imm
	case sf == 0x1 && opc == 0x1:
		// encoding_ORR_64_log_imm
	case sf == 0x1 && opc == 0x2:
		// encoding_EOR_64_log_imm
	case sf == 0x1 && opc == 0x3:
		// encoding_ANDS_64S_log_imm
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (opc << 29) | (N << 22) | (immr << 16) | (imms << 10) | (Rn << 5) | Rd
	ins |= 0x12000000

	return
}

func encode_movewide(sf, opc, hw, imm16, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || opc > 0x3 || hw > 0x3 || imm16 > 0xffff || Rd > 0x1f:
		err = errOverflow
	case opc == 0x1:
		err = errUnallocated
	case sf == 0x0 && (hw&0x2) == 0x2:
		err = errUnallocated
	case sf == 0x0 && opc == 0x0 && (hw&0x2) == 0x0:
		// encoding_MOVN_32_movewide
	case sf == 0x0 && opc == 0x2 && (hw&0x2) == 0x0:
		// encoding_MOVZ_32_movewide
	case sf == 0x0 && opc == 0x3 && (hw&0x2) == 0x0:
		// encoding_MOVK_32_movewide
	case sf == 0x1 && opc == 0x0:
		// encoding_MOVN_64_movewide
	case sf == 0x1 && opc == 0x2:
		// encoding_MOVZ_64_movewide
	case sf == 0x1 && opc == 0x3:
		// encoding_MOVK_64_movewide
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (opc << 29) | (hw << 21) | (imm16 << 5) | Rd
	ins |= 0x12800000

	return
}

func encode_bitfield(sf, opc, N, immr, imms, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || opc > 0x3 || N > 0x1 || immr > 0x3f || imms > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opc == 0x3:
		err = errUnallocated
	case sf == 0x0 && N == 0x1:
		err = errUnallocated
	case sf == 0x0 && opc == 0x0 && N == 0x0:
		// encoding_SBFM_32M_bitfield
	case sf == 0x0 && opc == 0x1 && N == 0x0:
		// encoding_BFM_32M_bitfield
	case sf == 0x0 && opc == 0x2 && N == 0x0:
		// encoding_UBFM_32M_bitfield
	case sf == 0x1 && N == 0x0:
		err = errUnallocated
	case sf == 0x1 && opc == 0x0 && N == 0x1:
		// encoding_SBFM_64M_bitfield
	case sf == 0x1 && opc == 0x1 && N == 0x1:
		// encoding_BFM_64M_bitfield
	case sf == 0x1 && opc == 0x2 && N == 0x1:
		// encoding_UBFM_64M_bitfield
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (opc << 29) | (N << 22) | (immr << 16) | (imms << 10) | (Rn << 5) | Rd
	ins |= 0x13000000

	return
}

func encode_extract(sf, op21, N, o0, Rm, imms, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op21 > 0x3 || N > 0x1 || o0 > 0x1 || Rm > 0x1f || imms > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (op21 & 0x1) == 0x1:
		err = errUnallocated
	case op21 == 0x0 && o0 == 0x1:
		err = errUnallocated
	case (op21 & 0x2) == 0x2:
		err = errUnallocated
	case sf == 0x0 && (imms&0x20) == 0x20:
		err = errUnallocated
	case sf == 0x0 && N == 0x1:
		err = errUnallocated
	case sf == 0x0 && op21 == 0x0 && N == 0x0 && o0 == 0x0 && (imms&0x20) == 0x0:
		// encoding_EXTR_32_extract
	case sf == 0x1 && N == 0x0:
		err = errUnallocated
	case sf == 0x1 && op21 == 0x0 && N == 0x1 && o0 == 0x0:
		// encoding_EXTR_64_extract
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op21 << 29) | (N << 22) | (o0 << 21) | (Rm << 16) | (imms << 10) | (Rn << 5) | Rd
	ins |= 0x13800000

	return
}

func encode_condbranch(o1, imm19, o0, cond uint32) (ins uint32, err error) {
	switch {
	case o1 > 0x1 || imm19 > 0x7ffff || o0 > 0x1 || cond > 0xf:
		err = errOverflow
	case o1 == 0x0 && o0 == 0x0:
		// encoding_B_only_condbranch
	case o1 == 0x0 && o0 == 0x1:
		// encoding_BC_only_condbranch
	case o1 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (o1 << 24) | (imm19 << 5) | (o0 << 4) | cond
	ins |= 0x54000000

	return
}

func encode_exception(opc, imm16, op2, LL uint32) (ins uint32, err error) {
	switch {
	case opc > 0x7 || imm16 > 0xffff || op2 > 0x7 || LL > 0x3:
		err = errOverflow
	case op2 == 0x1:
		err = errUnallocated
	case (op2 & 0x6) == 0x2:
		err = errUnallocated
	case (op2 & 0x4) == 0x4:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x0 && LL == 0x0:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x0 && LL == 0x1:
		// encoding_SVC_EX_exception
	case opc == 0x0 && op2 == 0x0 && LL == 0x2:
		// encoding_HVC_EX_exception
	case opc == 0x0 && op2 == 0x0 && LL == 0x3:
		// encoding_SMC_EX_exception
	case opc == 0x1 && op2 == 0x0 && (LL&0x1) == 0x1:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x0 && LL == 0x0:
		// encoding_BRK_EX_exception
	case opc == 0x1 && op2 == 0x0 && (LL&0x2) == 0x2:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x0 && (LL&0x1) == 0x1:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x0 && LL == 0x0:
		// encoding_HLT_EX_exception
	case opc == 0x2 && op2 == 0x0 && (LL&0x2) == 0x2:
		err = errUnallocated
	case opc == 0x3 && op2 == 0x0 && LL == 0x0:
		// encoding_TCANCEL_EX_exception
	case opc == 0x3 && op2 == 0x0 && LL == 0x1:
		err = errUnallocated
	case opc == 0x3 && op2 == 0x0 && (LL&0x2) == 0x2:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x0:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x0 && LL == 0x0:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x0 && LL == 0x1:
		// encoding_DCPS1_DC_exception
	case opc == 0x5 && op2 == 0x0 && LL == 0x2:
		// encoding_DCPS2_DC_exception
	case opc == 0x5 && op2 == 0x0 && LL == 0x3:
		// encoding_DCPS3_DC_exception
	case opc == 0x6 && op2 == 0x0:
		err = errUnallocated
	case opc == 0x7 && op2 == 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 21) | (imm16 << 5) | (op2 << 2) | LL
	ins |= 0xd4000000

	return
}

func encode_systeminstrswithreg(CRm, op2, Rt uint32) (ins uint32, err error) {
	switch {
	case CRm > 0xf || op2 > 0x7 || Rt > 0x1f:
		err = errOverflow
	case CRm != 0x0:
		err = errUnallocated
	case CRm == 0x0 && op2 == 0x0:
		// encoding_WFET_only_systeminstrswithreg
	case CRm == 0x0 && op2 == 0x1:
		// encoding_WFIT_only_systeminstrswithreg
	case CRm == 0x0 && (op2&0x6) == 0x2:
		err = errUnallocated
	case CRm == 0x0 && (op2&0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (CRm << 8) | (op2 << 5) | Rt
	ins |= 0xd5031000

	return
}

func encode_hints(CRm, op2 uint32) (ins uint32, err error) {
	switch {
	case CRm > 0xf || op2 > 0x7:
		err = errOverflow
	case CRm == 0x0 && op2 == 0x0:
		// encoding_NOP_HI_hints
	case CRm == 0x0 && op2 == 0x1:
		// encoding_YIELD_HI_hints
	case CRm == 0x0 && op2 == 0x2:
		// encoding_WFE_HI_hints
	case CRm == 0x0 && op2 == 0x3:
		// encoding_WFI_HI_hints
	case CRm == 0x0 && op2 == 0x4:
		// encoding_SEV_HI_hints
	case CRm == 0x0 && op2 == 0x5:
		// encoding_SEVL_HI_hints
	case CRm == 0x0 && op2 == 0x6:
		// encoding_DGH_HI_hints
	case CRm == 0x0 && op2 == 0x7:
		// encoding_XPACLRI_HI_hints
	case CRm == 0x1 && op2 == 0x0:
		// encoding_PACIA1716_HI_hints
	case CRm == 0x1 && op2 == 0x2:
		// encoding_PACIB1716_HI_hints
	case CRm == 0x1 && op2 == 0x4:
		// encoding_AUTIA1716_HI_hints
	case CRm == 0x1 && op2 == 0x6:
		// encoding_AUTIB1716_HI_hints
	case CRm == 0x2 && op2 == 0x0:
		// encoding_ESB_HI_hints
	case CRm == 0x2 && op2 == 0x1:
		// encoding_PSB_HC_hints
	case CRm == 0x2 && op2 == 0x2:
		// encoding_TSB_HC_hints
	case CRm == 0x2 && op2 == 0x4:
		// encoding_CSDB_HI_hints
	case CRm == 0x3 && op2 == 0x0:
		// encoding_PACIAZ_HI_hints
	case CRm == 0x3 && op2 == 0x1:
		// encoding_PACIASP_HI_hints
	case CRm == 0x3 && op2 == 0x2:
		// encoding_PACIBZ_HI_hints
	case CRm == 0x3 && op2 == 0x3:
		// encoding_PACIBSP_HI_hints
	case CRm == 0x3 && op2 == 0x4:
		// encoding_AUTIAZ_HI_hints
	case CRm == 0x3 && op2 == 0x5:
		// encoding_AUTIASP_HI_hints
	case CRm == 0x3 && op2 == 0x6:
		// encoding_AUTIBZ_HI_hints
	case CRm == 0x3 && op2 == 0x7:
		// encoding_AUTIBSP_HI_hints
	case CRm == 0x4 && (op2&0x1) == 0x0:
		// encoding_BTI_HB_hints
	default:
		// encoding_HINT_HM_hints
	}

	ins |= (CRm << 8) | (op2 << 5)
	ins |= 0xd503201f

	return
}

func encode_barriers(CRm, op2, Rt uint32) (ins uint32, err error) {
	switch {
	case CRm > 0xf || op2 > 0x7 || Rt > 0x1f:
		err = errOverflow
	case op2 == 0x0:
		err = errUnallocated
	case op2 == 0x1 && Rt != 0x1f:
		err = errUnallocated
	case op2 == 0x2 && Rt == 0x1f:
		// encoding_CLREX_BN_barriers
	case op2 == 0x4 && Rt == 0x1f:
		// encoding_DSB_BO_barriers
	case op2 == 0x5 && Rt == 0x1f:
		// encoding_DMB_BO_barriers
	case op2 == 0x6 && Rt == 0x1f:
		// encoding_ISB_BI_barriers
	case op2 == 0x7 && Rt != 0x1f:
		err = errUnallocated
	case op2 == 0x7 && Rt == 0x1f:
		// encoding_SB_only_barriers
	case (CRm&0x2) == 0x0 && op2 == 0x1 && Rt == 0x1f:
		err = errUnallocated
	case (CRm&0x3) == 0x2 && op2 == 0x1 && Rt == 0x1f:
		// encoding_DSB_BOn_barriers
	case (CRm&0x3) == 0x3 && op2 == 0x1 && Rt == 0x1f:
		err = errUnallocated
	case CRm == 0x0 && op2 == 0x3 && Rt == 0x1f:
		// encoding_TCOMMIT_only_barriers
	case CRm == 0x1 && op2 == 0x3:
		err = errUnallocated
	case (CRm&0xe) == 0x2 && op2 == 0x3:
		err = errUnallocated
	case (CRm&0xc) == 0x4 && op2 == 0x3:
		err = errUnallocated
	case (CRm&0x8) == 0x8 && op2 == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (CRm << 8) | (op2 << 5) | Rt
	ins |= 0xd5033000

	return
}

func encode_pstate(op1, CRm, op2, Rt uint32) (ins uint32, err error) {
	switch {
	case op1 > 0x7 || CRm > 0xf || op2 > 0x7 || Rt > 0x1f:
		err = errOverflow
	case Rt != 0x1f:
		err = errUnallocated
	case Rt == 0x1f:
		// encoding_MSR_SI_pstate
	case op1 == 0x0 && op2 == 0x0 && Rt == 0x1f:
		// encoding_CFINV_M_pstate
	case op1 == 0x0 && op2 == 0x1 && Rt == 0x1f:
		// encoding_XAFLAG_M_pstate
	case op1 == 0x0 && op2 == 0x2 && Rt == 0x1f:
		// encoding_AXFLAG_M_pstate
	default:
		err = errUnmatched
	}

	ins |= (op1 << 16) | (CRm << 8) | (op2 << 5) | Rt
	ins |= 0xd5004000

	return
}

func encode_systemresult(op1, CRn, CRm, op2, Rt uint32) (ins uint32, err error) {
	switch {
	case op1 > 0x7 || CRn > 0xf || CRm > 0xf || op2 > 0x7 || Rt > 0x1f:
		err = errOverflow
	case op1 != 0x3:
		err = errUnallocated
	case op1 == 0x3 && CRn != 0x3:
		err = errUnallocated
	case op1 == 0x3 && CRn == 0x3 && op2 != 0x3:
		err = errUnallocated
	case op1 == 0x3 && CRn == 0x3 && (CRm&0xe) != 0x0 && op2 == 0x3:
		err = errUnallocated
	case op1 == 0x3 && CRn == 0x3 && CRm == 0x0 && op2 == 0x3:
		// encoding_TSTART_BR_systemresult
	case op1 == 0x3 && CRn == 0x3 && CRm == 0x1 && op2 == 0x3:
		// encoding_TTEST_BR_systemresult
	default:
		err = errUnmatched
	}

	ins |= (op1 << 16) | (CRn << 12) | (CRm << 8) | (op2 << 5) | Rt
	ins |= 0xd5200000

	return
}

func encode_systeminstrs(L, op1, CRn, CRm, op2, Rt uint32) (ins uint32, err error) {
	switch {
	case L > 0x1 || op1 > 0x7 || CRn > 0xf || CRm > 0xf || op2 > 0x7 || Rt > 0x1f:
		err = errOverflow
	case L == 0x0:
		// encoding_SYS_CR_systeminstrs
	case L == 0x1:
		// encoding_SYSL_RC_systeminstrs
	default:
		err = errUnmatched
	}

	ins |= (L << 21) | (op1 << 16) | (CRn << 12) | (CRm << 8) | (op2 << 5) | Rt
	ins |= 0xd5080000

	return
}

func encode_systemmove(L, o0, op1, CRn, CRm, op2, Rt uint32) (ins uint32, err error) {
	switch {
	case L > 0x1 || o0 > 0x1 || op1 > 0x7 || CRn > 0xf || CRm > 0xf || op2 > 0x7 || Rt > 0x1f:
		err = errOverflow
	case L == 0x0:
		// encoding_MSR_SR_systemmove
	case L == 0x1:
		// encoding_MRS_RS_systemmove
	default:
		err = errUnmatched
	}

	ins |= (L << 21) | (o0 << 19) | (op1 << 16) | (CRn << 12) | (CRm << 8) | (op2 << 5) | Rt
	ins |= 0xd5100000

	return
}

func encode_branch_reg(opc, op2, op3, Rn, op4 uint32) (ins uint32, err error) {
	switch {
	case opc > 0xf || op2 > 0x1f || op3 > 0x3f || Rn > 0x1f || op4 > 0x1f:
		err = errOverflow
	case op2 != 0x1f:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && op3 == 0x0 && op4 != 0x0:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && op3 == 0x0 && op4 == 0x0:
		// encoding_BR_64_branch_reg
	case opc == 0x0 && op2 == 0x1f && op3 == 0x1:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && op3 == 0x2 && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && op3 == 0x2 && op4 == 0x1f:
		// encoding_BRAAZ_64_branch_reg
	case opc == 0x0 && op2 == 0x1f && op3 == 0x3 && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && op3 == 0x3 && op4 == 0x1f:
		// encoding_BRABZ_64_branch_reg
	case opc == 0x0 && op2 == 0x1f && (op3&0x3c) == 0x4:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && (op3&0x38) == 0x8:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && (op3&0x30) == 0x10:
		err = errUnallocated
	case opc == 0x0 && op2 == 0x1f && (op3&0x20) == 0x20:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && op3 == 0x0 && op4 != 0x0:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && op3 == 0x0 && op4 == 0x0:
		// encoding_BLR_64_branch_reg
	case opc == 0x1 && op2 == 0x1f && op3 == 0x1:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && op3 == 0x2 && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && op3 == 0x2 && op4 == 0x1f:
		// encoding_BLRAAZ_64_branch_reg
	case opc == 0x1 && op2 == 0x1f && op3 == 0x3 && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && op3 == 0x3 && op4 == 0x1f:
		// encoding_BLRABZ_64_branch_reg
	case opc == 0x1 && op2 == 0x1f && (op3&0x3c) == 0x4:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && (op3&0x38) == 0x8:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && (op3&0x30) == 0x10:
		err = errUnallocated
	case opc == 0x1 && op2 == 0x1f && (op3&0x20) == 0x20:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && op3 == 0x0 && op4 != 0x0:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && op3 == 0x0 && op4 == 0x0:
		// encoding_RET_64R_branch_reg
	case opc == 0x2 && op2 == 0x1f && op3 == 0x1:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && op3 == 0x2 && Rn != 0x1f && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && op3 == 0x2 && Rn == 0x1f && op4 == 0x1f:
		// encoding_RETAA_64E_branch_reg
	case opc == 0x2 && op2 == 0x1f && op3 == 0x3 && Rn != 0x1f && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && op3 == 0x3 && Rn == 0x1f && op4 == 0x1f:
		// encoding_RETAB_64E_branch_reg
	case opc == 0x2 && op2 == 0x1f && (op3&0x3c) == 0x4:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && (op3&0x38) == 0x8:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && (op3&0x30) == 0x10:
		err = errUnallocated
	case opc == 0x2 && op2 == 0x1f && (op3&0x20) == 0x20:
		err = errUnallocated
	case opc == 0x3 && op2 == 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x0 && Rn != 0x1f && op4 != 0x0:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x0 && Rn != 0x1f && op4 == 0x0:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x0 && Rn == 0x1f && op4 != 0x0:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x0 && Rn == 0x1f && op4 == 0x0:
		// encoding_ERET_64E_branch_reg
	case opc == 0x4 && op2 == 0x1f && op3 == 0x1:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x2 && Rn != 0x1f && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x2 && Rn != 0x1f && op4 == 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x2 && Rn == 0x1f && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x2 && Rn == 0x1f && op4 == 0x1f:
		// encoding_ERETAA_64E_branch_reg
	case opc == 0x4 && op2 == 0x1f && op3 == 0x3 && Rn != 0x1f && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x3 && Rn != 0x1f && op4 == 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x3 && Rn == 0x1f && op4 != 0x1f:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && op3 == 0x3 && Rn == 0x1f && op4 == 0x1f:
		// encoding_ERETAB_64E_branch_reg
	case opc == 0x4 && op2 == 0x1f && (op3&0x3c) == 0x4:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && (op3&0x38) == 0x8:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && (op3&0x30) == 0x10:
		err = errUnallocated
	case opc == 0x4 && op2 == 0x1f && (op3&0x20) == 0x20:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x1f && op3 != 0x0:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x1f && op3 == 0x0 && Rn != 0x1f && op4 != 0x0:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x1f && op3 == 0x0 && Rn != 0x1f && op4 == 0x0:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x1f && op3 == 0x0 && Rn == 0x1f && op4 != 0x0:
		err = errUnallocated
	case opc == 0x5 && op2 == 0x1f && op3 == 0x0 && Rn == 0x1f && op4 == 0x0:
		// encoding_DRPS_64E_branch_reg
	case (opc&0xe) == 0x6 && op2 == 0x1f:
		err = errUnallocated
	case opc == 0x8 && op2 == 0x1f && (op3&0x3e) == 0x0:
		err = errUnallocated
	case opc == 0x8 && op2 == 0x1f && op3 == 0x2:
		// encoding_BRAA_64P_branch_reg
	case opc == 0x8 && op2 == 0x1f && op3 == 0x3:
		// encoding_BRAB_64P_branch_reg
	case opc == 0x8 && op2 == 0x1f && (op3&0x3c) == 0x4:
		err = errUnallocated
	case opc == 0x8 && op2 == 0x1f && (op3&0x38) == 0x8:
		err = errUnallocated
	case opc == 0x8 && op2 == 0x1f && (op3&0x30) == 0x10:
		err = errUnallocated
	case opc == 0x8 && op2 == 0x1f && (op3&0x20) == 0x20:
		err = errUnallocated
	case opc == 0x9 && op2 == 0x1f && (op3&0x3e) == 0x0:
		err = errUnallocated
	case opc == 0x9 && op2 == 0x1f && op3 == 0x2:
		// encoding_BLRAA_64P_branch_reg
	case opc == 0x9 && op2 == 0x1f && op3 == 0x3:
		// encoding_BLRAB_64P_branch_reg
	case opc == 0x9 && op2 == 0x1f && (op3&0x3c) == 0x4:
		err = errUnallocated
	case opc == 0x9 && op2 == 0x1f && (op3&0x38) == 0x8:
		err = errUnallocated
	case opc == 0x9 && op2 == 0x1f && (op3&0x30) == 0x10:
		err = errUnallocated
	case opc == 0x9 && op2 == 0x1f && (op3&0x20) == 0x20:
		err = errUnallocated
	case (opc&0xe) == 0xa && op2 == 0x1f:
		err = errUnallocated
	case (opc&0xc) == 0xc && op2 == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 21) | (op2 << 16) | (op3 << 10) | (Rn << 5) | op4
	ins |= 0xd6000000

	return
}

func encode_branch_imm(op, imm26 uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || imm26 > 0x3ffffff:
		err = errOverflow
	case op == 0x0:
		// encoding_B_only_branch_imm
	case op == 0x1:
		// encoding_BL_only_branch_imm
	default:
		err = errUnmatched
	}

	ins |= (op << 31) | imm26
	ins |= 0x14000000

	return
}

func encode_compbranch(sf, op, imm19, Rt uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || imm19 > 0x7ffff || Rt > 0x1f:
		err = errOverflow
	case sf == 0x0 && op == 0x0:
		// encoding_CBZ_32_compbranch
	case sf == 0x0 && op == 0x1:
		// encoding_CBNZ_32_compbranch
	case sf == 0x1 && op == 0x0:
		// encoding_CBZ_64_compbranch
	case sf == 0x1 && op == 0x1:
		// encoding_CBNZ_64_compbranch
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 24) | (imm19 << 5) | Rt
	ins |= 0x34000000

	return
}

func encode_testbranch(b5, op, b40, imm14, Rt uint32) (ins uint32, err error) {
	switch {
	case b5 > 0x1 || op > 0x1 || b40 > 0x1f || imm14 > 0x3fff || Rt > 0x1f:
		err = errOverflow
	case op == 0x0:
		// encoding_TBZ_only_testbranch
	case op == 0x1:
		// encoding_TBNZ_only_testbranch
	default:
		err = errUnmatched
	}

	ins |= (b5 << 31) | (op << 24) | (b40 << 19) | (imm14 << 5) | Rt
	ins |= 0x36000000

	return
}

func encode_comswappr(sz, L, Rs, o0, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case sz > 0x1 || L > 0x1 || Rs > 0x1f || o0 > 0x1 || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case Rt2 != 0x1f:
		err = errUnallocated
	case sz == 0x0 && L == 0x0 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASP_CP32_comswappr
	case sz == 0x0 && L == 0x0 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASPL_CP32_comswappr
	case sz == 0x0 && L == 0x1 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASPA_CP32_comswappr
	case sz == 0x0 && L == 0x1 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASPAL_CP32_comswappr
	case sz == 0x1 && L == 0x0 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASP_CP64_comswappr
	case sz == 0x1 && L == 0x0 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASPL_CP64_comswappr
	case sz == 0x1 && L == 0x1 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASPA_CP64_comswappr
	case sz == 0x1 && L == 0x1 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASPAL_CP64_comswappr
	default:
		err = errUnmatched
	}

	ins |= (sz << 30) | (L << 22) | (Rs << 16) | (o0 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x8200000

	return
}

func encode_asisdlse(Q, L, opcode, size, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || L > 0x1 || opcode > 0xf || size > 0x3 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case L == 0x0 && opcode == 0x0:
		// encoding_ST4_asisdlse_R4
	case L == 0x0 && opcode == 0x1:
		err = errUnallocated
	case L == 0x0 && opcode == 0x2:
		// encoding_ST1_asisdlse_R4_4v
	case L == 0x0 && opcode == 0x3:
		err = errUnallocated
	case L == 0x0 && opcode == 0x4:
		// encoding_ST3_asisdlse_R3
	case L == 0x0 && opcode == 0x5:
		err = errUnallocated
	case L == 0x0 && opcode == 0x6:
		// encoding_ST1_asisdlse_R3_3v
	case L == 0x0 && opcode == 0x7:
		// encoding_ST1_asisdlse_R1_1v
	case L == 0x0 && opcode == 0x8:
		// encoding_ST2_asisdlse_R2
	case L == 0x0 && opcode == 0x9:
		err = errUnallocated
	case L == 0x0 && opcode == 0xa:
		// encoding_ST1_asisdlse_R2_2v
	case L == 0x0 && opcode == 0xb:
		err = errUnallocated
	case L == 0x0 && (opcode&0xc) == 0xc:
		err = errUnallocated
	case L == 0x1 && opcode == 0x0:
		// encoding_LD4_asisdlse_R4
	case L == 0x1 && opcode == 0x1:
		err = errUnallocated
	case L == 0x1 && opcode == 0x2:
		// encoding_LD1_asisdlse_R4_4v
	case L == 0x1 && opcode == 0x3:
		err = errUnallocated
	case L == 0x1 && opcode == 0x4:
		// encoding_LD3_asisdlse_R3
	case L == 0x1 && opcode == 0x5:
		err = errUnallocated
	case L == 0x1 && opcode == 0x6:
		// encoding_LD1_asisdlse_R3_3v
	case L == 0x1 && opcode == 0x7:
		// encoding_LD1_asisdlse_R1_1v
	case L == 0x1 && opcode == 0x8:
		// encoding_LD2_asisdlse_R2
	case L == 0x1 && opcode == 0x9:
		err = errUnallocated
	case L == 0x1 && opcode == 0xa:
		// encoding_LD1_asisdlse_R2_2v
	case L == 0x1 && opcode == 0xb:
		err = errUnallocated
	case L == 0x1 && (opcode&0xc) == 0xc:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (L << 22) | (opcode << 12) | (size << 10) | (Rn << 5) | Rt
	ins |= 0xc000000

	return
}

func encode_asisdlsep(Q, L, Rm, opcode, size, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || L > 0x1 || Rm > 0x1f || opcode > 0xf || size > 0x3 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case L == 0x0 && opcode == 0x1:
		err = errUnallocated
	case L == 0x0 && opcode == 0x3:
		err = errUnallocated
	case L == 0x0 && opcode == 0x5:
		err = errUnallocated
	case L == 0x0 && opcode == 0x9:
		err = errUnallocated
	case L == 0x0 && opcode == 0xb:
		err = errUnallocated
	case L == 0x0 && (opcode&0xc) == 0xc:
		err = errUnallocated
	case L == 0x0 && Rm != 0x1f && opcode == 0x0:
		// encoding_ST4_asisdlsep_R4_r
	case L == 0x0 && Rm != 0x1f && opcode == 0x2:
		// encoding_ST1_asisdlsep_R4_r4
	case L == 0x0 && Rm != 0x1f && opcode == 0x4:
		// encoding_ST3_asisdlsep_R3_r
	case L == 0x0 && Rm != 0x1f && opcode == 0x6:
		// encoding_ST1_asisdlsep_R3_r3
	case L == 0x0 && Rm != 0x1f && opcode == 0x7:
		// encoding_ST1_asisdlsep_R1_r1
	case L == 0x0 && Rm != 0x1f && opcode == 0x8:
		// encoding_ST2_asisdlsep_R2_r
	case L == 0x0 && Rm != 0x1f && opcode == 0xa:
		// encoding_ST1_asisdlsep_R2_r2
	case L == 0x0 && Rm == 0x1f && opcode == 0x0:
		// encoding_ST4_asisdlsep_I4_i
	case L == 0x0 && Rm == 0x1f && opcode == 0x2:
		// encoding_ST1_asisdlsep_I4_i4
	case L == 0x0 && Rm == 0x1f && opcode == 0x4:
		// encoding_ST3_asisdlsep_I3_i
	case L == 0x0 && Rm == 0x1f && opcode == 0x6:
		// encoding_ST1_asisdlsep_I3_i3
	case L == 0x0 && Rm == 0x1f && opcode == 0x7:
		// encoding_ST1_asisdlsep_I1_i1
	case L == 0x0 && Rm == 0x1f && opcode == 0x8:
		// encoding_ST2_asisdlsep_I2_i
	case L == 0x0 && Rm == 0x1f && opcode == 0xa:
		// encoding_ST1_asisdlsep_I2_i2
	case L == 0x1 && opcode == 0x1:
		err = errUnallocated
	case L == 0x1 && opcode == 0x3:
		err = errUnallocated
	case L == 0x1 && opcode == 0x5:
		err = errUnallocated
	case L == 0x1 && opcode == 0x9:
		err = errUnallocated
	case L == 0x1 && opcode == 0xb:
		err = errUnallocated
	case L == 0x1 && (opcode&0xc) == 0xc:
		err = errUnallocated
	case L == 0x1 && Rm != 0x1f && opcode == 0x0:
		// encoding_LD4_asisdlsep_R4_r
	case L == 0x1 && Rm != 0x1f && opcode == 0x2:
		// encoding_LD1_asisdlsep_R4_r4
	case L == 0x1 && Rm != 0x1f && opcode == 0x4:
		// encoding_LD3_asisdlsep_R3_r
	case L == 0x1 && Rm != 0x1f && opcode == 0x6:
		// encoding_LD1_asisdlsep_R3_r3
	case L == 0x1 && Rm != 0x1f && opcode == 0x7:
		// encoding_LD1_asisdlsep_R1_r1
	case L == 0x1 && Rm != 0x1f && opcode == 0x8:
		// encoding_LD2_asisdlsep_R2_r
	case L == 0x1 && Rm != 0x1f && opcode == 0xa:
		// encoding_LD1_asisdlsep_R2_r2
	case L == 0x1 && Rm == 0x1f && opcode == 0x0:
		// encoding_LD4_asisdlsep_I4_i
	case L == 0x1 && Rm == 0x1f && opcode == 0x2:
		// encoding_LD1_asisdlsep_I4_i4
	case L == 0x1 && Rm == 0x1f && opcode == 0x4:
		// encoding_LD3_asisdlsep_I3_i
	case L == 0x1 && Rm == 0x1f && opcode == 0x6:
		// encoding_LD1_asisdlsep_I3_i3
	case L == 0x1 && Rm == 0x1f && opcode == 0x7:
		// encoding_LD1_asisdlsep_I1_i1
	case L == 0x1 && Rm == 0x1f && opcode == 0x8:
		// encoding_LD2_asisdlsep_I2_i
	case L == 0x1 && Rm == 0x1f && opcode == 0xa:
		// encoding_LD1_asisdlsep_I2_i2
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (L << 22) | (Rm << 16) | (opcode << 12) | (size << 10) | (Rn << 5) | Rt
	ins |= 0xc800000

	return
}

func encode_asisdlso(Q, L, R, opcode, S, size, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || L > 0x1 || R > 0x1 || opcode > 0x7 || S > 0x1 || size > 0x3 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case L == 0x0 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x0:
		// encoding_ST1_asisdlso_B1_1b
	case L == 0x0 && R == 0x0 && opcode == 0x1:
		// encoding_ST3_asisdlso_B3_3b
	case L == 0x0 && R == 0x0 && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_ST1_asisdlso_H1_1h
	case L == 0x0 && R == 0x0 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_ST3_asisdlso_H3_3h
	case L == 0x0 && R == 0x0 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x4 && size == 0x0:
		// encoding_ST1_asisdlso_S1_1s
	case L == 0x0 && R == 0x0 && opcode == 0x4 && (size&0x2) == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_ST1_asisdlso_D1_1d
	case L == 0x0 && R == 0x0 && opcode == 0x4 && S == 0x1 && size == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x5 && size == 0x0:
		// encoding_ST3_asisdlso_S3_3s
	case L == 0x0 && R == 0x0 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_ST3_asisdlso_D3_3d
	case L == 0x0 && R == 0x0 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x0:
		// encoding_ST2_asisdlso_B2_2b
	case L == 0x0 && R == 0x1 && opcode == 0x1:
		// encoding_ST4_asisdlso_B4_4b
	case L == 0x0 && R == 0x1 && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_ST2_asisdlso_H2_2h
	case L == 0x0 && R == 0x1 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_ST4_asisdlso_H4_4h
	case L == 0x0 && R == 0x1 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x4 && size == 0x0:
		// encoding_ST2_asisdlso_S2_2s
	case L == 0x0 && R == 0x1 && opcode == 0x4 && size == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_ST2_asisdlso_D2_2d
	case L == 0x0 && R == 0x1 && opcode == 0x4 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x4 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x5 && size == 0x0:
		// encoding_ST4_asisdlso_S4_4s
	case L == 0x0 && R == 0x1 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_ST4_asisdlso_D4_4d
	case L == 0x0 && R == 0x1 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x0:
		// encoding_LD1_asisdlso_B1_1b
	case L == 0x1 && R == 0x0 && opcode == 0x1:
		// encoding_LD3_asisdlso_B3_3b
	case L == 0x1 && R == 0x0 && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_LD1_asisdlso_H1_1h
	case L == 0x1 && R == 0x0 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_LD3_asisdlso_H3_3h
	case L == 0x1 && R == 0x0 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x4 && size == 0x0:
		// encoding_LD1_asisdlso_S1_1s
	case L == 0x1 && R == 0x0 && opcode == 0x4 && (size&0x2) == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_LD1_asisdlso_D1_1d
	case L == 0x1 && R == 0x0 && opcode == 0x4 && S == 0x1 && size == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x5 && size == 0x0:
		// encoding_LD3_asisdlso_S3_3s
	case L == 0x1 && R == 0x0 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_LD3_asisdlso_D3_3d
	case L == 0x1 && R == 0x0 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x6 && S == 0x0:
		// encoding_LD1R_asisdlso_R1
	case L == 0x1 && R == 0x0 && opcode == 0x6 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x7 && S == 0x0:
		// encoding_LD3R_asisdlso_R3
	case L == 0x1 && R == 0x0 && opcode == 0x7 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x0:
		// encoding_LD2_asisdlso_B2_2b
	case L == 0x1 && R == 0x1 && opcode == 0x1:
		// encoding_LD4_asisdlso_B4_4b
	case L == 0x1 && R == 0x1 && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_LD2_asisdlso_H2_2h
	case L == 0x1 && R == 0x1 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_LD4_asisdlso_H4_4h
	case L == 0x1 && R == 0x1 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x4 && size == 0x0:
		// encoding_LD2_asisdlso_S2_2s
	case L == 0x1 && R == 0x1 && opcode == 0x4 && size == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_LD2_asisdlso_D2_2d
	case L == 0x1 && R == 0x1 && opcode == 0x4 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x4 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x5 && size == 0x0:
		// encoding_LD4_asisdlso_S4_4s
	case L == 0x1 && R == 0x1 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_LD4_asisdlso_D4_4d
	case L == 0x1 && R == 0x1 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x6 && S == 0x0:
		// encoding_LD2R_asisdlso_R2
	case L == 0x1 && R == 0x1 && opcode == 0x6 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x7 && S == 0x0:
		// encoding_LD4R_asisdlso_R4
	case L == 0x1 && R == 0x1 && opcode == 0x7 && S == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (L << 22) | (R << 21) | (opcode << 13) | (S << 12) | (size << 10) | (Rn << 5) | Rt
	ins |= 0xd000000

	return
}

func encode_asisdlsop(Q, L, R, Rm, opcode, S, size, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || L > 0x1 || R > 0x1 || Rm > 0x1f || opcode > 0x7 || S > 0x1 || size > 0x3 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case L == 0x0 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x4 && (size&0x2) == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x4 && S == 0x1 && size == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x0:
		// encoding_ST1_asisdlsop_BX1_r1b
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x1:
		// encoding_ST3_asisdlsop_BX3_r3b
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_ST1_asisdlsop_HX1_r1h
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_ST3_asisdlsop_HX3_r3h
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_ST1_asisdlsop_SX1_r1s
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_ST1_asisdlsop_DX1_r1d
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_ST3_asisdlsop_SX3_r3s
	case L == 0x0 && R == 0x0 && Rm != 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_ST3_asisdlsop_DX3_r3d
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x0:
		// encoding_ST1_asisdlsop_B1_i1b
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x1:
		// encoding_ST3_asisdlsop_B3_i3b
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_ST1_asisdlsop_H1_i1h
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_ST3_asisdlsop_H3_i3h
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_ST1_asisdlsop_S1_i1s
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_ST1_asisdlsop_D1_i1d
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_ST3_asisdlsop_S3_i3s
	case L == 0x0 && R == 0x0 && Rm == 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_ST3_asisdlsop_D3_i3d
	case L == 0x0 && R == 0x1 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x4 && size == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x4 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x4 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x0:
		// encoding_ST2_asisdlsop_BX2_r2b
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x1:
		// encoding_ST4_asisdlsop_BX4_r4b
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_ST2_asisdlsop_HX2_r2h
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_ST4_asisdlsop_HX4_r4h
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_ST2_asisdlsop_SX2_r2s
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_ST2_asisdlsop_DX2_r2d
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_ST4_asisdlsop_SX4_r4s
	case L == 0x0 && R == 0x1 && Rm != 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_ST4_asisdlsop_DX4_r4d
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x0:
		// encoding_ST2_asisdlsop_B2_i2b
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x1:
		// encoding_ST4_asisdlsop_B4_i4b
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_ST2_asisdlsop_H2_i2h
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_ST4_asisdlsop_H4_i4h
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_ST2_asisdlsop_S2_i2s
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_ST2_asisdlsop_D2_i2d
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_ST4_asisdlsop_S4_i4s
	case L == 0x0 && R == 0x1 && Rm == 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_ST4_asisdlsop_D4_i4d
	case L == 0x1 && R == 0x0 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x4 && (size&0x2) == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x4 && S == 0x1 && size == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x6 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && opcode == 0x7 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x0:
		// encoding_LD1_asisdlsop_BX1_r1b
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x1:
		// encoding_LD3_asisdlsop_BX3_r3b
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_LD1_asisdlsop_HX1_r1h
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_LD3_asisdlsop_HX3_r3h
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_LD1_asisdlsop_SX1_r1s
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_LD1_asisdlsop_DX1_r1d
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_LD3_asisdlsop_SX3_r3s
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_LD3_asisdlsop_DX3_r3d
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x6 && S == 0x0:
		// encoding_LD1R_asisdlsop_RX1_r
	case L == 0x1 && R == 0x0 && Rm != 0x1f && opcode == 0x7 && S == 0x0:
		// encoding_LD3R_asisdlsop_RX3_r
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x0:
		// encoding_LD1_asisdlsop_B1_i1b
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x1:
		// encoding_LD3_asisdlsop_B3_i3b
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_LD1_asisdlsop_H1_i1h
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_LD3_asisdlsop_H3_i3h
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_LD1_asisdlsop_S1_i1s
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_LD1_asisdlsop_D1_i1d
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_LD3_asisdlsop_S3_i3s
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_LD3_asisdlsop_D3_i3d
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x6 && S == 0x0:
		// encoding_LD1R_asisdlsop_R1_i
	case L == 0x1 && R == 0x0 && Rm == 0x1f && opcode == 0x7 && S == 0x0:
		// encoding_LD3R_asisdlsop_R3_i
	case L == 0x1 && R == 0x1 && opcode == 0x2 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x3 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x4 && size == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x4 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x4 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x5 && size == 0x2:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x5 && S == 0x0 && size == 0x3:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x5 && S == 0x1 && (size&0x1) == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x6 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && opcode == 0x7 && S == 0x1:
		err = errUnallocated
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x0:
		// encoding_LD2_asisdlsop_BX2_r2b
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x1:
		// encoding_LD4_asisdlsop_BX4_r4b
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_LD2_asisdlsop_HX2_r2h
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_LD4_asisdlsop_HX4_r4h
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_LD2_asisdlsop_SX2_r2s
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_LD2_asisdlsop_DX2_r2d
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_LD4_asisdlsop_SX4_r4s
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_LD4_asisdlsop_DX4_r4d
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x6 && S == 0x0:
		// encoding_LD2R_asisdlsop_RX2_r
	case L == 0x1 && R == 0x1 && Rm != 0x1f && opcode == 0x7 && S == 0x0:
		// encoding_LD4R_asisdlsop_RX4_r
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x0:
		// encoding_LD2_asisdlsop_B2_i2b
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x1:
		// encoding_LD4_asisdlsop_B4_i4b
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x2 && (size&0x1) == 0x0:
		// encoding_LD2_asisdlsop_H2_i2h
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x3 && (size&0x1) == 0x0:
		// encoding_LD4_asisdlsop_H4_i4h
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x4 && size == 0x0:
		// encoding_LD2_asisdlsop_S2_i2s
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x4 && S == 0x0 && size == 0x1:
		// encoding_LD2_asisdlsop_D2_i2d
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x5 && size == 0x0:
		// encoding_LD4_asisdlsop_S4_i4s
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x5 && S == 0x0 && size == 0x1:
		// encoding_LD4_asisdlsop_D4_i4d
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x6 && S == 0x0:
		// encoding_LD2R_asisdlsop_R2_i
	case L == 0x1 && R == 0x1 && Rm == 0x1f && opcode == 0x7 && S == 0x0:
		// encoding_LD4R_asisdlsop_R4_i
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (L << 22) | (R << 21) | (Rm << 16) | (opcode << 13) | (S << 12) | (size << 10) | (Rn << 5) | Rt
	ins |= 0xd800000

	return
}

func encode_ldsttags(opc, imm9, op2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || imm9 > 0x1ff || op2 > 0x3 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case opc == 0x0 && op2 == 0x1:
		// encoding_STG_64Spost_ldsttags
	case opc == 0x0 && op2 == 0x2:
		// encoding_STG_64Soffset_ldsttags
	case opc == 0x0 && op2 == 0x3:
		// encoding_STG_64Spre_ldsttags
	case opc == 0x0 && imm9 == 0x0 && op2 == 0x0:
		// encoding_STZGM_64bulk_ldsttags
	case opc == 0x1 && op2 == 0x0:
		// encoding_LDG_64Loffset_ldsttags
	case opc == 0x1 && op2 == 0x1:
		// encoding_STZG_64Spost_ldsttags
	case opc == 0x1 && op2 == 0x2:
		// encoding_STZG_64Soffset_ldsttags
	case opc == 0x1 && op2 == 0x3:
		// encoding_STZG_64Spre_ldsttags
	case opc == 0x2 && op2 == 0x1:
		// encoding_ST2G_64Spost_ldsttags
	case opc == 0x2 && op2 == 0x2:
		// encoding_ST2G_64Soffset_ldsttags
	case opc == 0x2 && op2 == 0x3:
		// encoding_ST2G_64Spre_ldsttags
	case opc == 0x2 && imm9 != 0x0 && op2 == 0x0:
		err = errUnallocated
	case opc == 0x2 && imm9 == 0x0 && op2 == 0x0:
		// encoding_STGM_64bulk_ldsttags
	case opc == 0x3 && op2 == 0x1:
		// encoding_STZ2G_64Spost_ldsttags
	case opc == 0x3 && op2 == 0x2:
		// encoding_STZ2G_64Soffset_ldsttags
	case opc == 0x3 && op2 == 0x3:
		// encoding_STZ2G_64Spre_ldsttags
	case opc == 0x3 && imm9 != 0x0 && op2 == 0x0:
		err = errUnallocated
	case opc == 0x3 && imm9 == 0x0 && op2 == 0x0:
		// encoding_LDGM_64bulk_ldsttags
	default:
		err = errUnmatched
	}

	ins |= (opc << 22) | (imm9 << 12) | (op2 << 10) | (Rn << 5) | Rt
	ins |= 0xd9200000

	return
}

func encode_ldstexclp(sz, L, Rs, o0, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case sz > 0x1 || L > 0x1 || Rs > 0x1f || o0 > 0x1 || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case sz == 0x0 && L == 0x0 && o0 == 0x0:
		// encoding_STXP_SP32_ldstexclp
	case sz == 0x0 && L == 0x0 && o0 == 0x1:
		// encoding_STLXP_SP32_ldstexclp
	case sz == 0x0 && L == 0x1 && o0 == 0x0:
		// encoding_LDXP_LP32_ldstexclp
	case sz == 0x0 && L == 0x1 && o0 == 0x1:
		// encoding_LDAXP_LP32_ldstexclp
	case sz == 0x1 && L == 0x0 && o0 == 0x0:
		// encoding_STXP_SP64_ldstexclp
	case sz == 0x1 && L == 0x0 && o0 == 0x1:
		// encoding_STLXP_SP64_ldstexclp
	case sz == 0x1 && L == 0x1 && o0 == 0x0:
		// encoding_LDXP_LP64_ldstexclp
	case sz == 0x1 && L == 0x1 && o0 == 0x1:
		// encoding_LDAXP_LP64_ldstexclp
	default:
		err = errUnmatched
	}

	ins |= (sz << 30) | (L << 22) | (Rs << 16) | (o0 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x88200000

	return
}

func encode_ldstexclr(size, L, Rs, o0, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || L > 0x1 || Rs > 0x1f || o0 > 0x1 || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case size == 0x0 && L == 0x0 && o0 == 0x0:
		// encoding_STXRB_SR32_ldstexclr
	case size == 0x0 && L == 0x0 && o0 == 0x1:
		// encoding_STLXRB_SR32_ldstexclr
	case size == 0x0 && L == 0x1 && o0 == 0x0:
		// encoding_LDXRB_LR32_ldstexclr
	case size == 0x0 && L == 0x1 && o0 == 0x1:
		// encoding_LDAXRB_LR32_ldstexclr
	case size == 0x1 && L == 0x0 && o0 == 0x0:
		// encoding_STXRH_SR32_ldstexclr
	case size == 0x1 && L == 0x0 && o0 == 0x1:
		// encoding_STLXRH_SR32_ldstexclr
	case size == 0x1 && L == 0x1 && o0 == 0x0:
		// encoding_LDXRH_LR32_ldstexclr
	case size == 0x1 && L == 0x1 && o0 == 0x1:
		// encoding_LDAXRH_LR32_ldstexclr
	case size == 0x2 && L == 0x0 && o0 == 0x0:
		// encoding_STXR_SR32_ldstexclr
	case size == 0x2 && L == 0x0 && o0 == 0x1:
		// encoding_STLXR_SR32_ldstexclr
	case size == 0x2 && L == 0x1 && o0 == 0x0:
		// encoding_LDXR_LR32_ldstexclr
	case size == 0x2 && L == 0x1 && o0 == 0x1:
		// encoding_LDAXR_LR32_ldstexclr
	case size == 0x3 && L == 0x0 && o0 == 0x0:
		// encoding_STXR_SR64_ldstexclr
	case size == 0x3 && L == 0x0 && o0 == 0x1:
		// encoding_STLXR_SR64_ldstexclr
	case size == 0x3 && L == 0x1 && o0 == 0x0:
		// encoding_LDXR_LR64_ldstexclr
	case size == 0x3 && L == 0x1 && o0 == 0x1:
		// encoding_LDAXR_LR64_ldstexclr
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (L << 22) | (Rs << 16) | (o0 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x8000000

	return
}

func encode_ldstord(size, L, Rs, o0, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || L > 0x1 || Rs > 0x1f || o0 > 0x1 || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case size == 0x0 && L == 0x0 && o0 == 0x0:
		// encoding_STLLRB_SL32_ldstord
	case size == 0x0 && L == 0x0 && o0 == 0x1:
		// encoding_STLRB_SL32_ldstord
	case size == 0x0 && L == 0x1 && o0 == 0x0:
		// encoding_LDLARB_LR32_ldstord
	case size == 0x0 && L == 0x1 && o0 == 0x1:
		// encoding_LDARB_LR32_ldstord
	case size == 0x1 && L == 0x0 && o0 == 0x0:
		// encoding_STLLRH_SL32_ldstord
	case size == 0x1 && L == 0x0 && o0 == 0x1:
		// encoding_STLRH_SL32_ldstord
	case size == 0x1 && L == 0x1 && o0 == 0x0:
		// encoding_LDLARH_LR32_ldstord
	case size == 0x1 && L == 0x1 && o0 == 0x1:
		// encoding_LDARH_LR32_ldstord
	case size == 0x2 && L == 0x0 && o0 == 0x0:
		// encoding_STLLR_SL32_ldstord
	case size == 0x2 && L == 0x0 && o0 == 0x1:
		// encoding_STLR_SL32_ldstord
	case size == 0x2 && L == 0x1 && o0 == 0x0:
		// encoding_LDLAR_LR32_ldstord
	case size == 0x2 && L == 0x1 && o0 == 0x1:
		// encoding_LDAR_LR32_ldstord
	case size == 0x3 && L == 0x0 && o0 == 0x0:
		// encoding_STLLR_SL64_ldstord
	case size == 0x3 && L == 0x0 && o0 == 0x1:
		// encoding_STLR_SL64_ldstord
	case size == 0x3 && L == 0x1 && o0 == 0x0:
		// encoding_LDLAR_LR64_ldstord
	case size == 0x3 && L == 0x1 && o0 == 0x1:
		// encoding_LDAR_LR64_ldstord
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (L << 22) | (Rs << 16) | (o0 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x8800000

	return
}

func encode_comswap(size, L, Rs, o0, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || L > 0x1 || Rs > 0x1f || o0 > 0x1 || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case Rt2 != 0x1f:
		err = errUnallocated
	case size == 0x0 && L == 0x0 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASB_C32_comswap
	case size == 0x0 && L == 0x0 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASLB_C32_comswap
	case size == 0x0 && L == 0x1 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASAB_C32_comswap
	case size == 0x0 && L == 0x1 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASALB_C32_comswap
	case size == 0x1 && L == 0x0 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASH_C32_comswap
	case size == 0x1 && L == 0x0 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASLH_C32_comswap
	case size == 0x1 && L == 0x1 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASAH_C32_comswap
	case size == 0x1 && L == 0x1 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASALH_C32_comswap
	case size == 0x2 && L == 0x0 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CAS_C32_comswap
	case size == 0x2 && L == 0x0 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASL_C32_comswap
	case size == 0x2 && L == 0x1 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASA_C32_comswap
	case size == 0x2 && L == 0x1 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASAL_C32_comswap
	case size == 0x3 && L == 0x0 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CAS_C64_comswap
	case size == 0x3 && L == 0x0 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASL_C64_comswap
	case size == 0x3 && L == 0x1 && o0 == 0x0 && Rt2 == 0x1f:
		// encoding_CASA_C64_comswap
	case size == 0x3 && L == 0x1 && o0 == 0x1 && Rt2 == 0x1f:
		// encoding_CASAL_C64_comswap
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (L << 22) | (Rs << 16) | (o0 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x8a00000

	return
}

func encode_ldapstl_unscaled(size, opc, imm9, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opc > 0x3 || imm9 > 0x1ff || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case size == 0x0 && opc == 0x0:
		// encoding_STLURB_32_ldapstl_unscaled
	case size == 0x0 && opc == 0x1:
		// encoding_LDAPURB_32_ldapstl_unscaled
	case size == 0x0 && opc == 0x2:
		// encoding_LDAPURSB_64_ldapstl_unscaled
	case size == 0x0 && opc == 0x3:
		// encoding_LDAPURSB_32_ldapstl_unscaled
	case size == 0x1 && opc == 0x0:
		// encoding_STLURH_32_ldapstl_unscaled
	case size == 0x1 && opc == 0x1:
		// encoding_LDAPURH_32_ldapstl_unscaled
	case size == 0x1 && opc == 0x2:
		// encoding_LDAPURSH_64_ldapstl_unscaled
	case size == 0x1 && opc == 0x3:
		// encoding_LDAPURSH_32_ldapstl_unscaled
	case size == 0x2 && opc == 0x0:
		// encoding_STLUR_32_ldapstl_unscaled
	case size == 0x2 && opc == 0x1:
		// encoding_LDAPUR_32_ldapstl_unscaled
	case size == 0x2 && opc == 0x2:
		// encoding_LDAPURSW_64_ldapstl_unscaled
	case size == 0x2 && opc == 0x3:
		err = errUnallocated
	case size == 0x3 && opc == 0x0:
		// encoding_STLUR_64_ldapstl_unscaled
	case size == 0x3 && opc == 0x1:
		// encoding_LDAPUR_64_ldapstl_unscaled
	case size == 0x3 && opc == 0x2:
		err = errUnallocated
	case size == 0x3 && opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (opc << 22) | (imm9 << 12) | (Rn << 5) | Rt
	ins |= 0x19000000

	return
}

func encode_loadlit(opc, V, imm19, Rt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || V > 0x1 || imm19 > 0x7ffff || Rt > 0x1f:
		err = errOverflow
	case opc == 0x0 && V == 0x0:
		// encoding_LDR_32_loadlit
	case opc == 0x0 && V == 0x1:
		// encoding_LDR_S_loadlit
	case opc == 0x1 && V == 0x0:
		// encoding_LDR_64_loadlit
	case opc == 0x1 && V == 0x1:
		// encoding_LDR_D_loadlit
	case opc == 0x2 && V == 0x0:
		// encoding_LDRSW_64_loadlit
	case opc == 0x2 && V == 0x1:
		// encoding_LDR_Q_loadlit
	case opc == 0x3 && V == 0x0:
		// encoding_PRFM_P_loadlit
	case opc == 0x3 && V == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 30) | (V << 26) | (imm19 << 5) | Rt
	ins |= 0x18000000

	return
}

func encode_memcms(size, o0, op1, Rs, op2, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || o0 > 0x1 || op1 > 0x3 || Rs > 0x1f || op2 > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x0:
		// encoding_CPYFP_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x1:
		// encoding_CPYFPWT_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x2:
		// encoding_CPYFPRT_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x3:
		// encoding_CPYFPT_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x4:
		// encoding_CPYFPWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x5:
		// encoding_CPYFPWTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x6:
		// encoding_CPYFPRTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x7:
		// encoding_CPYFPTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x8:
		// encoding_CPYFPRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0x9:
		// encoding_CPYFPWTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0xa:
		// encoding_CPYFPRTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0xb:
		// encoding_CPYFPTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0xc:
		// encoding_CPYFPN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0xd:
		// encoding_CPYFPWTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0xe:
		// encoding_CPYFPRTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x0 && op2 == 0xf:
		// encoding_CPYFPTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x0:
		// encoding_CPYFM_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x1:
		// encoding_CPYFMWT_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x2:
		// encoding_CPYFMRT_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x3:
		// encoding_CPYFMT_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x4:
		// encoding_CPYFMWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x5:
		// encoding_CPYFMWTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x6:
		// encoding_CPYFMRTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x7:
		// encoding_CPYFMTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x8:
		// encoding_CPYFMRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0x9:
		// encoding_CPYFMWTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0xa:
		// encoding_CPYFMRTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0xb:
		// encoding_CPYFMTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0xc:
		// encoding_CPYFMN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0xd:
		// encoding_CPYFMWTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0xe:
		// encoding_CPYFMRTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x1 && op2 == 0xf:
		// encoding_CPYFMTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x0:
		// encoding_CPYFE_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x1:
		// encoding_CPYFEWT_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x2:
		// encoding_CPYFERT_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x3:
		// encoding_CPYFET_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x4:
		// encoding_CPYFEWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x5:
		// encoding_CPYFEWTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x6:
		// encoding_CPYFERTWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x7:
		// encoding_CPYFETWN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x8:
		// encoding_CPYFERN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0x9:
		// encoding_CPYFEWTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0xa:
		// encoding_CPYFERTRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0xb:
		// encoding_CPYFETRN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0xc:
		// encoding_CPYFEN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0xd:
		// encoding_CPYFEWTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0xe:
		// encoding_CPYFERTN_CPY_memcms
	case o0 == 0x0 && op1 == 0x2 && op2 == 0xf:
		// encoding_CPYFETN_CPY_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x0:
		// encoding_SETP_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x1:
		// encoding_SETPT_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x2:
		// encoding_SETPN_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x3:
		// encoding_SETPTN_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x4:
		// encoding_SETM_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x5:
		// encoding_SETMT_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x6:
		// encoding_SETMN_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x7:
		// encoding_SETMTN_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x8:
		// encoding_SETE_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0x9:
		// encoding_SETET_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0xa:
		// encoding_SETEN_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && op2 == 0xb:
		// encoding_SETETN_SET_memcms
	case o0 == 0x0 && op1 == 0x3 && (op2&0xc) == 0xc:
		err = errUnallocated
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x0:
		// encoding_CPYP_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x1:
		// encoding_CPYPWT_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x2:
		// encoding_CPYPRT_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x3:
		// encoding_CPYPT_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x4:
		// encoding_CPYPWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x5:
		// encoding_CPYPWTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x6:
		// encoding_CPYPRTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x7:
		// encoding_CPYPTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x8:
		// encoding_CPYPRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0x9:
		// encoding_CPYPWTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0xa:
		// encoding_CPYPRTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0xb:
		// encoding_CPYPTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0xc:
		// encoding_CPYPN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0xd:
		// encoding_CPYPWTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0xe:
		// encoding_CPYPRTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x0 && op2 == 0xf:
		// encoding_CPYPTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x0:
		// encoding_CPYM_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x1:
		// encoding_CPYMWT_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x2:
		// encoding_CPYMRT_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x3:
		// encoding_CPYMT_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x4:
		// encoding_CPYMWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x5:
		// encoding_CPYMWTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x6:
		// encoding_CPYMRTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x7:
		// encoding_CPYMTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x8:
		// encoding_CPYMRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0x9:
		// encoding_CPYMWTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0xa:
		// encoding_CPYMRTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0xb:
		// encoding_CPYMTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0xc:
		// encoding_CPYMN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0xd:
		// encoding_CPYMWTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0xe:
		// encoding_CPYMRTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x1 && op2 == 0xf:
		// encoding_CPYMTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x0:
		// encoding_CPYE_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x1:
		// encoding_CPYEWT_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x2:
		// encoding_CPYERT_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x3:
		// encoding_CPYET_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x4:
		// encoding_CPYEWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x5:
		// encoding_CPYEWTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x6:
		// encoding_CPYERTWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x7:
		// encoding_CPYETWN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x8:
		// encoding_CPYERN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0x9:
		// encoding_CPYEWTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0xa:
		// encoding_CPYERTRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0xb:
		// encoding_CPYETRN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0xc:
		// encoding_CPYEN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0xd:
		// encoding_CPYEWTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0xe:
		// encoding_CPYERTN_CPY_memcms
	case o0 == 0x1 && op1 == 0x2 && op2 == 0xf:
		// encoding_CPYETN_CPY_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x0:
		// encoding_SETGP_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x1:
		// encoding_SETGPT_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x2:
		// encoding_SETGPN_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x3:
		// encoding_SETGPTN_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x4:
		// encoding_SETGM_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x5:
		// encoding_SETGMT_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x6:
		// encoding_SETGMN_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x7:
		// encoding_SETGMTN_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x8:
		// encoding_SETGE_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0x9:
		// encoding_SETGET_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0xa:
		// encoding_SETGEN_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && op2 == 0xb:
		// encoding_SETGETN_SET_memcms
	case o0 == 0x1 && op1 == 0x3 && (op2&0xc) == 0xc:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (o0 << 26) | (op1 << 22) | (Rs << 16) | (op2 << 12) | (Rn << 5) | Rd
	ins |= 0x19000400

	return
}

func encode_ldstnapair_offs(opc, V, L, imm7, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || V > 0x1 || L > 0x1 || imm7 > 0x7f || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case opc == 0x0 && V == 0x0 && L == 0x0:
		// encoding_STNP_32_ldstnapair_offs
	case opc == 0x0 && V == 0x0 && L == 0x1:
		// encoding_LDNP_32_ldstnapair_offs
	case opc == 0x0 && V == 0x1 && L == 0x0:
		// encoding_STNP_S_ldstnapair_offs
	case opc == 0x0 && V == 0x1 && L == 0x1:
		// encoding_LDNP_S_ldstnapair_offs
	case opc == 0x1 && V == 0x0:
		err = errUnallocated
	case opc == 0x1 && V == 0x1 && L == 0x0:
		// encoding_STNP_D_ldstnapair_offs
	case opc == 0x1 && V == 0x1 && L == 0x1:
		// encoding_LDNP_D_ldstnapair_offs
	case opc == 0x2 && V == 0x0 && L == 0x0:
		// encoding_STNP_64_ldstnapair_offs
	case opc == 0x2 && V == 0x0 && L == 0x1:
		// encoding_LDNP_64_ldstnapair_offs
	case opc == 0x2 && V == 0x1 && L == 0x0:
		// encoding_STNP_Q_ldstnapair_offs
	case opc == 0x2 && V == 0x1 && L == 0x1:
		// encoding_LDNP_Q_ldstnapair_offs
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 30) | (V << 26) | (L << 22) | (imm7 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x28000000

	return
}

func encode_ldstpair_post(opc, V, L, imm7, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || V > 0x1 || L > 0x1 || imm7 > 0x7f || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case opc == 0x0 && V == 0x0 && L == 0x0:
		// encoding_STP_32_ldstpair_post
	case opc == 0x0 && V == 0x0 && L == 0x1:
		// encoding_LDP_32_ldstpair_post
	case opc == 0x0 && V == 0x1 && L == 0x0:
		// encoding_STP_S_ldstpair_post
	case opc == 0x0 && V == 0x1 && L == 0x1:
		// encoding_LDP_S_ldstpair_post
	case opc == 0x1 && V == 0x0 && L == 0x0:
		// encoding_STGP_64_ldstpair_post
	case opc == 0x1 && V == 0x0 && L == 0x1:
		// encoding_LDPSW_64_ldstpair_post
	case opc == 0x1 && V == 0x1 && L == 0x0:
		// encoding_STP_D_ldstpair_post
	case opc == 0x1 && V == 0x1 && L == 0x1:
		// encoding_LDP_D_ldstpair_post
	case opc == 0x2 && V == 0x0 && L == 0x0:
		// encoding_STP_64_ldstpair_post
	case opc == 0x2 && V == 0x0 && L == 0x1:
		// encoding_LDP_64_ldstpair_post
	case opc == 0x2 && V == 0x1 && L == 0x0:
		// encoding_STP_Q_ldstpair_post
	case opc == 0x2 && V == 0x1 && L == 0x1:
		// encoding_LDP_Q_ldstpair_post
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 30) | (V << 26) | (L << 22) | (imm7 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x28800000

	return
}

func encode_ldstpair_off(opc, V, L, imm7, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || V > 0x1 || L > 0x1 || imm7 > 0x7f || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case opc == 0x0 && V == 0x0 && L == 0x0:
		// encoding_STP_32_ldstpair_off
	case opc == 0x0 && V == 0x0 && L == 0x1:
		// encoding_LDP_32_ldstpair_off
	case opc == 0x0 && V == 0x1 && L == 0x0:
		// encoding_STP_S_ldstpair_off
	case opc == 0x0 && V == 0x1 && L == 0x1:
		// encoding_LDP_S_ldstpair_off
	case opc == 0x1 && V == 0x0 && L == 0x0:
		// encoding_STGP_64_ldstpair_off
	case opc == 0x1 && V == 0x0 && L == 0x1:
		// encoding_LDPSW_64_ldstpair_off
	case opc == 0x1 && V == 0x1 && L == 0x0:
		// encoding_STP_D_ldstpair_off
	case opc == 0x1 && V == 0x1 && L == 0x1:
		// encoding_LDP_D_ldstpair_off
	case opc == 0x2 && V == 0x0 && L == 0x0:
		// encoding_STP_64_ldstpair_off
	case opc == 0x2 && V == 0x0 && L == 0x1:
		// encoding_LDP_64_ldstpair_off
	case opc == 0x2 && V == 0x1 && L == 0x0:
		// encoding_STP_Q_ldstpair_off
	case opc == 0x2 && V == 0x1 && L == 0x1:
		// encoding_LDP_Q_ldstpair_off
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 30) | (V << 26) | (L << 22) | (imm7 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x29000000

	return
}

func encode_ldstpair_pre(opc, V, L, imm7, Rt2, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case opc > 0x3 || V > 0x1 || L > 0x1 || imm7 > 0x7f || Rt2 > 0x1f || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case opc == 0x0 && V == 0x0 && L == 0x0:
		// encoding_STP_32_ldstpair_pre
	case opc == 0x0 && V == 0x0 && L == 0x1:
		// encoding_LDP_32_ldstpair_pre
	case opc == 0x0 && V == 0x1 && L == 0x0:
		// encoding_STP_S_ldstpair_pre
	case opc == 0x0 && V == 0x1 && L == 0x1:
		// encoding_LDP_S_ldstpair_pre
	case opc == 0x1 && V == 0x0 && L == 0x0:
		// encoding_STGP_64_ldstpair_pre
	case opc == 0x1 && V == 0x0 && L == 0x1:
		// encoding_LDPSW_64_ldstpair_pre
	case opc == 0x1 && V == 0x1 && L == 0x0:
		// encoding_STP_D_ldstpair_pre
	case opc == 0x1 && V == 0x1 && L == 0x1:
		// encoding_LDP_D_ldstpair_pre
	case opc == 0x2 && V == 0x0 && L == 0x0:
		// encoding_STP_64_ldstpair_pre
	case opc == 0x2 && V == 0x0 && L == 0x1:
		// encoding_LDP_64_ldstpair_pre
	case opc == 0x2 && V == 0x1 && L == 0x0:
		// encoding_STP_Q_ldstpair_pre
	case opc == 0x2 && V == 0x1 && L == 0x1:
		// encoding_LDP_Q_ldstpair_pre
	case opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opc << 30) | (V << 26) | (L << 22) | (imm7 << 15) | (Rt2 << 10) | (Rn << 5) | Rt
	ins |= 0x29800000

	return
}

func encode_ldst_unscaled(size, V, opc, imm9, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || opc > 0x3 || imm9 > 0x1ff || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case (size&0x1) == 0x1 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && opc == 0x0:
		// encoding_STURB_32_ldst_unscaled
	case size == 0x0 && V == 0x0 && opc == 0x1:
		// encoding_LDURB_32_ldst_unscaled
	case size == 0x0 && V == 0x0 && opc == 0x2:
		// encoding_LDURSB_64_ldst_unscaled
	case size == 0x0 && V == 0x0 && opc == 0x3:
		// encoding_LDURSB_32_ldst_unscaled
	case size == 0x0 && V == 0x1 && opc == 0x0:
		// encoding_STUR_B_ldst_unscaled
	case size == 0x0 && V == 0x1 && opc == 0x1:
		// encoding_LDUR_B_ldst_unscaled
	case size == 0x0 && V == 0x1 && opc == 0x2:
		// encoding_STUR_Q_ldst_unscaled
	case size == 0x0 && V == 0x1 && opc == 0x3:
		// encoding_LDUR_Q_ldst_unscaled
	case size == 0x1 && V == 0x0 && opc == 0x0:
		// encoding_STURH_32_ldst_unscaled
	case size == 0x1 && V == 0x0 && opc == 0x1:
		// encoding_LDURH_32_ldst_unscaled
	case size == 0x1 && V == 0x0 && opc == 0x2:
		// encoding_LDURSH_64_ldst_unscaled
	case size == 0x1 && V == 0x0 && opc == 0x3:
		// encoding_LDURSH_32_ldst_unscaled
	case size == 0x1 && V == 0x1 && opc == 0x0:
		// encoding_STUR_H_ldst_unscaled
	case size == 0x1 && V == 0x1 && opc == 0x1:
		// encoding_LDUR_H_ldst_unscaled
	case (size&0x2) == 0x2 && V == 0x0 && opc == 0x3:
		err = errUnallocated
	case (size&0x2) == 0x2 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && opc == 0x0:
		// encoding_STUR_32_ldst_unscaled
	case size == 0x2 && V == 0x0 && opc == 0x1:
		// encoding_LDUR_32_ldst_unscaled
	case size == 0x2 && V == 0x0 && opc == 0x2:
		// encoding_LDURSW_64_ldst_unscaled
	case size == 0x2 && V == 0x1 && opc == 0x0:
		// encoding_STUR_S_ldst_unscaled
	case size == 0x2 && V == 0x1 && opc == 0x1:
		// encoding_LDUR_S_ldst_unscaled
	case size == 0x3 && V == 0x0 && opc == 0x0:
		// encoding_STUR_64_ldst_unscaled
	case size == 0x3 && V == 0x0 && opc == 0x1:
		// encoding_LDUR_64_ldst_unscaled
	case size == 0x3 && V == 0x0 && opc == 0x2:
		// encoding_PRFUM_P_ldst_unscaled
	case size == 0x3 && V == 0x1 && opc == 0x0:
		// encoding_STUR_D_ldst_unscaled
	case size == 0x3 && V == 0x1 && opc == 0x1:
		// encoding_LDUR_D_ldst_unscaled
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (opc << 22) | (imm9 << 12) | (Rn << 5) | Rt
	ins |= 0x38000000

	return
}

func encode_ldst_immpost(size, V, opc, imm9, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || opc > 0x3 || imm9 > 0x1ff || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case (size&0x1) == 0x1 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && opc == 0x0:
		// encoding_STRB_32_ldst_immpost
	case size == 0x0 && V == 0x0 && opc == 0x1:
		// encoding_LDRB_32_ldst_immpost
	case size == 0x0 && V == 0x0 && opc == 0x2:
		// encoding_LDRSB_64_ldst_immpost
	case size == 0x0 && V == 0x0 && opc == 0x3:
		// encoding_LDRSB_32_ldst_immpost
	case size == 0x0 && V == 0x1 && opc == 0x0:
		// encoding_STR_B_ldst_immpost
	case size == 0x0 && V == 0x1 && opc == 0x1:
		// encoding_LDR_B_ldst_immpost
	case size == 0x0 && V == 0x1 && opc == 0x2:
		// encoding_STR_Q_ldst_immpost
	case size == 0x0 && V == 0x1 && opc == 0x3:
		// encoding_LDR_Q_ldst_immpost
	case size == 0x1 && V == 0x0 && opc == 0x0:
		// encoding_STRH_32_ldst_immpost
	case size == 0x1 && V == 0x0 && opc == 0x1:
		// encoding_LDRH_32_ldst_immpost
	case size == 0x1 && V == 0x0 && opc == 0x2:
		// encoding_LDRSH_64_ldst_immpost
	case size == 0x1 && V == 0x0 && opc == 0x3:
		// encoding_LDRSH_32_ldst_immpost
	case size == 0x1 && V == 0x1 && opc == 0x0:
		// encoding_STR_H_ldst_immpost
	case size == 0x1 && V == 0x1 && opc == 0x1:
		// encoding_LDR_H_ldst_immpost
	case (size&0x2) == 0x2 && V == 0x0 && opc == 0x3:
		err = errUnallocated
	case (size&0x2) == 0x2 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && opc == 0x0:
		// encoding_STR_32_ldst_immpost
	case size == 0x2 && V == 0x0 && opc == 0x1:
		// encoding_LDR_32_ldst_immpost
	case size == 0x2 && V == 0x0 && opc == 0x2:
		// encoding_LDRSW_64_ldst_immpost
	case size == 0x2 && V == 0x1 && opc == 0x0:
		// encoding_STR_S_ldst_immpost
	case size == 0x2 && V == 0x1 && opc == 0x1:
		// encoding_LDR_S_ldst_immpost
	case size == 0x3 && V == 0x0 && opc == 0x0:
		// encoding_STR_64_ldst_immpost
	case size == 0x3 && V == 0x0 && opc == 0x1:
		// encoding_LDR_64_ldst_immpost
	case size == 0x3 && V == 0x0 && opc == 0x2:
		err = errUnallocated
	case size == 0x3 && V == 0x1 && opc == 0x0:
		// encoding_STR_D_ldst_immpost
	case size == 0x3 && V == 0x1 && opc == 0x1:
		// encoding_LDR_D_ldst_immpost
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (opc << 22) | (imm9 << 12) | (Rn << 5) | Rt
	ins |= 0x38000400

	return
}

func encode_ldst_unpriv(size, V, opc, imm9, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || opc > 0x3 || imm9 > 0x1ff || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case V == 0x1:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && opc == 0x0:
		// encoding_STTRB_32_ldst_unpriv
	case size == 0x0 && V == 0x0 && opc == 0x1:
		// encoding_LDTRB_32_ldst_unpriv
	case size == 0x0 && V == 0x0 && opc == 0x2:
		// encoding_LDTRSB_64_ldst_unpriv
	case size == 0x0 && V == 0x0 && opc == 0x3:
		// encoding_LDTRSB_32_ldst_unpriv
	case size == 0x1 && V == 0x0 && opc == 0x0:
		// encoding_STTRH_32_ldst_unpriv
	case size == 0x1 && V == 0x0 && opc == 0x1:
		// encoding_LDTRH_32_ldst_unpriv
	case size == 0x1 && V == 0x0 && opc == 0x2:
		// encoding_LDTRSH_64_ldst_unpriv
	case size == 0x1 && V == 0x0 && opc == 0x3:
		// encoding_LDTRSH_32_ldst_unpriv
	case (size&0x2) == 0x2 && V == 0x0 && opc == 0x3:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && opc == 0x0:
		// encoding_STTR_32_ldst_unpriv
	case size == 0x2 && V == 0x0 && opc == 0x1:
		// encoding_LDTR_32_ldst_unpriv
	case size == 0x2 && V == 0x0 && opc == 0x2:
		// encoding_LDTRSW_64_ldst_unpriv
	case size == 0x3 && V == 0x0 && opc == 0x0:
		// encoding_STTR_64_ldst_unpriv
	case size == 0x3 && V == 0x0 && opc == 0x1:
		// encoding_LDTR_64_ldst_unpriv
	case size == 0x3 && V == 0x0 && opc == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (opc << 22) | (imm9 << 12) | (Rn << 5) | Rt
	ins |= 0x38000800

	return
}

func encode_ldst_immpre(size, V, opc, imm9, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || opc > 0x3 || imm9 > 0x1ff || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case (size&0x1) == 0x1 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && opc == 0x0:
		// encoding_STRB_32_ldst_immpre
	case size == 0x0 && V == 0x0 && opc == 0x1:
		// encoding_LDRB_32_ldst_immpre
	case size == 0x0 && V == 0x0 && opc == 0x2:
		// encoding_LDRSB_64_ldst_immpre
	case size == 0x0 && V == 0x0 && opc == 0x3:
		// encoding_LDRSB_32_ldst_immpre
	case size == 0x0 && V == 0x1 && opc == 0x0:
		// encoding_STR_B_ldst_immpre
	case size == 0x0 && V == 0x1 && opc == 0x1:
		// encoding_LDR_B_ldst_immpre
	case size == 0x0 && V == 0x1 && opc == 0x2:
		// encoding_STR_Q_ldst_immpre
	case size == 0x0 && V == 0x1 && opc == 0x3:
		// encoding_LDR_Q_ldst_immpre
	case size == 0x1 && V == 0x0 && opc == 0x0:
		// encoding_STRH_32_ldst_immpre
	case size == 0x1 && V == 0x0 && opc == 0x1:
		// encoding_LDRH_32_ldst_immpre
	case size == 0x1 && V == 0x0 && opc == 0x2:
		// encoding_LDRSH_64_ldst_immpre
	case size == 0x1 && V == 0x0 && opc == 0x3:
		// encoding_LDRSH_32_ldst_immpre
	case size == 0x1 && V == 0x1 && opc == 0x0:
		// encoding_STR_H_ldst_immpre
	case size == 0x1 && V == 0x1 && opc == 0x1:
		// encoding_LDR_H_ldst_immpre
	case (size&0x2) == 0x2 && V == 0x0 && opc == 0x3:
		err = errUnallocated
	case (size&0x2) == 0x2 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && opc == 0x0:
		// encoding_STR_32_ldst_immpre
	case size == 0x2 && V == 0x0 && opc == 0x1:
		// encoding_LDR_32_ldst_immpre
	case size == 0x2 && V == 0x0 && opc == 0x2:
		// encoding_LDRSW_64_ldst_immpre
	case size == 0x2 && V == 0x1 && opc == 0x0:
		// encoding_STR_S_ldst_immpre
	case size == 0x2 && V == 0x1 && opc == 0x1:
		// encoding_LDR_S_ldst_immpre
	case size == 0x3 && V == 0x0 && opc == 0x0:
		// encoding_STR_64_ldst_immpre
	case size == 0x3 && V == 0x0 && opc == 0x1:
		// encoding_LDR_64_ldst_immpre
	case size == 0x3 && V == 0x0 && opc == 0x2:
		err = errUnallocated
	case size == 0x3 && V == 0x1 && opc == 0x0:
		// encoding_STR_D_ldst_immpre
	case size == 0x3 && V == 0x1 && opc == 0x1:
		// encoding_LDR_D_ldst_immpre
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (opc << 22) | (imm9 << 12) | (Rn << 5) | Rt
	ins |= 0x38000c00

	return
}

func encode_memop(size, V, A, R, Rs, o3, opc, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || A > 0x1 || R > 0x1 || Rs > 0x1f || o3 > 0x1 || opc > 0x7 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case V == 0x0 && o3 == 0x1 && (opc&0x6) == 0x6:
		err = errUnallocated
	case V == 0x0 && A == 0x0 && o3 == 0x1 && opc == 0x4:
		err = errUnallocated
	case V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x1:
		err = errUnallocated
	case V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x2:
		err = errUnallocated
	case V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x3:
		err = errUnallocated
	case V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x5:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x1:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x2:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x3:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x5:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x1:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x2:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x3:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x4:
		err = errUnallocated
	case V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x5:
		err = errUnallocated
	case V == 0x1:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x1:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x2:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x3:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x5:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPLB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPAB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x4:
		// encoding_LDAPRB_32L_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINALB_32_memop
	case size == 0x0 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPALB_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x1:
		err = errUnallocated
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x2:
		err = errUnallocated
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x3:
		err = errUnallocated
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x5:
		err = errUnallocated
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPLH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPAH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x4:
		// encoding_LDAPRH_32L_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINALH_32_memop
	case size == 0x1 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPALH_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADD_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLR_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEOR_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSET_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAX_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMIN_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAX_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMIN_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWP_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x1:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x2:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x3:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x5:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPA_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x4:
		// encoding_LDAPR_32L_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINAL_32_memop
	case size == 0x2 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPAL_32_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADD_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLR_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEOR_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSET_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAX_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMIN_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAX_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMIN_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWP_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x2:
		// encoding_ST64BV0_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && o3 == 0x1 && opc == 0x3:
		// encoding_ST64BV_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && Rs == 0x1f && o3 == 0x1 && opc == 0x1:
		// encoding_ST64B_64L_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x0 && Rs == 0x1f && o3 == 0x1 && opc == 0x5:
		// encoding_LD64B_64L_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x0 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPA_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x0 && o3 == 0x1 && opc == 0x4:
		// encoding_LDAPR_64L_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x0:
		// encoding_LDADDAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x1:
		// encoding_LDCLRAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x2:
		// encoding_LDEORAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x3:
		// encoding_LDSETAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x4:
		// encoding_LDSMAXAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x5:
		// encoding_LDSMINAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x6:
		// encoding_LDUMAXAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x0 && opc == 0x7:
		// encoding_LDUMINAL_64_memop
	case size == 0x3 && V == 0x0 && A == 0x1 && R == 0x1 && o3 == 0x1 && opc == 0x0:
		// encoding_SWPAL_64_memop
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (A << 23) | (R << 22) | (Rs << 16) | (o3 << 15) | (opc << 12) | (Rn << 5) | Rt
	ins |= 0x38200000

	return
}

func encode_ldst_regoff(size, V, opc, Rm, option, S, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || opc > 0x3 || Rm > 0x1f || option > 0x7 || S > 0x1 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case (size&0x1) == 0x1 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && opc == 0x0 && option != 0x3:
		// encoding_STRB_32B_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x0 && option == 0x3:
		// encoding_STRB_32BL_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x1 && option != 0x3:
		// encoding_LDRB_32B_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x1 && option == 0x3:
		// encoding_LDRB_32BL_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x2 && option != 0x3:
		// encoding_LDRSB_64B_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x2 && option == 0x3:
		// encoding_LDRSB_64BL_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x3 && option != 0x3:
		// encoding_LDRSB_32B_ldst_regoff
	case size == 0x0 && V == 0x0 && opc == 0x3 && option == 0x3:
		// encoding_LDRSB_32BL_ldst_regoff
	case size == 0x0 && V == 0x1 && opc == 0x0 && option != 0x3:
		// encoding_STR_B_ldst_regoff
	case size == 0x0 && V == 0x1 && opc == 0x0 && option == 0x3:
		// encoding_STR_BL_ldst_regoff
	case size == 0x0 && V == 0x1 && opc == 0x1 && option != 0x3:
		// encoding_LDR_B_ldst_regoff
	case size == 0x0 && V == 0x1 && opc == 0x1 && option == 0x3:
		// encoding_LDR_BL_ldst_regoff
	case size == 0x0 && V == 0x1 && opc == 0x2:
		// encoding_STR_Q_ldst_regoff
	case size == 0x0 && V == 0x1 && opc == 0x3:
		// encoding_LDR_Q_ldst_regoff
	case size == 0x1 && V == 0x0 && opc == 0x0:
		// encoding_STRH_32_ldst_regoff
	case size == 0x1 && V == 0x0 && opc == 0x1:
		// encoding_LDRH_32_ldst_regoff
	case size == 0x1 && V == 0x0 && opc == 0x2:
		// encoding_LDRSH_64_ldst_regoff
	case size == 0x1 && V == 0x0 && opc == 0x3:
		// encoding_LDRSH_32_ldst_regoff
	case size == 0x1 && V == 0x1 && opc == 0x0:
		// encoding_STR_H_ldst_regoff
	case size == 0x1 && V == 0x1 && opc == 0x1:
		// encoding_LDR_H_ldst_regoff
	case (size&0x2) == 0x2 && V == 0x0 && opc == 0x3:
		err = errUnallocated
	case (size&0x2) == 0x2 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && opc == 0x0:
		// encoding_STR_32_ldst_regoff
	case size == 0x2 && V == 0x0 && opc == 0x1:
		// encoding_LDR_32_ldst_regoff
	case size == 0x2 && V == 0x0 && opc == 0x2:
		// encoding_LDRSW_64_ldst_regoff
	case size == 0x2 && V == 0x1 && opc == 0x0:
		// encoding_STR_S_ldst_regoff
	case size == 0x2 && V == 0x1 && opc == 0x1:
		// encoding_LDR_S_ldst_regoff
	case size == 0x3 && V == 0x0 && opc == 0x0:
		// encoding_STR_64_ldst_regoff
	case size == 0x3 && V == 0x0 && opc == 0x1:
		// encoding_LDR_64_ldst_regoff
	case size == 0x3 && V == 0x0 && opc == 0x2:
		// encoding_PRFM_P_ldst_regoff
	case size == 0x3 && V == 0x1 && opc == 0x0:
		// encoding_STR_D_ldst_regoff
	case size == 0x3 && V == 0x1 && opc == 0x1:
		// encoding_LDR_D_ldst_regoff
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (opc << 22) | (Rm << 16) | (option << 13) | (S << 12) | (Rn << 5) | Rt
	ins |= 0x38200800

	return
}

func encode_ldst_pac(size, V, M, S, imm9, W, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || M > 0x1 || S > 0x1 || imm9 > 0x1ff || W > 0x1 || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case size != 0x3:
		err = errUnallocated
	case size == 0x3 && V == 0x0 && M == 0x0 && W == 0x0:
		// encoding_LDRAA_64_ldst_pac
	case size == 0x3 && V == 0x0 && M == 0x0 && W == 0x1:
		// encoding_LDRAA_64W_ldst_pac
	case size == 0x3 && V == 0x0 && M == 0x1 && W == 0x0:
		// encoding_LDRAB_64_ldst_pac
	case size == 0x3 && V == 0x0 && M == 0x1 && W == 0x1:
		// encoding_LDRAB_64W_ldst_pac
	case size == 0x3 && V == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (M << 23) | (S << 22) | (imm9 << 12) | (W << 11) | (Rn << 5) | Rt
	ins |= 0x38200400

	return
}

func encode_ldst_pos(size, V, opc, imm12, Rn, Rt uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || V > 0x1 || opc > 0x3 || imm12 > 0xfff || Rn > 0x1f || Rt > 0x1f:
		err = errOverflow
	case (size&0x1) == 0x1 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x0 && V == 0x0 && opc == 0x0:
		// encoding_STRB_32_ldst_pos
	case size == 0x0 && V == 0x0 && opc == 0x1:
		// encoding_LDRB_32_ldst_pos
	case size == 0x0 && V == 0x0 && opc == 0x2:
		// encoding_LDRSB_64_ldst_pos
	case size == 0x0 && V == 0x0 && opc == 0x3:
		// encoding_LDRSB_32_ldst_pos
	case size == 0x0 && V == 0x1 && opc == 0x0:
		// encoding_STR_B_ldst_pos
	case size == 0x0 && V == 0x1 && opc == 0x1:
		// encoding_LDR_B_ldst_pos
	case size == 0x0 && V == 0x1 && opc == 0x2:
		// encoding_STR_Q_ldst_pos
	case size == 0x0 && V == 0x1 && opc == 0x3:
		// encoding_LDR_Q_ldst_pos
	case size == 0x1 && V == 0x0 && opc == 0x0:
		// encoding_STRH_32_ldst_pos
	case size == 0x1 && V == 0x0 && opc == 0x1:
		// encoding_LDRH_32_ldst_pos
	case size == 0x1 && V == 0x0 && opc == 0x2:
		// encoding_LDRSH_64_ldst_pos
	case size == 0x1 && V == 0x0 && opc == 0x3:
		// encoding_LDRSH_32_ldst_pos
	case size == 0x1 && V == 0x1 && opc == 0x0:
		// encoding_STR_H_ldst_pos
	case size == 0x1 && V == 0x1 && opc == 0x1:
		// encoding_LDR_H_ldst_pos
	case (size&0x2) == 0x2 && V == 0x0 && opc == 0x3:
		err = errUnallocated
	case (size&0x2) == 0x2 && V == 0x1 && (opc&0x2) == 0x2:
		err = errUnallocated
	case size == 0x2 && V == 0x0 && opc == 0x0:
		// encoding_STR_32_ldst_pos
	case size == 0x2 && V == 0x0 && opc == 0x1:
		// encoding_LDR_32_ldst_pos
	case size == 0x2 && V == 0x0 && opc == 0x2:
		// encoding_LDRSW_64_ldst_pos
	case size == 0x2 && V == 0x1 && opc == 0x0:
		// encoding_STR_S_ldst_pos
	case size == 0x2 && V == 0x1 && opc == 0x1:
		// encoding_LDR_S_ldst_pos
	case size == 0x3 && V == 0x0 && opc == 0x0:
		// encoding_STR_64_ldst_pos
	case size == 0x3 && V == 0x0 && opc == 0x1:
		// encoding_LDR_64_ldst_pos
	case size == 0x3 && V == 0x0 && opc == 0x2:
		// encoding_PRFM_P_ldst_pos
	case size == 0x3 && V == 0x1 && opc == 0x0:
		// encoding_STR_D_ldst_pos
	case size == 0x3 && V == 0x1 && opc == 0x1:
		// encoding_LDR_D_ldst_pos
	default:
		err = errUnmatched
	}

	ins |= (size << 30) | (V << 26) | (opc << 22) | (imm12 << 10) | (Rn << 5) | Rt
	ins |= 0x39000000

	return
}

func encode_dp_2src(sf, S, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || S > 0x1 || Rm > 0x1f || opcode > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x1:
		err = errUnallocated
	case (opcode & 0x38) == 0x18:
		err = errUnallocated
	case (opcode & 0x20) == 0x20:
		err = errUnallocated
	case S == 0x0 && (opcode&0x3e) == 0x6:
		err = errUnallocated
	case S == 0x0 && opcode == 0xd:
		err = errUnallocated
	case S == 0x0 && (opcode&0x3e) == 0xe:
		err = errUnallocated
	case S == 0x1 && (opcode&0x3e) == 0x2:
		err = errUnallocated
	case S == 0x1 && (opcode&0x3c) == 0x4:
		err = errUnallocated
	case S == 0x1 && (opcode&0x38) == 0x8:
		err = errUnallocated
	case S == 0x1 && (opcode&0x30) == 0x10:
		err = errUnallocated
	case sf == 0x0 && opcode == 0x0:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && opcode == 0x2:
		// encoding_UDIV_32_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x3:
		// encoding_SDIV_32_dp_2src
	case sf == 0x0 && S == 0x0 && (opcode&0x3e) == 0x4:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && opcode == 0x8:
		// encoding_LSLV_32_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x9:
		// encoding_LSRV_32_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0xa:
		// encoding_ASRV_32_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0xb:
		// encoding_RORV_32_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0xc:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && (opcode&0x3b) == 0x13:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && opcode == 0x10:
		// encoding_CRC32B_32C_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x11:
		// encoding_CRC32H_32C_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x12:
		// encoding_CRC32W_32C_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x14:
		// encoding_CRC32CB_32C_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x15:
		// encoding_CRC32CH_32C_dp_2src
	case sf == 0x0 && S == 0x0 && opcode == 0x16:
		// encoding_CRC32CW_32C_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x0:
		// encoding_SUBP_64S_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x2:
		// encoding_UDIV_64_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x3:
		// encoding_SDIV_64_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x4:
		// encoding_IRG_64I_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x5:
		// encoding_GMI_64G_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x8:
		// encoding_LSLV_64_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x9:
		// encoding_LSRV_64_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0xa:
		// encoding_ASRV_64_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0xb:
		// encoding_RORV_64_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0xc:
		// encoding_PACGA_64P_dp_2src
	case sf == 0x1 && S == 0x0 && (opcode&0x39) == 0x10:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && (opcode&0x3a) == 0x10:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && opcode == 0x13:
		// encoding_CRC32X_64C_dp_2src
	case sf == 0x1 && S == 0x0 && opcode == 0x17:
		// encoding_CRC32CX_64C_dp_2src
	case sf == 0x1 && S == 0x1 && opcode == 0x0:
		// encoding_SUBPS_64S_dp_2src
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (S << 29) | (Rm << 16) | (opcode << 10) | (Rn << 5) | Rd
	ins |= 0x1ac00000

	return
}

func encode_dp_1src(sf, S, opcode2, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || S > 0x1 || opcode2 > 0x1f || opcode > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x20) == 0x20:
		err = errUnallocated
	case (opcode2 & 0x2) == 0x2:
		err = errUnallocated
	case (opcode2 & 0x4) == 0x4:
		err = errUnallocated
	case (opcode2 & 0x8) == 0x8:
		err = errUnallocated
	case (opcode2 & 0x10) == 0x10:
		err = errUnallocated
	case S == 0x0 && opcode2 == 0x0 && (opcode&0x3e) == 0x6:
		err = errUnallocated
	case S == 0x0 && opcode2 == 0x0 && (opcode&0x38) == 0x8:
		err = errUnallocated
	case S == 0x0 && opcode2 == 0x0 && (opcode&0x30) == 0x10:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case sf == 0x0 && opcode2 == 0x1:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && opcode2 == 0x0 && opcode == 0x0:
		// encoding_RBIT_32_dp_1src
	case sf == 0x0 && S == 0x0 && opcode2 == 0x0 && opcode == 0x1:
		// encoding_REV16_32_dp_1src
	case sf == 0x0 && S == 0x0 && opcode2 == 0x0 && opcode == 0x2:
		// encoding_REV_32_dp_1src
	case sf == 0x0 && S == 0x0 && opcode2 == 0x0 && opcode == 0x3:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && opcode2 == 0x0 && opcode == 0x4:
		// encoding_CLZ_32_dp_1src
	case sf == 0x0 && S == 0x0 && opcode2 == 0x0 && opcode == 0x5:
		// encoding_CLS_32_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x0 && opcode == 0x0:
		// encoding_RBIT_64_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x0 && opcode == 0x1:
		// encoding_REV16_64_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x0 && opcode == 0x2:
		// encoding_REV32_64_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x0 && opcode == 0x3:
		// encoding_REV_64_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x0 && opcode == 0x4:
		// encoding_CLZ_64_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x0 && opcode == 0x5:
		// encoding_CLS_64_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x0:
		// encoding_PACIA_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x1:
		// encoding_PACIB_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x2:
		// encoding_PACDA_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x3:
		// encoding_PACDB_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x4:
		// encoding_AUTIA_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x5:
		// encoding_AUTIB_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x6:
		// encoding_AUTDA_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x7:
		// encoding_AUTDB_64P_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x8 && Rn == 0x1f:
		// encoding_PACIZA_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x9 && Rn == 0x1f:
		// encoding_PACIZB_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0xa && Rn == 0x1f:
		// encoding_PACDZA_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0xb && Rn == 0x1f:
		// encoding_PACDZB_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0xc && Rn == 0x1f:
		// encoding_AUTIZA_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0xd && Rn == 0x1f:
		// encoding_AUTIZB_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0xe && Rn == 0x1f:
		// encoding_AUTDZA_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0xf && Rn == 0x1f:
		// encoding_AUTDZB_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x10 && Rn == 0x1f:
		// encoding_XPACI_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && opcode == 0x11 && Rn == 0x1f:
		// encoding_XPACD_64Z_dp_1src
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && (opcode&0x3e) == 0x12:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && (opcode&0x3c) == 0x14:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && opcode2 == 0x1 && (opcode&0x38) == 0x18:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (S << 29) | (opcode2 << 16) | (opcode << 10) | (Rn << 5) | Rd
	ins |= 0x5ac00000

	return
}

func encode_log_shift(sf, opc, shift, N, Rm, imm6, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || opc > 0x3 || shift > 0x3 || N > 0x1 || Rm > 0x1f || imm6 > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case sf == 0x0 && (imm6&0x20) == 0x20:
		err = errUnallocated
	case sf == 0x0 && opc == 0x0 && N == 0x0:
		// encoding_AND_32_log_shift
	case sf == 0x0 && opc == 0x0 && N == 0x1:
		// encoding_BIC_32_log_shift
	case sf == 0x0 && opc == 0x1 && N == 0x0:
		// encoding_ORR_32_log_shift
	case sf == 0x0 && opc == 0x1 && N == 0x1:
		// encoding_ORN_32_log_shift
	case sf == 0x0 && opc == 0x2 && N == 0x0:
		// encoding_EOR_32_log_shift
	case sf == 0x0 && opc == 0x2 && N == 0x1:
		// encoding_EON_32_log_shift
	case sf == 0x0 && opc == 0x3 && N == 0x0:
		// encoding_ANDS_32_log_shift
	case sf == 0x0 && opc == 0x3 && N == 0x1:
		// encoding_BICS_32_log_shift
	case sf == 0x1 && opc == 0x0 && N == 0x0:
		// encoding_AND_64_log_shift
	case sf == 0x1 && opc == 0x0 && N == 0x1:
		// encoding_BIC_64_log_shift
	case sf == 0x1 && opc == 0x1 && N == 0x0:
		// encoding_ORR_64_log_shift
	case sf == 0x1 && opc == 0x1 && N == 0x1:
		// encoding_ORN_64_log_shift
	case sf == 0x1 && opc == 0x2 && N == 0x0:
		// encoding_EOR_64_log_shift
	case sf == 0x1 && opc == 0x2 && N == 0x1:
		// encoding_EON_64_log_shift
	case sf == 0x1 && opc == 0x3 && N == 0x0:
		// encoding_ANDS_64_log_shift
	case sf == 0x1 && opc == 0x3 && N == 0x1:
		// encoding_BICS_64_log_shift
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (opc << 29) | (shift << 22) | (N << 21) | (Rm << 16) | (imm6 << 10) | (Rn << 5) | Rd
	ins |= 0xa000000

	return
}

func encode_addsub_shift(sf, op, S, shift, Rm, imm6, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || shift > 0x3 || Rm > 0x1f || imm6 > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case shift == 0x3:
		err = errUnallocated
	case sf == 0x0 && (imm6&0x20) == 0x20:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x0:
		// encoding_ADD_32_addsub_shift
	case sf == 0x0 && op == 0x0 && S == 0x1:
		// encoding_ADDS_32_addsub_shift
	case sf == 0x0 && op == 0x1 && S == 0x0:
		// encoding_SUB_32_addsub_shift
	case sf == 0x0 && op == 0x1 && S == 0x1:
		// encoding_SUBS_32_addsub_shift
	case sf == 0x1 && op == 0x0 && S == 0x0:
		// encoding_ADD_64_addsub_shift
	case sf == 0x1 && op == 0x0 && S == 0x1:
		// encoding_ADDS_64_addsub_shift
	case sf == 0x1 && op == 0x1 && S == 0x0:
		// encoding_SUB_64_addsub_shift
	case sf == 0x1 && op == 0x1 && S == 0x1:
		// encoding_SUBS_64_addsub_shift
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (shift << 22) | (Rm << 16) | (imm6 << 10) | (Rn << 5) | Rd
	ins |= 0xb000000

	return
}

func encode_addsub_ext(sf, op, S, opt, Rm, option, imm3, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || opt > 0x3 || Rm > 0x1f || option > 0x7 || imm3 > 0x7 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (imm3 & 0x5) == 0x5:
		err = errUnallocated
	case (imm3 & 0x6) == 0x6:
		err = errUnallocated
	case (opt & 0x1) == 0x1:
		err = errUnallocated
	case (opt & 0x2) == 0x2:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x0 && opt == 0x0:
		// encoding_ADD_32_addsub_ext
	case sf == 0x0 && op == 0x0 && S == 0x1 && opt == 0x0:
		// encoding_ADDS_32S_addsub_ext
	case sf == 0x0 && op == 0x1 && S == 0x0 && opt == 0x0:
		// encoding_SUB_32_addsub_ext
	case sf == 0x0 && op == 0x1 && S == 0x1 && opt == 0x0:
		// encoding_SUBS_32S_addsub_ext
	case sf == 0x1 && op == 0x0 && S == 0x0 && opt == 0x0:
		// encoding_ADD_64_addsub_ext
	case sf == 0x1 && op == 0x0 && S == 0x1 && opt == 0x0:
		// encoding_ADDS_64S_addsub_ext
	case sf == 0x1 && op == 0x1 && S == 0x0 && opt == 0x0:
		// encoding_SUB_64_addsub_ext
	case sf == 0x1 && op == 0x1 && S == 0x1 && opt == 0x0:
		// encoding_SUBS_64S_addsub_ext
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (opt << 22) | (Rm << 16) | (option << 13) | (imm3 << 10) | (Rn << 5) | Rd
	ins |= 0xb200000

	return
}

func encode_addsub_carry(sf, op, S, Rm, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || Rm > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case sf == 0x0 && op == 0x0 && S == 0x0:
		// encoding_ADC_32_addsub_carry
	case sf == 0x0 && op == 0x0 && S == 0x1:
		// encoding_ADCS_32_addsub_carry
	case sf == 0x0 && op == 0x1 && S == 0x0:
		// encoding_SBC_32_addsub_carry
	case sf == 0x0 && op == 0x1 && S == 0x1:
		// encoding_SBCS_32_addsub_carry
	case sf == 0x1 && op == 0x0 && S == 0x0:
		// encoding_ADC_64_addsub_carry
	case sf == 0x1 && op == 0x0 && S == 0x1:
		// encoding_ADCS_64_addsub_carry
	case sf == 0x1 && op == 0x1 && S == 0x0:
		// encoding_SBC_64_addsub_carry
	case sf == 0x1 && op == 0x1 && S == 0x1:
		// encoding_SBCS_64_addsub_carry
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (Rm << 16) | (Rn << 5) | Rd
	ins |= 0x1a000000

	return
}

func encode_rmif(sf, op, S, imm6, Rn, o2, mask uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || imm6 > 0x3f || Rn > 0x1f || o2 > 0x1 || mask > 0xf:
		err = errOverflow
	case sf == 0x0:
		err = errUnallocated
	case sf == 0x1 && op == 0x0 && S == 0x0:
		err = errUnallocated
	case sf == 0x1 && op == 0x0 && S == 0x1 && o2 == 0x0:
		// encoding_RMIF_only_rmif
	case sf == 0x1 && op == 0x0 && S == 0x1 && o2 == 0x1:
		err = errUnallocated
	case sf == 0x1 && op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (imm6 << 15) | (Rn << 5) | (o2 << 4) | mask
	ins |= 0x1a000400

	return
}

func encode_setf(sf, op, S, opcode2, sz, Rn, o3, mask uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || opcode2 > 0x3f || sz > 0x1 || Rn > 0x1f || o3 > 0x1 || mask > 0xf:
		err = errOverflow
	case sf == 0x0 && op == 0x0 && S == 0x0:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x1 && opcode2 != 0x0:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x1 && opcode2 == 0x0 && o3 == 0x0 && mask != 0xd:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x1 && opcode2 == 0x0 && o3 == 0x1:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x1 && opcode2 == 0x0 && sz == 0x0 && o3 == 0x0 && mask == 0xd:
		// encoding_SETF8_only_setf
	case sf == 0x0 && op == 0x0 && S == 0x1 && opcode2 == 0x0 && sz == 0x1 && o3 == 0x0 && mask == 0xd:
		// encoding_SETF16_only_setf
	case sf == 0x0 && op == 0x1:
		err = errUnallocated
	case sf == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (opcode2 << 15) | (sz << 14) | (Rn << 5) | (o3 << 4) | mask
	ins |= 0x1a000800

	return
}

func encode_condcmp_reg(sf, op, S, Rm, cond, o2, Rn, o3, nzcv uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || Rm > 0x1f || cond > 0xf || o2 > 0x1 || Rn > 0x1f || o3 > 0x1 || nzcv > 0xf:
		err = errOverflow
	case o3 == 0x1:
		err = errUnallocated
	case o2 == 0x1:
		err = errUnallocated
	case S == 0x0:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMN_32_condcmp_reg
	case sf == 0x0 && op == 0x1 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMP_32_condcmp_reg
	case sf == 0x1 && op == 0x0 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMN_64_condcmp_reg
	case sf == 0x1 && op == 0x1 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMP_64_condcmp_reg
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (Rm << 16) | (cond << 12) | (o2 << 10) | (Rn << 5) | (o3 << 4) | nzcv
	ins |= 0x1a400000

	return
}

func encode_condcmp_imm(sf, op, S, imm5, cond, o2, Rn, o3, nzcv uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || imm5 > 0x1f || cond > 0xf || o2 > 0x1 || Rn > 0x1f || o3 > 0x1 || nzcv > 0xf:
		err = errOverflow
	case o3 == 0x1:
		err = errUnallocated
	case o2 == 0x1:
		err = errUnallocated
	case S == 0x0:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMN_32_condcmp_imm
	case sf == 0x0 && op == 0x1 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMP_32_condcmp_imm
	case sf == 0x1 && op == 0x0 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMN_64_condcmp_imm
	case sf == 0x1 && op == 0x1 && S == 0x1 && o2 == 0x0 && o3 == 0x0:
		// encoding_CCMP_64_condcmp_imm
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (imm5 << 16) | (cond << 12) | (o2 << 10) | (Rn << 5) | (o3 << 4) | nzcv
	ins |= 0x1a400800

	return
}

func encode_condsel(sf, op, S, Rm, cond, op2, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op > 0x1 || S > 0x1 || Rm > 0x1f || cond > 0xf || op2 > 0x3 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (op2 & 0x2) == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case sf == 0x0 && op == 0x0 && S == 0x0 && op2 == 0x0:
		// encoding_CSEL_32_condsel
	case sf == 0x0 && op == 0x0 && S == 0x0 && op2 == 0x1:
		// encoding_CSINC_32_condsel
	case sf == 0x0 && op == 0x1 && S == 0x0 && op2 == 0x0:
		// encoding_CSINV_32_condsel
	case sf == 0x0 && op == 0x1 && S == 0x0 && op2 == 0x1:
		// encoding_CSNEG_32_condsel
	case sf == 0x1 && op == 0x0 && S == 0x0 && op2 == 0x0:
		// encoding_CSEL_64_condsel
	case sf == 0x1 && op == 0x0 && S == 0x0 && op2 == 0x1:
		// encoding_CSINC_64_condsel
	case sf == 0x1 && op == 0x1 && S == 0x0 && op2 == 0x0:
		// encoding_CSINV_64_condsel
	case sf == 0x1 && op == 0x1 && S == 0x0 && op2 == 0x1:
		// encoding_CSNEG_64_condsel
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op << 30) | (S << 29) | (Rm << 16) | (cond << 12) | (op2 << 10) | (Rn << 5) | Rd
	ins |= 0x1a800000

	return
}

func encode_dp_3src(sf, op54, op31, Rm, o0, Ra, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || op54 > 0x3 || op31 > 0x7 || Rm > 0x1f || o0 > 0x1 || Ra > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case op54 == 0x0 && op31 == 0x2 && o0 == 0x1:
		err = errUnallocated
	case op54 == 0x0 && op31 == 0x3:
		err = errUnallocated
	case op54 == 0x0 && op31 == 0x4:
		err = errUnallocated
	case op54 == 0x0 && op31 == 0x6 && o0 == 0x1:
		err = errUnallocated
	case op54 == 0x0 && op31 == 0x7:
		err = errUnallocated
	case op54 == 0x1:
		err = errUnallocated
	case (op54 & 0x2) == 0x2:
		err = errUnallocated
	case sf == 0x0 && op54 == 0x0 && op31 == 0x0 && o0 == 0x0:
		// encoding_MADD_32A_dp_3src
	case sf == 0x0 && op54 == 0x0 && op31 == 0x0 && o0 == 0x1:
		// encoding_MSUB_32A_dp_3src
	case sf == 0x0 && op54 == 0x0 && op31 == 0x1 && o0 == 0x0:
		err = errUnallocated
	case sf == 0x0 && op54 == 0x0 && op31 == 0x1 && o0 == 0x1:
		err = errUnallocated
	case sf == 0x0 && op54 == 0x0 && op31 == 0x2 && o0 == 0x0:
		err = errUnallocated
	case sf == 0x0 && op54 == 0x0 && op31 == 0x5 && o0 == 0x0:
		err = errUnallocated
	case sf == 0x0 && op54 == 0x0 && op31 == 0x5 && o0 == 0x1:
		err = errUnallocated
	case sf == 0x0 && op54 == 0x0 && op31 == 0x6 && o0 == 0x0:
		err = errUnallocated
	case sf == 0x1 && op54 == 0x0 && op31 == 0x0 && o0 == 0x0:
		// encoding_MADD_64A_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x0 && o0 == 0x1:
		// encoding_MSUB_64A_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x1 && o0 == 0x0:
		// encoding_SMADDL_64WA_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x1 && o0 == 0x1:
		// encoding_SMSUBL_64WA_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x2 && o0 == 0x0:
		// encoding_SMULH_64_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x5 && o0 == 0x0:
		// encoding_UMADDL_64WA_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x5 && o0 == 0x1:
		// encoding_UMSUBL_64WA_dp_3src
	case sf == 0x1 && op54 == 0x0 && op31 == 0x6 && o0 == 0x0:
		// encoding_UMULH_64_dp_3src
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (op54 << 29) | (op31 << 21) | (Rm << 16) | (o0 << 15) | (Ra << 10) | (Rn << 5) | Rd
	ins |= 0x1b000000

	return
}

func encode_cryptoaes(size, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x8) == 0x8:
		err = errUnallocated
	case (opcode & 0x1c) == 0x0:
		err = errUnallocated
	case (opcode & 0x10) == 0x10:
		err = errUnallocated
	case (size & 0x1) == 0x1:
		err = errUnallocated
	case size == 0x0 && opcode == 0x4:
		// encoding_AESE_B_cryptoaes
	case size == 0x0 && opcode == 0x5:
		// encoding_AESD_B_cryptoaes
	case size == 0x0 && opcode == 0x6:
		// encoding_AESMC_B_cryptoaes
	case size == 0x0 && opcode == 0x7:
		// encoding_AESIMC_B_cryptoaes
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x4e280800

	return
}

func encode_cryptosha3(size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || Rm > 0x1f || opcode > 0x7 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x7:
		err = errUnallocated
	case (size & 0x1) == 0x1:
		err = errUnallocated
	case size == 0x0 && opcode == 0x0:
		// encoding_SHA1C_QSV_cryptosha3
	case size == 0x0 && opcode == 0x1:
		// encoding_SHA1P_QSV_cryptosha3
	case size == 0x0 && opcode == 0x2:
		// encoding_SHA1M_QSV_cryptosha3
	case size == 0x0 && opcode == 0x3:
		// encoding_SHA1SU0_VVV_cryptosha3
	case size == 0x0 && opcode == 0x4:
		// encoding_SHA256H_QQV_cryptosha3
	case size == 0x0 && opcode == 0x5:
		// encoding_SHA256H2_QQV_cryptosha3
	case size == 0x0 && opcode == 0x6:
		// encoding_SHA256SU1_VVV_cryptosha3
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (Rm << 16) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x5e000000

	return
}

func encode_cryptosha2(size, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case size > 0x3 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x4) == 0x4:
		err = errUnallocated
	case (opcode & 0x8) == 0x8:
		err = errUnallocated
	case (opcode & 0x10) == 0x10:
		err = errUnallocated
	case (size & 0x1) == 0x1:
		err = errUnallocated
	case size == 0x0 && opcode == 0x0:
		// encoding_SHA1H_SS_cryptosha2
	case size == 0x0 && opcode == 0x1:
		// encoding_SHA1SU1_VV_cryptosha2
	case size == 0x0 && opcode == 0x2:
		// encoding_SHA256SU0_VV_cryptosha2
	case size == 0x0 && opcode == 0x3:
		err = errUnallocated
	case (size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (size << 22) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x5e280800

	return
}

func encode_asisdone(op, imm5, imm4, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case op > 0x1 || imm5 > 0x1f || imm4 > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case op == 0x0 && (imm4&0x1) == 0x1:
		err = errUnallocated
	case op == 0x0 && (imm4&0x2) == 0x2:
		err = errUnallocated
	case op == 0x0 && (imm4&0x4) == 0x4:
		err = errUnallocated
	case op == 0x0 && imm4 == 0x0:
		// encoding_DUP_asisdone_only
	case op == 0x0 && (imm4&0x8) == 0x8:
		err = errUnallocated
	case op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (op << 29) | (imm5 << 16) | (imm4 << 11) | (Rn << 5) | Rd
	ins |= 0x5e000400

	return
}

func encode_asisdsamefp16(U, a, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || a > 0x1 || Rm > 0x1f || opcode > 0x7 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x6:
		err = errUnallocated
	case a == 0x1 && opcode == 0x3:
		err = errUnallocated
	case U == 0x0 && a == 0x0 && opcode == 0x3:
		// encoding_FMULX_asisdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x4:
		// encoding_FCMEQ_asisdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x5:
		err = errUnallocated
	case U == 0x0 && a == 0x0 && opcode == 0x7:
		// encoding_FRECPS_asisdsamefp16_only
	case U == 0x0 && a == 0x1 && opcode == 0x4:
		err = errUnallocated
	case U == 0x0 && a == 0x1 && opcode == 0x5:
		err = errUnallocated
	case U == 0x0 && a == 0x1 && opcode == 0x7:
		// encoding_FRSQRTS_asisdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x3:
		err = errUnallocated
	case U == 0x1 && a == 0x0 && opcode == 0x4:
		// encoding_FCMGE_asisdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x5:
		// encoding_FACGE_asisdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x7:
		err = errUnallocated
	case U == 0x1 && a == 0x1 && opcode == 0x2:
		// encoding_FABD_asisdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x4:
		// encoding_FCMGT_asisdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x5:
		// encoding_FACGT_asisdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (a << 23) | (Rm << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0x5e400400

	return
}

func encode_asisdmiscfp16(U, a, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || a > 0x1 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x18) == 0x0:
		err = errUnallocated
	case (opcode & 0x1c) == 0x8:
		err = errUnallocated
	case (opcode & 0x18) == 0x10:
		err = errUnallocated
	case (opcode & 0x1e) == 0x18:
		err = errUnallocated
	case opcode == 0x1e:
		err = errUnallocated
	case a == 0x0 && (opcode&0x1c) == 0xc:
		err = errUnallocated
	case a == 0x0 && opcode == 0x1f:
		err = errUnallocated
	case a == 0x1 && opcode == 0xf:
		err = errUnallocated
	case a == 0x1 && opcode == 0x1c:
		err = errUnallocated
	case U == 0x0 && a == 0x0 && opcode == 0x1a:
		// encoding_FCVTNS_asisdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1b:
		// encoding_FCVTMS_asisdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1c:
		// encoding_FCVTAS_asisdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1d:
		// encoding_SCVTF_asisdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0xc:
		// encoding_FCMGT_asisdmiscfp16_FZ
	case U == 0x0 && a == 0x1 && opcode == 0xd:
		// encoding_FCMEQ_asisdmiscfp16_FZ
	case U == 0x0 && a == 0x1 && opcode == 0xe:
		// encoding_FCMLT_asisdmiscfp16_FZ
	case U == 0x0 && a == 0x1 && opcode == 0x1a:
		// encoding_FCVTPS_asisdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1b:
		// encoding_FCVTZS_asisdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1d:
		// encoding_FRECPE_asisdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1f:
		// encoding_FRECPX_asisdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1a:
		// encoding_FCVTNU_asisdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1b:
		// encoding_FCVTMU_asisdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1c:
		// encoding_FCVTAU_asisdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1d:
		// encoding_UCVTF_asisdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0xc:
		// encoding_FCMGE_asisdmiscfp16_FZ
	case U == 0x1 && a == 0x1 && opcode == 0xd:
		// encoding_FCMLE_asisdmiscfp16_FZ
	case U == 0x1 && a == 0x1 && opcode == 0xe:
		err = errUnallocated
	case U == 0x1 && a == 0x1 && opcode == 0x1a:
		// encoding_FCVTPU_asisdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1b:
		// encoding_FCVTZU_asisdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1d:
		// encoding_FRSQRTE_asisdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (a << 23) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x5e780800

	return
}

func encode_asisdsame2(U, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0xe) == 0x2:
		err = errUnallocated
	case (opcode & 0xc) == 0x4:
		err = errUnallocated
	case (opcode & 0x8) == 0x8:
		err = errUnallocated
	case U == 0x0 && opcode == 0x0:
		err = errUnallocated
	case U == 0x0 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && opcode == 0x0:
		// encoding_SQRDMLAH_asisdsame2_only
	case U == 0x1 && opcode == 0x1:
		// encoding_SQRDMLSH_asisdsame2_only
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (size << 22) | (Rm << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0x5e008400

	return
}

func encode_asisdmisc(U, size, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || size > 0x3 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x1e) == 0x0:
		err = errUnallocated
	case opcode == 0x2:
		err = errUnallocated
	case (opcode & 0x1e) == 0x4:
		err = errUnallocated
	case opcode == 0x6:
		err = errUnallocated
	case opcode == 0xf:
		err = errUnallocated
	case (opcode & 0x1e) == 0x10:
		err = errUnallocated
	case opcode == 0x13:
		err = errUnallocated
	case opcode == 0x15:
		err = errUnallocated
	case opcode == 0x17:
		err = errUnallocated
	case (opcode & 0x1e) == 0x18:
		err = errUnallocated
	case opcode == 0x1e:
		err = errUnallocated
	case (size&0x2) == 0x0 && (opcode&0x1c) == 0xc:
		err = errUnallocated
	case (size&0x2) == 0x0 && opcode == 0x1f:
		err = errUnallocated
	case (size&0x2) == 0x2 && opcode == 0x16:
		err = errUnallocated
	case (size&0x2) == 0x2 && opcode == 0x1c:
		err = errUnallocated
	case U == 0x0 && opcode == 0x3:
		// encoding_SUQADD_asisdmisc_R
	case U == 0x0 && opcode == 0x7:
		// encoding_SQABS_asisdmisc_R
	case U == 0x0 && opcode == 0x8:
		// encoding_CMGT_asisdmisc_Z
	case U == 0x0 && opcode == 0x9:
		// encoding_CMEQ_asisdmisc_Z
	case U == 0x0 && opcode == 0xa:
		// encoding_CMLT_asisdmisc_Z
	case U == 0x0 && opcode == 0xb:
		// encoding_ABS_asisdmisc_R
	case U == 0x0 && opcode == 0x12:
		err = errUnallocated
	case U == 0x0 && opcode == 0x14:
		// encoding_SQXTN_asisdmisc_N
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x16:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1a:
		// encoding_FCVTNS_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FCVTMS_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCVTAS_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1d:
		// encoding_SCVTF_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FCMGT_asisdmisc_FZ
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xd:
		// encoding_FCMEQ_asisdmisc_FZ
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xe:
		// encoding_FCMLT_asisdmisc_FZ
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FCVTPS_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1b:
		// encoding_FCVTZS_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1d:
		// encoding_FRECPE_asisdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1f:
		// encoding_FRECPX_asisdmisc_R
	case U == 0x1 && opcode == 0x3:
		// encoding_USQADD_asisdmisc_R
	case U == 0x1 && opcode == 0x7:
		// encoding_SQNEG_asisdmisc_R
	case U == 0x1 && opcode == 0x8:
		// encoding_CMGE_asisdmisc_Z
	case U == 0x1 && opcode == 0x9:
		// encoding_CMLE_asisdmisc_Z
	case U == 0x1 && opcode == 0xa:
		err = errUnallocated
	case U == 0x1 && opcode == 0xb:
		// encoding_NEG_asisdmisc_R
	case U == 0x1 && opcode == 0x12:
		// encoding_SQXTUN_asisdmisc_N
	case U == 0x1 && opcode == 0x14:
		// encoding_UQXTN_asisdmisc_N
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x16:
		// encoding_FCVTXN_asisdmisc_N
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1a:
		// encoding_FCVTNU_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FCVTMU_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCVTAU_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1d:
		// encoding_UCVTF_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FCMGE_asisdmisc_FZ
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xd:
		// encoding_FCMLE_asisdmisc_FZ
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xe:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FCVTPU_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1b:
		// encoding_FCVTZU_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1d:
		// encoding_FRSQRTE_asisdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (size << 22) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x5e200800

	return
}

func encode_asisdpair(U, size, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || size > 0x3 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x18) == 0x0:
		err = errUnallocated
	case (opcode & 0x1c) == 0x8:
		err = errUnallocated
	case opcode == 0xe:
		err = errUnallocated
	case (opcode & 0x18) == 0x10:
		err = errUnallocated
	case (opcode & 0x1e) == 0x18:
		err = errUnallocated
	case opcode == 0x1a:
		err = errUnallocated
	case (opcode & 0x1c) == 0x1c:
		err = errUnallocated
	case (size&0x2) == 0x2 && opcode == 0xd:
		err = errUnallocated
	case U == 0x0 && opcode == 0x1b:
		// encoding_ADDP_asisdpair_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0xc:
		// encoding_FMAXNMP_asisdpair_only_H
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0xd:
		// encoding_FADDP_asisdpair_only_H
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0xf:
		// encoding_FMAXP_asisdpair_only_H
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FMINNMP_asisdpair_only_H
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xf:
		// encoding_FMINP_asisdpair_only_H
	case U == 0x1 && opcode == 0x1b:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0xc:
		// encoding_FMAXNMP_asisdpair_only_SD
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0xd:
		// encoding_FADDP_asisdpair_only_SD
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0xf:
		// encoding_FMAXP_asisdpair_only_SD
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FMINNMP_asisdpair_only_SD
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xf:
		// encoding_FMINP_asisdpair_only_SD
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (size << 22) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x5e300800

	return
}

func encode_asisddiff(U, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0xc) == 0x0:
		err = errUnallocated
	case (opcode & 0xc) == 0x4:
		err = errUnallocated
	case opcode == 0x8:
		err = errUnallocated
	case opcode == 0xa:
		err = errUnallocated
	case opcode == 0xc:
		err = errUnallocated
	case (opcode & 0xe) == 0xe:
		err = errUnallocated
	case U == 0x0 && opcode == 0x9:
		// encoding_SQDMLAL_asisddiff_only
	case U == 0x0 && opcode == 0xb:
		// encoding_SQDMLSL_asisddiff_only
	case U == 0x0 && opcode == 0xd:
		// encoding_SQDMULL_asisddiff_only
	case U == 0x1 && opcode == 0x9:
		err = errUnallocated
	case U == 0x1 && opcode == 0xb:
		err = errUnallocated
	case U == 0x1 && opcode == 0xd:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (size << 22) | (Rm << 16) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x5e200000

	return
}

func encode_asisdsame(U, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x0:
		err = errUnallocated
	case (opcode & 0x1e) == 0x2:
		err = errUnallocated
	case opcode == 0x4:
		err = errUnallocated
	case (opcode & 0x1c) == 0xc:
		err = errUnallocated
	case (opcode & 0x1e) == 0x12:
		err = errUnallocated
	case (size&0x2) == 0x2 && opcode == 0x1b:
		err = errUnallocated
	case U == 0x0 && opcode == 0x1:
		// encoding_SQADD_asisdsame_only
	case U == 0x0 && opcode == 0x5:
		// encoding_SQSUB_asisdsame_only
	case U == 0x0 && opcode == 0x6:
		// encoding_CMGT_asisdsame_only
	case U == 0x0 && opcode == 0x7:
		// encoding_CMGE_asisdsame_only
	case U == 0x0 && opcode == 0x8:
		// encoding_SSHL_asisdsame_only
	case U == 0x0 && opcode == 0x9:
		// encoding_SQSHL_asisdsame_only
	case U == 0x0 && opcode == 0xa:
		// encoding_SRSHL_asisdsame_only
	case U == 0x0 && opcode == 0xb:
		// encoding_SQRSHL_asisdsame_only
	case U == 0x0 && opcode == 0x10:
		// encoding_ADD_asisdsame_only
	case U == 0x0 && opcode == 0x11:
		// encoding_CMTST_asisdsame_only
	case U == 0x0 && opcode == 0x14:
		err = errUnallocated
	case U == 0x0 && opcode == 0x15:
		err = errUnallocated
	case U == 0x0 && opcode == 0x16:
		// encoding_SQDMULH_asisdsame_only
	case U == 0x0 && opcode == 0x17:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x18:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x19:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1a:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FMULX_asisdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCMEQ_asisdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1d:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1e:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1f:
		// encoding_FRECPS_asisdsame_only
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x18:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x19:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1a:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1c:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1d:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1e:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1f:
		// encoding_FRSQRTS_asisdsame_only
	case U == 0x1 && opcode == 0x1:
		// encoding_UQADD_asisdsame_only
	case U == 0x1 && opcode == 0x5:
		// encoding_UQSUB_asisdsame_only
	case U == 0x1 && opcode == 0x6:
		// encoding_CMHI_asisdsame_only
	case U == 0x1 && opcode == 0x7:
		// encoding_CMHS_asisdsame_only
	case U == 0x1 && opcode == 0x8:
		// encoding_USHL_asisdsame_only
	case U == 0x1 && opcode == 0x9:
		// encoding_UQSHL_asisdsame_only
	case U == 0x1 && opcode == 0xa:
		// encoding_URSHL_asisdsame_only
	case U == 0x1 && opcode == 0xb:
		// encoding_UQRSHL_asisdsame_only
	case U == 0x1 && opcode == 0x10:
		// encoding_SUB_asisdsame_only
	case U == 0x1 && opcode == 0x11:
		// encoding_CMEQ_asisdsame_only
	case U == 0x1 && opcode == 0x14:
		err = errUnallocated
	case U == 0x1 && opcode == 0x15:
		err = errUnallocated
	case U == 0x1 && opcode == 0x16:
		// encoding_SQRDMULH_asisdsame_only
	case U == 0x1 && opcode == 0x17:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x18:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x19:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1a:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1b:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCMGE_asisdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1d:
		// encoding_FACGE_asisdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1e:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1f:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x18:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x19:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FABD_asisdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1c:
		// encoding_FCMGT_asisdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1d:
		// encoding_FACGT_asisdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1e:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (size << 22) | (Rm << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0x5e200400

	return
}

func encode_asisdshf(U, immh, immb, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || immh > 0xf || immb > 0x7 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case immh != 0x0 && opcode == 0x1:
		err = errUnallocated
	case immh != 0x0 && opcode == 0x3:
		err = errUnallocated
	case immh != 0x0 && opcode == 0x5:
		err = errUnallocated
	case immh != 0x0 && opcode == 0x7:
		err = errUnallocated
	case immh != 0x0 && opcode == 0x9:
		err = errUnallocated
	case immh != 0x0 && opcode == 0xb:
		err = errUnallocated
	case immh != 0x0 && opcode == 0xd:
		err = errUnallocated
	case immh != 0x0 && opcode == 0xf:
		err = errUnallocated
	case immh != 0x0 && (opcode&0x1c) == 0x14:
		err = errUnallocated
	case immh != 0x0 && (opcode&0x1c) == 0x18:
		err = errUnallocated
	case immh != 0x0 && opcode == 0x1d:
		err = errUnallocated
	case immh != 0x0 && opcode == 0x1e:
		err = errUnallocated
	case immh == 0x0:
		err = errUnallocated
	case U == 0x0 && immh != 0x0 && opcode == 0x0:
		// encoding_SSHR_asisdshf_R
	case U == 0x0 && immh != 0x0 && opcode == 0x2:
		// encoding_SSRA_asisdshf_R
	case U == 0x0 && immh != 0x0 && opcode == 0x4:
		// encoding_SRSHR_asisdshf_R
	case U == 0x0 && immh != 0x0 && opcode == 0x6:
		// encoding_SRSRA_asisdshf_R
	case U == 0x0 && immh != 0x0 && opcode == 0x8:
		err = errUnallocated
	case U == 0x0 && immh != 0x0 && opcode == 0xa:
		// encoding_SHL_asisdshf_R
	case U == 0x0 && immh != 0x0 && opcode == 0xc:
		err = errUnallocated
	case U == 0x0 && immh != 0x0 && opcode == 0xe:
		// encoding_SQSHL_asisdshf_R
	case U == 0x0 && immh != 0x0 && opcode == 0x10:
		err = errUnallocated
	case U == 0x0 && immh != 0x0 && opcode == 0x11:
		err = errUnallocated
	case U == 0x0 && immh != 0x0 && opcode == 0x12:
		// encoding_SQSHRN_asisdshf_N
	case U == 0x0 && immh != 0x0 && opcode == 0x13:
		// encoding_SQRSHRN_asisdshf_N
	case U == 0x0 && immh != 0x0 && opcode == 0x1c:
		// encoding_SCVTF_asisdshf_C
	case U == 0x0 && immh != 0x0 && opcode == 0x1f:
		// encoding_FCVTZS_asisdshf_C
	case U == 0x1 && immh != 0x0 && opcode == 0x0:
		// encoding_USHR_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0x2:
		// encoding_USRA_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0x4:
		// encoding_URSHR_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0x6:
		// encoding_URSRA_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0x8:
		// encoding_SRI_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0xa:
		// encoding_SLI_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0xc:
		// encoding_SQSHLU_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0xe:
		// encoding_UQSHL_asisdshf_R
	case U == 0x1 && immh != 0x0 && opcode == 0x10:
		// encoding_SQSHRUN_asisdshf_N
	case U == 0x1 && immh != 0x0 && opcode == 0x11:
		// encoding_SQRSHRUN_asisdshf_N
	case U == 0x1 && immh != 0x0 && opcode == 0x12:
		// encoding_UQSHRN_asisdshf_N
	case U == 0x1 && immh != 0x0 && opcode == 0x13:
		// encoding_UQRSHRN_asisdshf_N
	case U == 0x1 && immh != 0x0 && opcode == 0x1c:
		// encoding_UCVTF_asisdshf_C
	case U == 0x1 && immh != 0x0 && opcode == 0x1f:
		// encoding_FCVTZU_asisdshf_C
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (immh << 19) | (immb << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0x5f000400

	return
}

func encode_asisdelem(U, size, L, M, Rm, opcode, H, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case U > 0x1 || size > 0x3 || L > 0x1 || M > 0x1 || Rm > 0xf || opcode > 0xf || H > 0x1 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x0:
		err = errUnallocated
	case opcode == 0x2:
		err = errUnallocated
	case opcode == 0x4:
		err = errUnallocated
	case opcode == 0x6:
		err = errUnallocated
	case opcode == 0x8:
		err = errUnallocated
	case opcode == 0xa:
		err = errUnallocated
	case opcode == 0xe:
		err = errUnallocated
	case size == 0x1 && opcode == 0x1:
		err = errUnallocated
	case size == 0x1 && opcode == 0x5:
		err = errUnallocated
	case size == 0x1 && opcode == 0x9:
		err = errUnallocated
	case U == 0x0 && opcode == 0x3:
		// encoding_SQDMLAL_asisdelem_L
	case U == 0x0 && opcode == 0x7:
		// encoding_SQDMLSL_asisdelem_L
	case U == 0x0 && opcode == 0xb:
		// encoding_SQDMULL_asisdelem_L
	case U == 0x0 && opcode == 0xc:
		// encoding_SQDMULH_asisdelem_R
	case U == 0x0 && opcode == 0xd:
		// encoding_SQRDMULH_asisdelem_R
	case U == 0x0 && opcode == 0xf:
		err = errUnallocated
	case U == 0x0 && size == 0x0 && opcode == 0x1:
		// encoding_FMLA_asisdelem_RH_H
	case U == 0x0 && size == 0x0 && opcode == 0x5:
		// encoding_FMLS_asisdelem_RH_H
	case U == 0x0 && size == 0x0 && opcode == 0x9:
		// encoding_FMUL_asisdelem_RH_H
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1:
		// encoding_FMLA_asisdelem_R_SD
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x5:
		// encoding_FMLS_asisdelem_R_SD
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x9:
		// encoding_FMUL_asisdelem_R_SD
	case U == 0x1 && opcode == 0x3:
		err = errUnallocated
	case U == 0x1 && opcode == 0x7:
		err = errUnallocated
	case U == 0x1 && opcode == 0xb:
		err = errUnallocated
	case U == 0x1 && opcode == 0xc:
		err = errUnallocated
	case U == 0x1 && opcode == 0xd:
		// encoding_SQRDMLAH_asisdelem_R
	case U == 0x1 && opcode == 0xf:
		// encoding_SQRDMLSH_asisdelem_R
	case U == 0x1 && size == 0x0 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x5:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x9:
		// encoding_FMULX_asisdelem_RH_H
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x5:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x9:
		// encoding_FMULX_asisdelem_R_SD
	default:
		err = errUnmatched
	}

	ins |= (U << 29) | (size << 22) | (L << 21) | (M << 20) | (Rm << 16) | (opcode << 12) | (H << 11) | (Rn << 5) | Rd
	ins |= 0x5f000000

	return
}

func encode_asimdtbl(Q, op2, Rm, len, op, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || op2 > 0x3 || Rm > 0x1f || len > 0x3 || op > 0x1 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (op2 & 0x1) == 0x1:
		err = errUnallocated
	case op2 == 0x0 && len == 0x0 && op == 0x0:
		// encoding_TBL_asimdtbl_L1_1
	case op2 == 0x0 && len == 0x0 && op == 0x1:
		// encoding_TBX_asimdtbl_L1_1
	case op2 == 0x0 && len == 0x1 && op == 0x0:
		// encoding_TBL_asimdtbl_L2_2
	case op2 == 0x0 && len == 0x1 && op == 0x1:
		// encoding_TBX_asimdtbl_L2_2
	case op2 == 0x0 && len == 0x2 && op == 0x0:
		// encoding_TBL_asimdtbl_L3_3
	case op2 == 0x0 && len == 0x2 && op == 0x1:
		// encoding_TBX_asimdtbl_L3_3
	case op2 == 0x0 && len == 0x3 && op == 0x0:
		// encoding_TBL_asimdtbl_L4_4
	case op2 == 0x0 && len == 0x3 && op == 0x1:
		// encoding_TBX_asimdtbl_L4_4
	case (op2 & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (op2 << 22) | (Rm << 16) | (len << 13) | (op << 12) | (Rn << 5) | Rd
	ins |= 0xe000000

	return
}

func encode_asimdperm(Q, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0x7 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x0:
		err = errUnallocated
	case opcode == 0x1:
		// encoding_UZP1_asimdperm_only
	case opcode == 0x2:
		// encoding_TRN1_asimdperm_only
	case opcode == 0x3:
		// encoding_ZIP1_asimdperm_only
	case opcode == 0x4:
		err = errUnallocated
	case opcode == 0x5:
		// encoding_UZP2_asimdperm_only
	case opcode == 0x6:
		// encoding_TRN2_asimdperm_only
	case opcode == 0x7:
		// encoding_ZIP2_asimdperm_only
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (size << 22) | (Rm << 16) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0xe000800

	return
}

func encode_asimdext(Q, op2, Rm, imm4, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || op2 > 0x3 || Rm > 0x1f || imm4 > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (op2 & 0x1) == 0x1:
		err = errUnallocated
	case op2 == 0x0:
		// encoding_EXT_asimdext_only
	case (op2 & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (op2 << 22) | (Rm << 16) | (imm4 << 11) | (Rn << 5) | Rd
	ins |= 0x2e000000

	return
}

func encode_asimdins(Q, op, imm5, imm4, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || op > 0x1 || imm5 > 0x1f || imm4 > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (imm5 & 0xf) == 0x0:
		err = errUnallocated
	case op == 0x0 && imm4 == 0x0:
		// encoding_DUP_asimdins_DV_v
	case op == 0x0 && imm4 == 0x1:
		// encoding_DUP_asimdins_DR_r
	case op == 0x0 && imm4 == 0x2:
		err = errUnallocated
	case op == 0x0 && imm4 == 0x4:
		err = errUnallocated
	case op == 0x0 && imm4 == 0x6:
		err = errUnallocated
	case op == 0x0 && (imm4&0x8) == 0x8:
		err = errUnallocated
	case Q == 0x0 && op == 0x0 && imm4 == 0x3:
		err = errUnallocated
	case Q == 0x0 && op == 0x0 && imm4 == 0x5:
		// encoding_SMOV_asimdins_W_w
	case Q == 0x0 && op == 0x0 && imm4 == 0x7:
		// encoding_UMOV_asimdins_W_w
	case Q == 0x0 && op == 0x1:
		err = errUnallocated
	case Q == 0x1 && op == 0x0 && imm4 == 0x3:
		// encoding_INS_asimdins_IR_r
	case Q == 0x1 && op == 0x0 && imm4 == 0x5:
		// encoding_SMOV_asimdins_X_x
	case Q == 0x1 && op == 0x0 && (imm5&0xf) == 0x8 && imm4 == 0x7:
		// encoding_UMOV_asimdins_X_x
	case Q == 0x1 && op == 0x1:
		// encoding_INS_asimdins_IV_v
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (op << 29) | (imm5 << 16) | (imm4 << 11) | (Rn << 5) | Rd
	ins |= 0xe000400

	return
}

func encode_asimdsamefp16(Q, U, a, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || a > 0x1 || Rm > 0x1f || opcode > 0x7 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case U == 0x0 && a == 0x0 && opcode == 0x0:
		// encoding_FMAXNM_asimdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x1:
		// encoding_FMLA_asimdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x2:
		// encoding_FADD_asimdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x3:
		// encoding_FMULX_asimdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x4:
		// encoding_FCMEQ_asimdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x5:
		err = errUnallocated
	case U == 0x0 && a == 0x0 && opcode == 0x6:
		// encoding_FMAX_asimdsamefp16_only
	case U == 0x0 && a == 0x0 && opcode == 0x7:
		// encoding_FRECPS_asimdsamefp16_only
	case U == 0x0 && a == 0x1 && opcode == 0x0:
		// encoding_FMINNM_asimdsamefp16_only
	case U == 0x0 && a == 0x1 && opcode == 0x1:
		// encoding_FMLS_asimdsamefp16_only
	case U == 0x0 && a == 0x1 && opcode == 0x2:
		// encoding_FSUB_asimdsamefp16_only
	case U == 0x0 && a == 0x1 && opcode == 0x3:
		err = errUnallocated
	case U == 0x0 && a == 0x1 && opcode == 0x4:
		err = errUnallocated
	case U == 0x0 && a == 0x1 && opcode == 0x5:
		err = errUnallocated
	case U == 0x0 && a == 0x1 && opcode == 0x6:
		// encoding_FMIN_asimdsamefp16_only
	case U == 0x0 && a == 0x1 && opcode == 0x7:
		// encoding_FRSQRTS_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x0:
		// encoding_FMAXNMP_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && a == 0x0 && opcode == 0x2:
		// encoding_FADDP_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x3:
		// encoding_FMUL_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x4:
		// encoding_FCMGE_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x5:
		// encoding_FACGE_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x6:
		// encoding_FMAXP_asimdsamefp16_only
	case U == 0x1 && a == 0x0 && opcode == 0x7:
		// encoding_FDIV_asimdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x0:
		// encoding_FMINNMP_asimdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && a == 0x1 && opcode == 0x2:
		// encoding_FABD_asimdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x3:
		err = errUnallocated
	case U == 0x1 && a == 0x1 && opcode == 0x4:
		// encoding_FCMGT_asimdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x5:
		// encoding_FACGT_asimdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x6:
		// encoding_FMINP_asimdsamefp16_only
	case U == 0x1 && a == 0x1 && opcode == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (a << 23) | (Rm << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0xe400400

	return
}

func encode_asimdmiscfp16(Q, U, a, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || a > 0x1 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x18) == 0x0:
		err = errUnallocated
	case (opcode & 0x1c) == 0x8:
		err = errUnallocated
	case (opcode & 0x18) == 0x10:
		err = errUnallocated
	case opcode == 0x1e:
		err = errUnallocated
	case a == 0x0 && (opcode&0x1c) == 0xc:
		err = errUnallocated
	case a == 0x0 && opcode == 0x1f:
		err = errUnallocated
	case a == 0x1 && opcode == 0x1c:
		err = errUnallocated
	case U == 0x0 && a == 0x0 && opcode == 0x18:
		// encoding_FRINTN_asimdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x19:
		// encoding_FRINTM_asimdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1a:
		// encoding_FCVTNS_asimdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1b:
		// encoding_FCVTMS_asimdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1c:
		// encoding_FCVTAS_asimdmiscfp16_R
	case U == 0x0 && a == 0x0 && opcode == 0x1d:
		// encoding_SCVTF_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0xc:
		// encoding_FCMGT_asimdmiscfp16_FZ
	case U == 0x0 && a == 0x1 && opcode == 0xd:
		// encoding_FCMEQ_asimdmiscfp16_FZ
	case U == 0x0 && a == 0x1 && opcode == 0xe:
		// encoding_FCMLT_asimdmiscfp16_FZ
	case U == 0x0 && a == 0x1 && opcode == 0xf:
		// encoding_FABS_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x18:
		// encoding_FRINTP_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x19:
		// encoding_FRINTZ_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1a:
		// encoding_FCVTPS_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1b:
		// encoding_FCVTZS_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1d:
		// encoding_FRECPE_asimdmiscfp16_R
	case U == 0x0 && a == 0x1 && opcode == 0x1f:
		err = errUnallocated
	case U == 0x1 && a == 0x0 && opcode == 0x18:
		// encoding_FRINTA_asimdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x19:
		// encoding_FRINTX_asimdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1a:
		// encoding_FCVTNU_asimdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1b:
		// encoding_FCVTMU_asimdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1c:
		// encoding_FCVTAU_asimdmiscfp16_R
	case U == 0x1 && a == 0x0 && opcode == 0x1d:
		// encoding_UCVTF_asimdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0xc:
		// encoding_FCMGE_asimdmiscfp16_FZ
	case U == 0x1 && a == 0x1 && opcode == 0xd:
		// encoding_FCMLE_asimdmiscfp16_FZ
	case U == 0x1 && a == 0x1 && opcode == 0xe:
		err = errUnallocated
	case U == 0x1 && a == 0x1 && opcode == 0xf:
		// encoding_FNEG_asimdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x18:
		err = errUnallocated
	case U == 0x1 && a == 0x1 && opcode == 0x19:
		// encoding_FRINTI_asimdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1a:
		// encoding_FCVTPU_asimdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1b:
		// encoding_FCVTZU_asimdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1d:
		// encoding_FRSQRTE_asimdmiscfp16_R
	case U == 0x1 && a == 0x1 && opcode == 0x1f:
		// encoding_FSQRT_asimdmiscfp16_R
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (a << 23) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0xe780800

	return
}

func encode_asimdsame2(Q, U, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (size&0x2) == 0x0 && opcode == 0x3:
		err = errUnallocated
	case size == 0x3 && opcode == 0x3:
		err = errUnallocated
	case U == 0x0 && opcode == 0x0:
		err = errUnallocated
	case U == 0x0 && opcode == 0x1:
		err = errUnallocated
	case U == 0x0 && opcode == 0x2:
		// encoding_SDOT_asimdsame2_D
	case U == 0x0 && (opcode&0x8) == 0x8:
		err = errUnallocated
	case U == 0x0 && size == 0x2 && opcode == 0x3:
		// encoding_USDOT_asimdsame2_D
	case U == 0x1 && opcode == 0x0:
		// encoding_SQRDMLAH_asimdsame2_only
	case U == 0x1 && opcode == 0x1:
		// encoding_SQRDMLSH_asimdsame2_only
	case U == 0x1 && opcode == 0x2:
		// encoding_UDOT_asimdsame2_D
	case U == 0x1 && (opcode&0xc) == 0x8:
		// encoding_FCMLA_asimdsame2_C
	case U == 0x1 && (opcode&0xd) == 0xc:
		// encoding_FCADD_asimdsame2_C
	case U == 0x1 && size == 0x0 && opcode == 0xd:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0xf:
		err = errUnallocated
	case U == 0x1 && size == 0x1 && opcode == 0xf:
		// encoding_BFDOT_asimdsame2_D
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xd:
		err = errUnallocated
	case U == 0x1 && size == 0x2 && opcode == 0x3:
		err = errUnallocated
	case U == 0x1 && size == 0x2 && opcode == 0xf:
		err = errUnallocated
	case U == 0x1 && size == 0x3 && opcode == 0xf:
		// encoding_BFMLAL_asimdsame2_F_
	case Q == 0x0 && (opcode&0xc) == 0x4:
		err = errUnallocated
	case Q == 0x0 && U == 0x1 && size == 0x1 && opcode == 0xd:
		err = errUnallocated
	case Q == 0x1 && (size&0x2) == 0x0 && (opcode&0xc) == 0x4:
		err = errUnallocated
	case Q == 0x1 && (size&0x2) == 0x2 && (opcode&0xe) == 0x6:
		err = errUnallocated
	case Q == 0x1 && U == 0x0 && size == 0x2 && opcode == 0x4:
		// encoding_SMMLA_asimdsame2_G
	case Q == 0x1 && U == 0x0 && size == 0x2 && opcode == 0x5:
		// encoding_USMMLA_asimdsame2_G
	case Q == 0x1 && U == 0x1 && size == 0x1 && opcode == 0xd:
		// encoding_BFMMLA_asimdsame2_E
	case Q == 0x1 && U == 0x1 && size == 0x2 && opcode == 0x4:
		// encoding_UMMLA_asimdsame2_G
	case Q == 0x1 && U == 0x1 && size == 0x2 && opcode == 0x5:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (size << 22) | (Rm << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0xe008400

	return
}

func encode_asimdmisc(Q, U, size, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || size > 0x3 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x1e) == 0x10:
		err = errUnallocated
	case opcode == 0x15:
		err = errUnallocated
	case (size&0x2) == 0x0 && (opcode&0x1c) == 0xc:
		err = errUnallocated
	case (size&0x2) == 0x2 && opcode == 0x17:
		err = errUnallocated
	case (size&0x2) == 0x2 && opcode == 0x1e:
		err = errUnallocated
	case size == 0x3 && opcode == 0x16:
		err = errUnallocated
	case U == 0x0 && opcode == 0x0:
		// encoding_REV64_asimdmisc_R
	case U == 0x0 && opcode == 0x1:
		// encoding_REV16_asimdmisc_R
	case U == 0x0 && opcode == 0x2:
		// encoding_SADDLP_asimdmisc_P
	case U == 0x0 && opcode == 0x3:
		// encoding_SUQADD_asimdmisc_R
	case U == 0x0 && opcode == 0x4:
		// encoding_CLS_asimdmisc_R
	case U == 0x0 && opcode == 0x5:
		// encoding_CNT_asimdmisc_R
	case U == 0x0 && opcode == 0x6:
		// encoding_SADALP_asimdmisc_P
	case U == 0x0 && opcode == 0x7:
		// encoding_SQABS_asimdmisc_R
	case U == 0x0 && opcode == 0x8:
		// encoding_CMGT_asimdmisc_Z
	case U == 0x0 && opcode == 0x9:
		// encoding_CMEQ_asimdmisc_Z
	case U == 0x0 && opcode == 0xa:
		// encoding_CMLT_asimdmisc_Z
	case U == 0x0 && opcode == 0xb:
		// encoding_ABS_asimdmisc_R
	case U == 0x0 && opcode == 0x12:
		// encoding_XTN_asimdmisc_N
	case U == 0x0 && opcode == 0x13:
		err = errUnallocated
	case U == 0x0 && opcode == 0x14:
		// encoding_SQXTN_asimdmisc_N
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x16:
		// encoding_FCVTN_asimdmisc_N
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x17:
		// encoding_FCVTL_asimdmisc_L
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x18:
		// encoding_FRINTN_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x19:
		// encoding_FRINTM_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1a:
		// encoding_FCVTNS_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FCVTMS_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCVTAS_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1d:
		// encoding_SCVTF_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1e:
		// encoding_FRINT32Z_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1f:
		// encoding_FRINT64Z_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FCMGT_asimdmisc_FZ
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xd:
		// encoding_FCMEQ_asimdmisc_FZ
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xe:
		// encoding_FCMLT_asimdmisc_FZ
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0xf:
		// encoding_FABS_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x18:
		// encoding_FRINTP_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x19:
		// encoding_FRINTZ_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FCVTPS_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1b:
		// encoding_FCVTZS_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1c:
		// encoding_URECPE_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1d:
		// encoding_FRECPE_asimdmisc_R
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1f:
		err = errUnallocated
	case U == 0x0 && size == 0x2 && opcode == 0x16:
		// encoding_BFCVTN_asimdmisc_4S
	case U == 0x1 && opcode == 0x0:
		// encoding_REV32_asimdmisc_R
	case U == 0x1 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && opcode == 0x2:
		// encoding_UADDLP_asimdmisc_P
	case U == 0x1 && opcode == 0x3:
		// encoding_USQADD_asimdmisc_R
	case U == 0x1 && opcode == 0x4:
		// encoding_CLZ_asimdmisc_R
	case U == 0x1 && opcode == 0x6:
		// encoding_UADALP_asimdmisc_P
	case U == 0x1 && opcode == 0x7:
		// encoding_SQNEG_asimdmisc_R
	case U == 0x1 && opcode == 0x8:
		// encoding_CMGE_asimdmisc_Z
	case U == 0x1 && opcode == 0x9:
		// encoding_CMLE_asimdmisc_Z
	case U == 0x1 && opcode == 0xa:
		err = errUnallocated
	case U == 0x1 && opcode == 0xb:
		// encoding_NEG_asimdmisc_R
	case U == 0x1 && opcode == 0x12:
		// encoding_SQXTUN_asimdmisc_N
	case U == 0x1 && opcode == 0x13:
		// encoding_SHLL_asimdmisc_S
	case U == 0x1 && opcode == 0x14:
		// encoding_UQXTN_asimdmisc_N
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x16:
		// encoding_FCVTXN_asimdmisc_N
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x17:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x18:
		// encoding_FRINTA_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x19:
		// encoding_FRINTX_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1a:
		// encoding_FCVTNU_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FCVTMU_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCVTAU_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1d:
		// encoding_UCVTF_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1e:
		// encoding_FRINT32X_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1f:
		// encoding_FRINT64X_asimdmisc_R
	case U == 0x1 && size == 0x0 && opcode == 0x5:
		// encoding_NOT_asimdmisc_R
	case U == 0x1 && size == 0x1 && opcode == 0x5:
		// encoding_RBIT_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x5:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FCMGE_asimdmisc_FZ
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xd:
		// encoding_FCMLE_asimdmisc_FZ
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xe:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xf:
		// encoding_FNEG_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x18:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x19:
		// encoding_FRINTI_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FCVTPU_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1b:
		// encoding_FCVTZU_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1c:
		// encoding_URSQRTE_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1d:
		// encoding_FRSQRTE_asimdmisc_R
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1f:
		// encoding_FSQRT_asimdmisc_R
	case U == 0x1 && size == 0x2 && opcode == 0x16:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (size << 22) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0xe200800

	return
}

func encode_asimdall(Q, U, size, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || size > 0x3 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x1e) == 0x0:
		err = errUnallocated
	case opcode == 0x2:
		err = errUnallocated
	case (opcode & 0x1c) == 0x4:
		err = errUnallocated
	case (opcode & 0x1e) == 0x8:
		err = errUnallocated
	case opcode == 0xb:
		err = errUnallocated
	case opcode == 0xd:
		err = errUnallocated
	case opcode == 0xe:
		err = errUnallocated
	case (opcode & 0x18) == 0x10:
		err = errUnallocated
	case (opcode & 0x1e) == 0x18:
		err = errUnallocated
	case (opcode & 0x1c) == 0x1c:
		err = errUnallocated
	case U == 0x0 && opcode == 0x3:
		// encoding_SADDLV_asimdall_only
	case U == 0x0 && opcode == 0xa:
		// encoding_SMAXV_asimdall_only
	case U == 0x0 && opcode == 0x1a:
		// encoding_SMINV_asimdall_only
	case U == 0x0 && opcode == 0x1b:
		// encoding_ADDV_asimdall_only
	case U == 0x0 && size == 0x0 && opcode == 0xc:
		// encoding_FMAXNMV_asimdall_only_H
	case U == 0x0 && size == 0x0 && opcode == 0xf:
		// encoding_FMAXV_asimdall_only_H
	case U == 0x0 && size == 0x1 && opcode == 0xc:
		err = errUnallocated
	case U == 0x0 && size == 0x1 && opcode == 0xf:
		err = errUnallocated
	case U == 0x0 && size == 0x2 && opcode == 0xc:
		// encoding_FMINNMV_asimdall_only_H
	case U == 0x0 && size == 0x2 && opcode == 0xf:
		// encoding_FMINV_asimdall_only_H
	case U == 0x0 && size == 0x3 && opcode == 0xc:
		err = errUnallocated
	case U == 0x0 && size == 0x3 && opcode == 0xf:
		err = errUnallocated
	case U == 0x1 && opcode == 0x3:
		// encoding_UADDLV_asimdall_only
	case U == 0x1 && opcode == 0xa:
		// encoding_UMAXV_asimdall_only
	case U == 0x1 && opcode == 0x1a:
		// encoding_UMINV_asimdall_only
	case U == 0x1 && opcode == 0x1b:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0xc:
		// encoding_FMAXNMV_asimdall_only_SD
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0xf:
		// encoding_FMAXV_asimdall_only_SD
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xc:
		// encoding_FMINNMV_asimdall_only_SD
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0xf:
		// encoding_FMINV_asimdall_only_SD
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (size << 22) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0xe300800

	return
}

func encode_asimddiff(Q, U, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0xf:
		err = errUnallocated
	case U == 0x0 && opcode == 0x0:
		// encoding_SADDL_asimddiff_L
	case U == 0x0 && opcode == 0x1:
		// encoding_SADDW_asimddiff_W
	case U == 0x0 && opcode == 0x2:
		// encoding_SSUBL_asimddiff_L
	case U == 0x0 && opcode == 0x3:
		// encoding_SSUBW_asimddiff_W
	case U == 0x0 && opcode == 0x4:
		// encoding_ADDHN_asimddiff_N
	case U == 0x0 && opcode == 0x5:
		// encoding_SABAL_asimddiff_L
	case U == 0x0 && opcode == 0x6:
		// encoding_SUBHN_asimddiff_N
	case U == 0x0 && opcode == 0x7:
		// encoding_SABDL_asimddiff_L
	case U == 0x0 && opcode == 0x8:
		// encoding_SMLAL_asimddiff_L
	case U == 0x0 && opcode == 0x9:
		// encoding_SQDMLAL_asimddiff_L
	case U == 0x0 && opcode == 0xa:
		// encoding_SMLSL_asimddiff_L
	case U == 0x0 && opcode == 0xb:
		// encoding_SQDMLSL_asimddiff_L
	case U == 0x0 && opcode == 0xc:
		// encoding_SMULL_asimddiff_L
	case U == 0x0 && opcode == 0xd:
		// encoding_SQDMULL_asimddiff_L
	case U == 0x0 && opcode == 0xe:
		// encoding_PMULL_asimddiff_L
	case U == 0x1 && opcode == 0x0:
		// encoding_UADDL_asimddiff_L
	case U == 0x1 && opcode == 0x1:
		// encoding_UADDW_asimddiff_W
	case U == 0x1 && opcode == 0x2:
		// encoding_USUBL_asimddiff_L
	case U == 0x1 && opcode == 0x3:
		// encoding_USUBW_asimddiff_W
	case U == 0x1 && opcode == 0x4:
		// encoding_RADDHN_asimddiff_N
	case U == 0x1 && opcode == 0x5:
		// encoding_UABAL_asimddiff_L
	case U == 0x1 && opcode == 0x6:
		// encoding_RSUBHN_asimddiff_N
	case U == 0x1 && opcode == 0x7:
		// encoding_UABDL_asimddiff_L
	case U == 0x1 && opcode == 0x8:
		// encoding_UMLAL_asimddiff_L
	case U == 0x1 && opcode == 0x9:
		err = errUnallocated
	case U == 0x1 && opcode == 0xa:
		// encoding_UMLSL_asimddiff_L
	case U == 0x1 && opcode == 0xb:
		err = errUnallocated
	case U == 0x1 && opcode == 0xc:
		// encoding_UMULL_asimddiff_L
	case U == 0x1 && opcode == 0xd:
		err = errUnallocated
	case U == 0x1 && opcode == 0xe:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (size << 22) | (Rm << 16) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0xe200000

	return
}

func encode_asimdsame(Q, U, size, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || size > 0x3 || Rm > 0x1f || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case U == 0x0 && opcode == 0x0:
		// encoding_SHADD_asimdsame_only
	case U == 0x0 && opcode == 0x1:
		// encoding_SQADD_asimdsame_only
	case U == 0x0 && opcode == 0x2:
		// encoding_SRHADD_asimdsame_only
	case U == 0x0 && opcode == 0x4:
		// encoding_SHSUB_asimdsame_only
	case U == 0x0 && opcode == 0x5:
		// encoding_SQSUB_asimdsame_only
	case U == 0x0 && opcode == 0x6:
		// encoding_CMGT_asimdsame_only
	case U == 0x0 && opcode == 0x7:
		// encoding_CMGE_asimdsame_only
	case U == 0x0 && opcode == 0x8:
		// encoding_SSHL_asimdsame_only
	case U == 0x0 && opcode == 0x9:
		// encoding_SQSHL_asimdsame_only
	case U == 0x0 && opcode == 0xa:
		// encoding_SRSHL_asimdsame_only
	case U == 0x0 && opcode == 0xb:
		// encoding_SQRSHL_asimdsame_only
	case U == 0x0 && opcode == 0xc:
		// encoding_SMAX_asimdsame_only
	case U == 0x0 && opcode == 0xd:
		// encoding_SMIN_asimdsame_only
	case U == 0x0 && opcode == 0xe:
		// encoding_SABD_asimdsame_only
	case U == 0x0 && opcode == 0xf:
		// encoding_SABA_asimdsame_only
	case U == 0x0 && opcode == 0x10:
		// encoding_ADD_asimdsame_only
	case U == 0x0 && opcode == 0x11:
		// encoding_CMTST_asimdsame_only
	case U == 0x0 && opcode == 0x12:
		// encoding_MLA_asimdsame_only
	case U == 0x0 && opcode == 0x13:
		// encoding_MUL_asimdsame_only
	case U == 0x0 && opcode == 0x14:
		// encoding_SMAXP_asimdsame_only
	case U == 0x0 && opcode == 0x15:
		// encoding_SMINP_asimdsame_only
	case U == 0x0 && opcode == 0x16:
		// encoding_SQDMULH_asimdsame_only
	case U == 0x0 && opcode == 0x17:
		// encoding_ADDP_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x18:
		// encoding_FMAXNM_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x19:
		// encoding_FMLA_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1a:
		// encoding_FADD_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FMULX_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCMEQ_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1e:
		// encoding_FMAX_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x1f:
		// encoding_FRECPS_asimdsame_only
	case U == 0x0 && size == 0x0 && opcode == 0x3:
		// encoding_AND_asimdsame_only
	case U == 0x0 && size == 0x0 && opcode == 0x1d:
		// encoding_FMLAL_asimdsame_F
	case U == 0x0 && size == 0x1 && opcode == 0x3:
		// encoding_BIC_asimdsame_only
	case U == 0x0 && size == 0x1 && opcode == 0x1d:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x18:
		// encoding_FMINNM_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x19:
		// encoding_FMLS_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FSUB_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1b:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1c:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1e:
		// encoding_FMIN_asimdsame_only
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1f:
		// encoding_FRSQRTS_asimdsame_only
	case U == 0x0 && size == 0x2 && opcode == 0x3:
		// encoding_ORR_asimdsame_only
	case U == 0x0 && size == 0x2 && opcode == 0x1d:
		// encoding_FMLSL_asimdsame_F
	case U == 0x0 && size == 0x3 && opcode == 0x3:
		// encoding_ORN_asimdsame_only
	case U == 0x0 && size == 0x3 && opcode == 0x1d:
		err = errUnallocated
	case U == 0x1 && opcode == 0x0:
		// encoding_UHADD_asimdsame_only
	case U == 0x1 && opcode == 0x1:
		// encoding_UQADD_asimdsame_only
	case U == 0x1 && opcode == 0x2:
		// encoding_URHADD_asimdsame_only
	case U == 0x1 && opcode == 0x4:
		// encoding_UHSUB_asimdsame_only
	case U == 0x1 && opcode == 0x5:
		// encoding_UQSUB_asimdsame_only
	case U == 0x1 && opcode == 0x6:
		// encoding_CMHI_asimdsame_only
	case U == 0x1 && opcode == 0x7:
		// encoding_CMHS_asimdsame_only
	case U == 0x1 && opcode == 0x8:
		// encoding_USHL_asimdsame_only
	case U == 0x1 && opcode == 0x9:
		// encoding_UQSHL_asimdsame_only
	case U == 0x1 && opcode == 0xa:
		// encoding_URSHL_asimdsame_only
	case U == 0x1 && opcode == 0xb:
		// encoding_UQRSHL_asimdsame_only
	case U == 0x1 && opcode == 0xc:
		// encoding_UMAX_asimdsame_only
	case U == 0x1 && opcode == 0xd:
		// encoding_UMIN_asimdsame_only
	case U == 0x1 && opcode == 0xe:
		// encoding_UABD_asimdsame_only
	case U == 0x1 && opcode == 0xf:
		// encoding_UABA_asimdsame_only
	case U == 0x1 && opcode == 0x10:
		// encoding_SUB_asimdsame_only
	case U == 0x1 && opcode == 0x11:
		// encoding_CMEQ_asimdsame_only
	case U == 0x1 && opcode == 0x12:
		// encoding_MLS_asimdsame_only
	case U == 0x1 && opcode == 0x13:
		// encoding_PMUL_asimdsame_only
	case U == 0x1 && opcode == 0x14:
		// encoding_UMAXP_asimdsame_only
	case U == 0x1 && opcode == 0x15:
		// encoding_UMINP_asimdsame_only
	case U == 0x1 && opcode == 0x16:
		// encoding_SQRDMULH_asimdsame_only
	case U == 0x1 && opcode == 0x17:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x18:
		// encoding_FMAXNMP_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1a:
		// encoding_FADDP_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1b:
		// encoding_FMUL_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1c:
		// encoding_FCMGE_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1d:
		// encoding_FACGE_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1e:
		// encoding_FMAXP_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x1f:
		// encoding_FDIV_asimdsame_only
	case U == 0x1 && size == 0x0 && opcode == 0x3:
		// encoding_EOR_asimdsame_only
	case U == 0x1 && size == 0x0 && opcode == 0x19:
		// encoding_FMLAL2_asimdsame_F
	case U == 0x1 && size == 0x1 && opcode == 0x3:
		// encoding_BSL_asimdsame_only
	case U == 0x1 && size == 0x1 && opcode == 0x19:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x18:
		// encoding_FMINNMP_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1a:
		// encoding_FABD_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1b:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1c:
		// encoding_FCMGT_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1d:
		// encoding_FACGT_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1e:
		// encoding_FMINP_asimdsame_only
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x1f:
		err = errUnallocated
	case U == 0x1 && size == 0x2 && opcode == 0x3:
		// encoding_BIT_asimdsame_only
	case U == 0x1 && size == 0x2 && opcode == 0x19:
		// encoding_FMLSL2_asimdsame_F
	case U == 0x1 && size == 0x3 && opcode == 0x3:
		// encoding_BIF_asimdsame_only
	case U == 0x1 && size == 0x3 && opcode == 0x19:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (size << 22) | (Rm << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0xe200400

	return
}

func encode_asimdimm(Q, op, a, b, c, cmode, o2, d, e, f, g, h, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || op > 0x1 || a > 0x1 || b > 0x1 || c > 0x1 || cmode > 0xf || o2 > 0x1 || d > 0x1 || e > 0x1 || f > 0x1 || g > 0x1 || h > 0x1 || Rd > 0x1f:
		err = errOverflow
	case op == 0x0 && (cmode&0x8) == 0x0 && o2 == 0x1:
		err = errUnallocated
	case op == 0x0 && (cmode&0x9) == 0x0 && o2 == 0x0:
		// encoding_MOVI_asimdimm_L_sl
	case op == 0x0 && (cmode&0x9) == 0x1 && o2 == 0x0:
		// encoding_ORR_asimdimm_L_sl
	case op == 0x0 && (cmode&0xc) == 0x8 && o2 == 0x1:
		err = errUnallocated
	case op == 0x0 && (cmode&0xd) == 0x8 && o2 == 0x0:
		// encoding_MOVI_asimdimm_L_hl
	case op == 0x0 && (cmode&0xd) == 0x9 && o2 == 0x0:
		// encoding_ORR_asimdimm_L_hl
	case op == 0x0 && (cmode&0xe) == 0xc && o2 == 0x0:
		// encoding_MOVI_asimdimm_M_sm
	case op == 0x0 && (cmode&0xe) == 0xc && o2 == 0x1:
		err = errUnallocated
	case op == 0x0 && cmode == 0xe && o2 == 0x0:
		// encoding_MOVI_asimdimm_N_b
	case op == 0x0 && cmode == 0xe && o2 == 0x1:
		err = errUnallocated
	case op == 0x0 && cmode == 0xf && o2 == 0x0:
		// encoding_FMOV_asimdimm_S_s
	case op == 0x0 && cmode == 0xf && o2 == 0x1:
		// encoding_FMOV_asimdimm_H_h
	case op == 0x1 && o2 == 0x1:
		err = errUnallocated
	case op == 0x1 && (cmode&0x9) == 0x0 && o2 == 0x0:
		// encoding_MVNI_asimdimm_L_sl
	case op == 0x1 && (cmode&0x9) == 0x1 && o2 == 0x0:
		// encoding_BIC_asimdimm_L_sl
	case op == 0x1 && (cmode&0xd) == 0x8 && o2 == 0x0:
		// encoding_MVNI_asimdimm_L_hl
	case op == 0x1 && (cmode&0xd) == 0x9 && o2 == 0x0:
		// encoding_BIC_asimdimm_L_hl
	case op == 0x1 && (cmode&0xe) == 0xc && o2 == 0x0:
		// encoding_MVNI_asimdimm_M_sm
	case Q == 0x0 && op == 0x1 && cmode == 0xe && o2 == 0x0:
		// encoding_MOVI_asimdimm_D_ds
	case Q == 0x0 && op == 0x1 && cmode == 0xf && o2 == 0x0:
		err = errUnallocated
	case Q == 0x1 && op == 0x1 && cmode == 0xe && o2 == 0x0:
		// encoding_MOVI_asimdimm_D2_d
	case Q == 0x1 && op == 0x1 && cmode == 0xf && o2 == 0x0:
		// encoding_FMOV_asimdimm_D2_d
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (op << 29) | (a << 18) | (b << 17) | (c << 16) | (cmode << 12) | (o2 << 11) | (d << 9) | (e << 8) | (f << 7) | (g << 6) | (h << 5) | Rd
	ins |= 0xf000400

	return
}

func encode_asimdshf(Q, U, immh, immb, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || immh > 0xf || immb > 0x7 || opcode > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x1:
		err = errUnallocated
	case opcode == 0x3:
		err = errUnallocated
	case opcode == 0x5:
		err = errUnallocated
	case opcode == 0x7:
		err = errUnallocated
	case opcode == 0x9:
		err = errUnallocated
	case opcode == 0xb:
		err = errUnallocated
	case opcode == 0xd:
		err = errUnallocated
	case opcode == 0xf:
		err = errUnallocated
	case opcode == 0x15:
		err = errUnallocated
	case (opcode & 0x1e) == 0x16:
		err = errUnallocated
	case (opcode & 0x1c) == 0x18:
		err = errUnallocated
	case opcode == 0x1d:
		err = errUnallocated
	case opcode == 0x1e:
		err = errUnallocated
	case U == 0x0 && opcode == 0x0:
		// encoding_SSHR_asimdshf_R
	case U == 0x0 && opcode == 0x2:
		// encoding_SSRA_asimdshf_R
	case U == 0x0 && opcode == 0x4:
		// encoding_SRSHR_asimdshf_R
	case U == 0x0 && opcode == 0x6:
		// encoding_SRSRA_asimdshf_R
	case U == 0x0 && opcode == 0x8:
		err = errUnallocated
	case U == 0x0 && opcode == 0xa:
		// encoding_SHL_asimdshf_R
	case U == 0x0 && opcode == 0xc:
		err = errUnallocated
	case U == 0x0 && opcode == 0xe:
		// encoding_SQSHL_asimdshf_R
	case U == 0x0 && opcode == 0x10:
		// encoding_SHRN_asimdshf_N
	case U == 0x0 && opcode == 0x11:
		// encoding_RSHRN_asimdshf_N
	case U == 0x0 && opcode == 0x12:
		// encoding_SQSHRN_asimdshf_N
	case U == 0x0 && opcode == 0x13:
		// encoding_SQRSHRN_asimdshf_N
	case U == 0x0 && opcode == 0x14:
		// encoding_SSHLL_asimdshf_L
	case U == 0x0 && opcode == 0x1c:
		// encoding_SCVTF_asimdshf_C
	case U == 0x0 && opcode == 0x1f:
		// encoding_FCVTZS_asimdshf_C
	case U == 0x1 && opcode == 0x0:
		// encoding_USHR_asimdshf_R
	case U == 0x1 && opcode == 0x2:
		// encoding_USRA_asimdshf_R
	case U == 0x1 && opcode == 0x4:
		// encoding_URSHR_asimdshf_R
	case U == 0x1 && opcode == 0x6:
		// encoding_URSRA_asimdshf_R
	case U == 0x1 && opcode == 0x8:
		// encoding_SRI_asimdshf_R
	case U == 0x1 && opcode == 0xa:
		// encoding_SLI_asimdshf_R
	case U == 0x1 && opcode == 0xc:
		// encoding_SQSHLU_asimdshf_R
	case U == 0x1 && opcode == 0xe:
		// encoding_UQSHL_asimdshf_R
	case U == 0x1 && opcode == 0x10:
		// encoding_SQSHRUN_asimdshf_N
	case U == 0x1 && opcode == 0x11:
		// encoding_SQRSHRUN_asimdshf_N
	case U == 0x1 && opcode == 0x12:
		// encoding_UQSHRN_asimdshf_N
	case U == 0x1 && opcode == 0x13:
		// encoding_UQRSHRN_asimdshf_N
	case U == 0x1 && opcode == 0x14:
		// encoding_USHLL_asimdshf_L
	case U == 0x1 && opcode == 0x1c:
		// encoding_UCVTF_asimdshf_C
	case U == 0x1 && opcode == 0x1f:
		// encoding_FCVTZU_asimdshf_C
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (immh << 19) | (immb << 16) | (opcode << 11) | (Rn << 5) | Rd
	ins |= 0xf000400

	return
}

func encode_asimdelem(Q, U, size, L, M, Rm, opcode, H, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Q > 0x1 || U > 0x1 || size > 0x3 || L > 0x1 || M > 0x1 || Rm > 0xf || opcode > 0xf || H > 0x1 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case size == 0x1 && opcode == 0x9:
		err = errUnallocated
	case U == 0x0 && opcode == 0x2:
		// encoding_SMLAL_asimdelem_L
	case U == 0x0 && opcode == 0x3:
		// encoding_SQDMLAL_asimdelem_L
	case U == 0x0 && opcode == 0x6:
		// encoding_SMLSL_asimdelem_L
	case U == 0x0 && opcode == 0x7:
		// encoding_SQDMLSL_asimdelem_L
	case U == 0x0 && opcode == 0x8:
		// encoding_MUL_asimdelem_R
	case U == 0x0 && opcode == 0xa:
		// encoding_SMULL_asimdelem_L
	case U == 0x0 && opcode == 0xb:
		// encoding_SQDMULL_asimdelem_L
	case U == 0x0 && opcode == 0xc:
		// encoding_SQDMULH_asimdelem_R
	case U == 0x0 && opcode == 0xd:
		// encoding_SQRDMULH_asimdelem_R
	case U == 0x0 && opcode == 0xe:
		// encoding_SDOT_asimdelem_D
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x0:
		err = errUnallocated
	case U == 0x0 && (size&0x2) == 0x0 && opcode == 0x4:
		err = errUnallocated
	case U == 0x0 && size == 0x0 && opcode == 0x1:
		// encoding_FMLA_asimdelem_RH_H
	case U == 0x0 && size == 0x0 && opcode == 0x5:
		// encoding_FMLS_asimdelem_RH_H
	case U == 0x0 && size == 0x0 && opcode == 0x9:
		// encoding_FMUL_asimdelem_RH_H
	case U == 0x0 && size == 0x0 && opcode == 0xf:
		// encoding_SUDOT_asimdelem_D
	case U == 0x0 && size == 0x1 && opcode == 0x1:
		err = errUnallocated
	case U == 0x0 && size == 0x1 && opcode == 0x5:
		err = errUnallocated
	case U == 0x0 && size == 0x1 && opcode == 0xf:
		// encoding_BFDOT_asimdelem_E
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x1:
		// encoding_FMLA_asimdelem_R_SD
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x5:
		// encoding_FMLS_asimdelem_R_SD
	case U == 0x0 && (size&0x2) == 0x2 && opcode == 0x9:
		// encoding_FMUL_asimdelem_R_SD
	case U == 0x0 && size == 0x2 && opcode == 0x0:
		// encoding_FMLAL_asimdelem_LH
	case U == 0x0 && size == 0x2 && opcode == 0x4:
		// encoding_FMLSL_asimdelem_LH
	case U == 0x0 && size == 0x2 && opcode == 0xf:
		// encoding_USDOT_asimdelem_D
	case U == 0x0 && size == 0x3 && opcode == 0x0:
		err = errUnallocated
	case U == 0x0 && size == 0x3 && opcode == 0x4:
		err = errUnallocated
	case U == 0x0 && size == 0x3 && opcode == 0xf:
		// encoding_BFMLAL_asimdelem_F
	case U == 0x1 && opcode == 0x0:
		// encoding_MLA_asimdelem_R
	case U == 0x1 && opcode == 0x2:
		// encoding_UMLAL_asimdelem_L
	case U == 0x1 && opcode == 0x4:
		// encoding_MLS_asimdelem_R
	case U == 0x1 && opcode == 0x6:
		// encoding_UMLSL_asimdelem_L
	case U == 0x1 && opcode == 0xa:
		// encoding_UMULL_asimdelem_L
	case U == 0x1 && opcode == 0xb:
		err = errUnallocated
	case U == 0x1 && opcode == 0xd:
		// encoding_SQRDMLAH_asimdelem_R
	case U == 0x1 && opcode == 0xe:
		// encoding_UDOT_asimdelem_D
	case U == 0x1 && opcode == 0xf:
		// encoding_SQRDMLSH_asimdelem_R
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0x8:
		err = errUnallocated
	case U == 0x1 && (size&0x2) == 0x0 && opcode == 0xc:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x3:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x5:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x7:
		err = errUnallocated
	case U == 0x1 && size == 0x0 && opcode == 0x9:
		// encoding_FMULX_asimdelem_RH_H
	case U == 0x1 && size == 0x1 && (opcode&0x9) == 0x1:
		// encoding_FCMLA_asimdelem_C_H
	case U == 0x1 && (size&0x2) == 0x2 && opcode == 0x9:
		// encoding_FMULX_asimdelem_R_SD
	case U == 0x1 && size == 0x2 && (opcode&0x9) == 0x1:
		// encoding_FCMLA_asimdelem_C_S
	case U == 0x1 && size == 0x2 && opcode == 0x8:
		// encoding_FMLAL2_asimdelem_LH
	case U == 0x1 && size == 0x2 && opcode == 0xc:
		// encoding_FMLSL2_asimdelem_LH
	case U == 0x1 && size == 0x3 && opcode == 0x1:
		err = errUnallocated
	case U == 0x1 && size == 0x3 && opcode == 0x3:
		err = errUnallocated
	case U == 0x1 && size == 0x3 && opcode == 0x5:
		err = errUnallocated
	case U == 0x1 && size == 0x3 && opcode == 0x7:
		err = errUnallocated
	case U == 0x1 && size == 0x3 && opcode == 0x8:
		err = errUnallocated
	case U == 0x1 && size == 0x3 && opcode == 0xc:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Q << 30) | (U << 29) | (size << 22) | (L << 21) | (M << 20) | (Rm << 16) | (opcode << 12) | (H << 11) | (Rn << 5) | Rd
	ins |= 0xf000000

	return
}

func encode_crypto3_imm2(Rm, imm2, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Rm > 0x1f || imm2 > 0x3 || opcode > 0x3 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x0:
		// encoding_SM3TT1A_VVV4_crypto3_imm2
	case opcode == 0x1:
		// encoding_SM3TT1B_VVV4_crypto3_imm2
	case opcode == 0x2:
		// encoding_SM3TT2A_VVV4_crypto3_imm2
	case opcode == 0x3:
		// encoding_SM3TT2B_VVV_crypto3_imm2
	default:
		err = errUnmatched
	}

	ins |= (Rm << 16) | (imm2 << 12) | (opcode << 10) | (Rn << 5) | Rd
	ins |= 0xce408000

	return
}

func encode_cryptosha512_3(Rm, O, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Rm > 0x1f || O > 0x1 || opcode > 0x3 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case O == 0x0 && opcode == 0x0:
		// encoding_SHA512H_QQV_cryptosha512_3
	case O == 0x0 && opcode == 0x1:
		// encoding_SHA512H2_QQV_cryptosha512_3
	case O == 0x0 && opcode == 0x2:
		// encoding_SHA512SU1_VVV2_cryptosha512_3
	case O == 0x0 && opcode == 0x3:
		// encoding_RAX1_VVV2_cryptosha512_3
	case O == 0x1 && opcode == 0x0:
		// encoding_SM3PARTW1_VVV4_cryptosha512_3
	case O == 0x1 && opcode == 0x1:
		// encoding_SM3PARTW2_VVV4_cryptosha512_3
	case O == 0x1 && opcode == 0x2:
		// encoding_SM4EKEY_VVV4_cryptosha512_3
	case O == 0x1 && opcode == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Rm << 16) | (O << 14) | (opcode << 10) | (Rn << 5) | Rd
	ins |= 0xce608000

	return
}

func encode_crypto4(Op0, Rm, Ra, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Op0 > 0x3 || Rm > 0x1f || Ra > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case Op0 == 0x0:
		// encoding_EOR3_VVV16_crypto4
	case Op0 == 0x1:
		// encoding_BCAX_VVV16_crypto4
	case Op0 == 0x2:
		// encoding_SM3SS1_VVV4_crypto4
	case Op0 == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (Op0 << 21) | (Rm << 16) | (Ra << 10) | (Rn << 5) | Rd
	ins |= 0xce000000

	return
}

func encode_crypto3_imm6(Rm, imm6, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case Rm > 0x1f || imm6 > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	default:
		// encoding_XAR_VVV2_crypto3_imm6
	}

	ins |= (Rm << 16) | (imm6 << 10) | (Rn << 5) | Rd
	ins |= 0xce800000

	return
}

func encode_cryptosha512_2(opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case opcode > 0x3 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case opcode == 0x0:
		// encoding_SHA512SU0_VV2_cryptosha512_2
	case opcode == 0x1:
		// encoding_SM4E_VV4_cryptosha512_2
	case (opcode & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (opcode << 10) | (Rn << 5) | Rd
	ins |= 0xcec08000

	return
}

func encode_float2fix(sf, S, ptype, rmode, opcode, scale, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || S > 0x1 || ptype > 0x3 || rmode > 0x3 || opcode > 0x7 || scale > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x4) == 0x4:
		err = errUnallocated
	case (rmode&0x1) == 0x0 && (opcode&0x6) == 0x0:
		err = errUnallocated
	case (rmode&0x1) == 0x1 && (opcode&0x6) == 0x2:
		err = errUnallocated
	case (rmode&0x2) == 0x0 && (opcode&0x6) == 0x0:
		err = errUnallocated
	case (rmode&0x2) == 0x2 && (opcode&0x6) == 0x2:
		err = errUnallocated
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case sf == 0x0 && (scale&0x20) == 0x0:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_S32_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_S32_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_32S_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_32S_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_D32_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_D32_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_32D_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_32D_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_H32_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_H32_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_32H_float2fix
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_32H_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_S64_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_S64_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_64S_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_64S_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_D64_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_D64_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_64D_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_64D_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_H64_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_H64_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_64H_float2fix
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_64H_float2fix
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (S << 29) | (ptype << 22) | (rmode << 19) | (opcode << 16) | (scale << 10) | (Rn << 5) | Rd
	ins |= 0x1e000000

	return
}

func encode_float2int(sf, S, ptype, rmode, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case sf > 0x1 || S > 0x1 || ptype > 0x3 || rmode > 0x3 || opcode > 0x7 || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (rmode&0x1) == 0x1 && (opcode&0x6) == 0x2:
		err = errUnallocated
	case (rmode&0x1) == 0x1 && (opcode&0x6) == 0x4:
		err = errUnallocated
	case (rmode&0x2) == 0x2 && (opcode&0x6) == 0x2:
		err = errUnallocated
	case (rmode&0x2) == 0x2 && (opcode&0x6) == 0x4:
		err = errUnallocated
	case S == 0x0 && ptype == 0x2 && (opcode&0x4) == 0x0:
		err = errUnallocated
	case S == 0x0 && ptype == 0x2 && (opcode&0x6) == 0x4:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && (rmode&0x1) == 0x1 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x0:
		// encoding_FCVTNS_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x1:
		// encoding_FCVTNU_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_S32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_S32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x4:
		// encoding_FCVTAS_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x5:
		// encoding_FCVTAU_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x6:
		// encoding_FMOV_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x7:
		// encoding_FMOV_S32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x1 && opcode == 0x0:
		// encoding_FCVTPS_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x1 && opcode == 0x1:
		// encoding_FCVTPU_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && (rmode&0x2) == 0x2 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x2 && opcode == 0x0:
		// encoding_FCVTMS_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x2 && opcode == 0x1:
		// encoding_FCVTMU_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_32S_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && (rmode&0x2) == 0x0 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x0:
		// encoding_FCVTNS_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x1:
		// encoding_FCVTNU_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_D32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_D32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x4:
		// encoding_FCVTAS_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x5:
		// encoding_FCVTAU_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x1 && opcode == 0x0:
		// encoding_FCVTPS_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x1 && opcode == 0x1:
		// encoding_FCVTPU_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x2 && opcode == 0x0:
		// encoding_FCVTMS_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x2 && opcode == 0x1:
		// encoding_FCVTMU_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x2 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x6:
		// encoding_FJCVTZS_32D_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x7:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x2 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x0:
		// encoding_FCVTNS_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x1:
		// encoding_FCVTNU_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_H32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_H32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x4:
		// encoding_FCVTAS_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x5:
		// encoding_FCVTAU_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x6:
		// encoding_FMOV_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x7:
		// encoding_FMOV_H32_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x1 && opcode == 0x0:
		// encoding_FCVTPS_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x1 && opcode == 0x1:
		// encoding_FCVTPU_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x2 && opcode == 0x0:
		// encoding_FCVTMS_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x2 && opcode == 0x1:
		// encoding_FCVTMU_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_32H_float2int
	case sf == 0x0 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_32H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x0:
		// encoding_FCVTNS_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x1:
		// encoding_FCVTNU_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_S64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_S64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x4:
		// encoding_FCVTAS_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x0 && opcode == 0x5:
		// encoding_FCVTAU_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x1 && opcode == 0x0:
		// encoding_FCVTPS_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x1 && opcode == 0x1:
		// encoding_FCVTPU_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x2 && opcode == 0x0:
		// encoding_FCVTMS_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x2 && opcode == 0x1:
		// encoding_FCVTMU_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x0 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_64S_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && (rmode&0x1) == 0x1 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x0:
		// encoding_FCVTNS_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x1:
		// encoding_FCVTNU_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_D64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_D64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x4:
		// encoding_FCVTAS_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x5:
		// encoding_FCVTAU_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x6:
		// encoding_FMOV_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x0 && opcode == 0x7:
		// encoding_FMOV_D64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x1 && opcode == 0x0:
		// encoding_FCVTPS_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x1 && opcode == 0x1:
		// encoding_FCVTPU_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && (rmode&0x2) == 0x2 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x2 && opcode == 0x0:
		// encoding_FCVTMS_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x2 && opcode == 0x1:
		// encoding_FCVTMU_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x1 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_64D_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x2 && (rmode&0x1) == 0x0 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && ptype == 0x2 && rmode == 0x1 && opcode == 0x6:
		// encoding_FMOV_64VX_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x2 && rmode == 0x1 && opcode == 0x7:
		// encoding_FMOV_V64I_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x2 && (rmode&0x2) == 0x2 && (opcode&0x6) == 0x6:
		err = errUnallocated
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x0:
		// encoding_FCVTNS_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x1:
		// encoding_FCVTNU_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x2:
		// encoding_SCVTF_H64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x3:
		// encoding_UCVTF_H64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x4:
		// encoding_FCVTAS_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x5:
		// encoding_FCVTAU_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x6:
		// encoding_FMOV_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x0 && opcode == 0x7:
		// encoding_FMOV_H64_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x1 && opcode == 0x0:
		// encoding_FCVTPS_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x1 && opcode == 0x1:
		// encoding_FCVTPU_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x2 && opcode == 0x0:
		// encoding_FCVTMS_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x2 && opcode == 0x1:
		// encoding_FCVTMU_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x0:
		// encoding_FCVTZS_64H_float2int
	case sf == 0x1 && S == 0x0 && ptype == 0x3 && rmode == 0x3 && opcode == 0x1:
		// encoding_FCVTZU_64H_float2int
	default:
		err = errUnmatched
	}

	ins |= (sf << 31) | (S << 29) | (ptype << 22) | (rmode << 19) | (opcode << 16) | (Rn << 5) | Rd
	ins |= 0x1e200000

	return
}

func encode_floatdp1(M, S, ptype, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || opcode > 0x3f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x20) == 0x20:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x0:
		// encoding_FMOV_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x1:
		// encoding_FABS_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x2:
		// encoding_FNEG_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x3:
		// encoding_FSQRT_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x4:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x5:
		// encoding_FCVT_DS_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x6:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x7:
		// encoding_FCVT_HS_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x8:
		// encoding_FRINTN_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x9:
		// encoding_FRINTP_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0xa:
		// encoding_FRINTM_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0xb:
		// encoding_FRINTZ_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0xc:
		// encoding_FRINTA_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0xd:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0xe:
		// encoding_FRINTX_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0xf:
		// encoding_FRINTI_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x10:
		// encoding_FRINT32Z_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x11:
		// encoding_FRINT32X_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x12:
		// encoding_FRINT64Z_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x13:
		// encoding_FRINT64X_S_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x0 && (opcode&0x3c) == 0x14:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && (opcode&0x38) == 0x18:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x0:
		// encoding_FMOV_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x1:
		// encoding_FABS_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x2:
		// encoding_FNEG_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x3:
		// encoding_FSQRT_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x4:
		// encoding_FCVT_SD_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x5:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x6:
		// encoding_BFCVT_BS_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x7:
		// encoding_FCVT_HD_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x8:
		// encoding_FRINTN_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x9:
		// encoding_FRINTP_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0xa:
		// encoding_FRINTM_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0xb:
		// encoding_FRINTZ_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0xc:
		// encoding_FRINTA_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0xd:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0xe:
		// encoding_FRINTX_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0xf:
		// encoding_FRINTI_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x10:
		// encoding_FRINT32Z_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x11:
		// encoding_FRINT32X_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x12:
		// encoding_FRINT64Z_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x13:
		// encoding_FRINT64X_D_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x1 && (opcode&0x3c) == 0x14:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x1 && (opcode&0x38) == 0x18:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x2 && (opcode&0x20) == 0x0:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x0:
		// encoding_FMOV_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x1:
		// encoding_FABS_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x2:
		// encoding_FNEG_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x3:
		// encoding_FSQRT_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x4:
		// encoding_FCVT_SH_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x5:
		// encoding_FCVT_DH_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && (opcode&0x3e) == 0x6:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x8:
		// encoding_FRINTN_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x9:
		// encoding_FRINTP_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0xa:
		// encoding_FRINTM_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0xb:
		// encoding_FRINTZ_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0xc:
		// encoding_FRINTA_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0xd:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0xe:
		// encoding_FRINTX_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0xf:
		// encoding_FRINTI_H_floatdp1
	case M == 0x0 && S == 0x0 && ptype == 0x3 && (opcode&0x30) == 0x10:
		err = errUnallocated
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (opcode << 15) | (Rn << 5) | Rd
	ins |= 0x1e204000

	return
}

func encode_floatcmp(M, S, ptype, Rm, op, Rn, opcode2 uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || Rm > 0x1f || op > 0x3 || Rn > 0x1f || opcode2 > 0x1f:
		err = errOverflow
	case (opcode2 & 0x1) == 0x1:
		err = errUnallocated
	case (opcode2 & 0x2) == 0x2:
		err = errUnallocated
	case (opcode2 & 0x4) == 0x4:
		err = errUnallocated
	case (op & 0x1) == 0x1:
		err = errUnallocated
	case (op & 0x2) == 0x2:
		err = errUnallocated
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && op == 0x0 && opcode2 == 0x0:
		// encoding_FCMP_S_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x0 && op == 0x0 && opcode2 == 0x8:
		// encoding_FCMP_SZ_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x0 && op == 0x0 && opcode2 == 0x10:
		// encoding_FCMPE_S_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x0 && op == 0x0 && opcode2 == 0x18:
		// encoding_FCMPE_SZ_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x1 && op == 0x0 && opcode2 == 0x0:
		// encoding_FCMP_D_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x1 && op == 0x0 && opcode2 == 0x8:
		// encoding_FCMP_DZ_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x1 && op == 0x0 && opcode2 == 0x10:
		// encoding_FCMPE_D_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x1 && op == 0x0 && opcode2 == 0x18:
		// encoding_FCMPE_DZ_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x3 && op == 0x0 && opcode2 == 0x0:
		// encoding_FCMP_H_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x3 && op == 0x0 && opcode2 == 0x8:
		// encoding_FCMP_HZ_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x3 && op == 0x0 && opcode2 == 0x10:
		// encoding_FCMPE_H_floatcmp
	case M == 0x0 && S == 0x0 && ptype == 0x3 && op == 0x0 && opcode2 == 0x18:
		// encoding_FCMPE_HZ_floatcmp
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (Rm << 16) | (op << 14) | (Rn << 5) | opcode2
	ins |= 0x1e202000

	return
}

func encode_floatimm(M, S, ptype, imm8, imm5, Rd uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || imm8 > 0xff || imm5 > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (imm5 & 0x1) == 0x1:
		err = errUnallocated
	case (imm5 & 0x2) == 0x2:
		err = errUnallocated
	case (imm5 & 0x4) == 0x4:
		err = errUnallocated
	case (imm5 & 0x8) == 0x8:
		err = errUnallocated
	case (imm5 & 0x10) == 0x10:
		err = errUnallocated
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && imm5 == 0x0:
		// encoding_FMOV_S_floatimm
	case M == 0x0 && S == 0x0 && ptype == 0x1 && imm5 == 0x0:
		// encoding_FMOV_D_floatimm
	case M == 0x0 && S == 0x0 && ptype == 0x3 && imm5 == 0x0:
		// encoding_FMOV_H_floatimm
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (imm8 << 13) | (imm5 << 5) | Rd
	ins |= 0x1e201000

	return
}

func encode_floatccmp(M, S, ptype, Rm, cond, Rn, op, nzcv uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || Rm > 0x1f || cond > 0xf || Rn > 0x1f || op > 0x1 || nzcv > 0xf:
		err = errOverflow
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && op == 0x0:
		// encoding_FCCMP_S_floatccmp
	case M == 0x0 && S == 0x0 && ptype == 0x0 && op == 0x1:
		// encoding_FCCMPE_S_floatccmp
	case M == 0x0 && S == 0x0 && ptype == 0x1 && op == 0x0:
		// encoding_FCCMP_D_floatccmp
	case M == 0x0 && S == 0x0 && ptype == 0x1 && op == 0x1:
		// encoding_FCCMPE_D_floatccmp
	case M == 0x0 && S == 0x0 && ptype == 0x3 && op == 0x0:
		// encoding_FCCMP_H_floatccmp
	case M == 0x0 && S == 0x0 && ptype == 0x3 && op == 0x1:
		// encoding_FCCMPE_H_floatccmp
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (Rm << 16) | (cond << 12) | (Rn << 5) | (op << 4) | nzcv
	ins |= 0x1e200400

	return
}

func encode_floatdp2(M, S, ptype, Rm, opcode, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || Rm > 0x1f || opcode > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case (opcode & 0x9) == 0x9:
		err = errUnallocated
	case (opcode & 0xa) == 0xa:
		err = errUnallocated
	case (opcode & 0xc) == 0xc:
		err = errUnallocated
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x0:
		// encoding_FMUL_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x1:
		// encoding_FDIV_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x2:
		// encoding_FADD_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x3:
		// encoding_FSUB_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x4:
		// encoding_FMAX_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x5:
		// encoding_FMIN_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x6:
		// encoding_FMAXNM_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x7:
		// encoding_FMINNM_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x0 && opcode == 0x8:
		// encoding_FNMUL_S_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x0:
		// encoding_FMUL_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x1:
		// encoding_FDIV_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x2:
		// encoding_FADD_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x3:
		// encoding_FSUB_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x4:
		// encoding_FMAX_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x5:
		// encoding_FMIN_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x6:
		// encoding_FMAXNM_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x7:
		// encoding_FMINNM_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x1 && opcode == 0x8:
		// encoding_FNMUL_D_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x0:
		// encoding_FMUL_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x1:
		// encoding_FDIV_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x2:
		// encoding_FADD_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x3:
		// encoding_FSUB_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x4:
		// encoding_FMAX_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x5:
		// encoding_FMIN_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x6:
		// encoding_FMAXNM_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x7:
		// encoding_FMINNM_H_floatdp2
	case M == 0x0 && S == 0x0 && ptype == 0x3 && opcode == 0x8:
		// encoding_FNMUL_H_floatdp2
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (Rm << 16) | (opcode << 12) | (Rn << 5) | Rd
	ins |= 0x1e200800

	return
}

func encode_floatsel(M, S, ptype, Rm, cond, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || Rm > 0x1f || cond > 0xf || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0:
		// encoding_FCSEL_S_floatsel
	case M == 0x0 && S == 0x0 && ptype == 0x1:
		// encoding_FCSEL_D_floatsel
	case M == 0x0 && S == 0x0 && ptype == 0x3:
		// encoding_FCSEL_H_floatsel
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (Rm << 16) | (cond << 12) | (Rn << 5) | Rd
	ins |= 0x1e200c00

	return
}

func encode_floatdp3(M, S, ptype, o1, Rm, o0, Ra, Rn, Rd uint32) (ins uint32, err error) {
	switch {
	case M > 0x1 || S > 0x1 || ptype > 0x3 || o1 > 0x1 || Rm > 0x1f || o0 > 0x1 || Ra > 0x1f || Rn > 0x1f || Rd > 0x1f:
		err = errOverflow
	case ptype == 0x2:
		err = errUnallocated
	case S == 0x1:
		err = errUnallocated
	case M == 0x0 && S == 0x0 && ptype == 0x0 && o1 == 0x0 && o0 == 0x0:
		// encoding_FMADD_S_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x0 && o1 == 0x0 && o0 == 0x1:
		// encoding_FMSUB_S_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x0 && o1 == 0x1 && o0 == 0x0:
		// encoding_FNMADD_S_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x0 && o1 == 0x1 && o0 == 0x1:
		// encoding_FNMSUB_S_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x1 && o1 == 0x0 && o0 == 0x0:
		// encoding_FMADD_D_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x1 && o1 == 0x0 && o0 == 0x1:
		// encoding_FMSUB_D_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x1 && o1 == 0x1 && o0 == 0x0:
		// encoding_FNMADD_D_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x1 && o1 == 0x1 && o0 == 0x1:
		// encoding_FNMSUB_D_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x3 && o1 == 0x0 && o0 == 0x0:
		// encoding_FMADD_H_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x3 && o1 == 0x0 && o0 == 0x1:
		// encoding_FMSUB_H_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x3 && o1 == 0x1 && o0 == 0x0:
		// encoding_FNMADD_H_floatdp3
	case M == 0x0 && S == 0x0 && ptype == 0x3 && o1 == 0x1 && o0 == 0x1:
		// encoding_FNMSUB_H_floatdp3
	case M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}

	ins |= (M << 31) | (S << 29) | (ptype << 22) | (o1 << 21) | (Rm << 16) | (o0 << 15) | (Ra << 10) | (Rn << 5) | Rd
	ins |= 0x1f000000

	return
}
