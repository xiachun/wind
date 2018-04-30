package wind

// Freq used sector id in wind
const (
	C_Sec_AllAShare        = "a001010100000000" // 全部A股
	C_Sec_MainBoard        = "1000016326000000" // 全部主板（不含中小板）
	C_Sec_SmallMediumBoard = "1000009396000000" // 中小板
	C_Sec_GrowthBoard      = "a001010r00000000" // 创业板
	C_Sec_HS300            = "1000000090000000" // 沪深300
	C_Sec_CSI500           = "1000008491000000" // 中证500
	C_Sec_SSE50            = "1000000087000000" // 上证50
	C_Sec_NonFin           = "1000011888000000" // 全部非金融A股
	C_Sec_HShare           = "a002010500000000" // H股
	C_Sec_AHShare          = "a002010600000000" // 含A股的H股
	C_Sec_HKStocksThrough  = "1000025142000000" // 全部港股通
)

// AccCodeToWindField 财务报表科目代码与万得接口查询代码的对照字典
var AccCodeToWindField = map[string]string{
	/*****************************************************************************
	 * 资产
	 *****************************************************************************/
	"A0":  "MONETARY_CAP",                   // 货币资金
	"A1":  "SETTLE_RSRV",                    // 结算备付金
	"A2":  "LOANS_TO_OTH_BANKS",             // 拆出资金
	"A3":  "TRADABLE_FIN_ASSETS",            // 交易性金融资产
	"A4":  "NOTES_RCV",                      // 应收票据
	"A5":  "ACCT_RCV",                       // 应收账款
	"A6":  "PREPAY",                         // 预付款项
	"A7":  "PREM_RCV",                       // 应收保费
	"A8":  "RCV_FROM_REINSURER",             // 应收分保账款
	"A9":  "RCV_FROM_CEDED_INSUR_CONT_RSRV", // 应收分保合同准备金
	"A10": "INT_RCV",                        // 应收利息
	"A11": "DVD_RCV",                        // 应收股利
	"A12": "OTH_RCV",                        // 其它应收款
	"A13": "RED_MONETARY_CAP_FOR_SALE",      // 买入返售金融资产
	"A14": "INVENTORIES",                    // 存货
	"A15": "NON_CUR_ASSETS_DUE_WITHIN_1Y",   // 一年内到期的非流动资产
	"A16": "OTH_CUR_ASSETS",                 // 其它流动资产
	"A17": "TOT_CUR_ASSETS",                 // 流动资产合计
	"A18": "LOANS_AND_ADV_GRANTED",          // 发放贷款及垫款
	"A19": "FIN_ASSETS_AVAIL_FOR_SALE",      // 可供出售金融资产
	"A20": "HELD_TO_MTY_INVEST",             // 持有至到期投资
	"A21": "LONG_TERM_REC",                  // 长期应收款
	"A22": "LONG_TERM_EQY_INVEST",           // 长期股权投资
	"A23": "INVEST_REAL_ESTATE",             // 投资性房地产
	"A24": "FIX_ASSETS",                     // 固定资产原价
	"A25": "CONST_IN_PROG",                  // 在建工程
	"A26": "PROJ_MATL",                      // 工程物资
	"A27": "FIX_ASSETS_DISP",                // 固定资产清理
	"A28": "PRODUCTIVE_BIO_ASSETS",          // 生产性生物资产
	"A29": "OIL_AND_NATURAL_GAS_ASSETS",     // 油气资产
	"A30": "INTANG_ASSETS",                  // 无形资产
	"A31": "R_AND_D_COSTS",                  // 开发支出
	"A32": "GOODWILL",                       // 商誉
	"A33": "LONG_TERM_DEFERRED_EXP",         // 长期待摊费用
	"A34": "DEFERRED_TAX_ASSETS",            // 递延所得税资产
	"A35": "OTH_NON_CUR_ASSETS",             // 其他非流动资产
	"A36": "TOT_NON_CUR_ASSETS",             // 非流动资产合计
	"A37": "TOT_ASSETS",                     // 资产总计
	/*****************************************************************************
	 * 负债
	 *****************************************************************************/
	"L0":  "ST_BORROW",                     // 短期借款
	"L1":  "BORROW_CENTRAL_BANK",           // 向中央银行借款
	"L2":  "DEPOSIT_RECEIVED_IB_DEPOSITS",  // 吸收存款及同业存放
	"L3":  "LOANS_OTH_BANKS",               // 拆入资金
	"L4":  "TRADABLE_FIN_LIAB",             // 交易性金融负债
	"L5":  "NOTES_PAYABLE",                 // 应付票据
	"L6":  "ACCT_PAYABLE",                  // 应付账款
	"L7":  "ADV_FROM_CUST",                 // 预收款项
	"L8":  "FUND_SALES_FIN_ASSETS_RP",      // 卖出回购金融资产款
	"L9":  "HANDLING_CHARGES_COMM_PAYABLE", // 应付手续费及佣金
	"L10": "EMPL_BEN_PAYABLE",              // 应付职工薪酬
	"L11": "TAXES_SURCHARGES_PAYABLE",      // 应交税费
	"L12": "INT_PAYABLE",                   // 应付利息
	"L13": "DVD_PAYABLE",                   // 应付股利
	"L14": "OTH_PAYABLE",                   // 其他应付款
	"L15": "PAYABLE_TO_REINSURER",          // 应付分保账款
	"L16": "RSRV_INSUR_CONT",               // 保险合同准备金
	"L17": "ACTING_TRADING_SEC",            // 代理买卖证券款
	"L18": "ACTING_UW_SEC",                 // 代理承销证券款
	"L19": "NON_CUR_LIAB_DUE_WITHIN_1Y",    // 一年内到期的非流动负债
	"L20": "OTH_CUR_LIAB",                  // 其他流动负债
	"L21": "TOT_CUR_LIAB",                  // 流动负债合计
	"L22": "LT_BORROW",                     // 长期借款
	"L23": "BONDS_PAYABLE",                 // 应付债券
	"L24": "LT_PAYABLE",                    // 长期应付款
	"L25": "SPECIFIC_ITEM_PAYABLE",         // 专项应付款
	"L26": "PROVISIONS",                    // 预计负债
	"L27": "DEFERRED_TAX_LIAB",             // 递延所得税负债
	"L28": "OTH_NON_CUR_LIAB",              // 其他非流动负债
	"L29": "TOT_NON_CUR_LIAB",              // 非流动负债合计
	"L30": "TOT_LIAB",                      // 负债合计
	/*****************************************************************************
	 * 权益
	 *****************************************************************************/
	"E0":  "CAP_STK",                     // 实收资本（或股本）
	"E1":  "CAP_RSRV",                    // 资本公积
	"E2":  "TSY_STK",                     // 减：库存股
	"E3":  "SURPLUS_RSRV",                // 盈余公积
	"E4":  "SPECIAL_RSRV",                // 专项储备
	"E5":  "PROV_NOM_RISKS",              // 一般风险准备
	"E6":  "UNDISTRIBUTED_PROFIT",        // 未分配利润
	"E7":  "CNVD_DIFF_FOREIGN_CURR_STAT", // 外币报表折算差额
	"E8":  "EQY_BELONGTO_PARCOMSH",       // 归属母公司所有者权益合计
	"E9":  "MINORITY_INT",                // 少数股东权益
	"E10": "TOT_EQUITY",                  // 所有者权益合计
	"E11": "TOT_LIAB_SHRHLDR_EQY",        // 负债和所有者权益总计
	/*****************************************************************************
	 * 利润表
	 *****************************************************************************/
	"P0":  "TOT_OPER_REV",                // 营业总收入
	"P1":  "OPER_REV",                    // 营业收入
	"P2":  "INT_INC",                     // 利息收入
	"P3":  "INSUR_PREM_UNEARNED",         // 已赚保费
	"P4":  "HANDLING_CHRG_COMM_INC",      // 手续费及佣金收入
	"P5":  "TOT_OPER_COST",               // 营业总成本
	"P6":  "OPER_COST",                   // 营业成本
	"P7":  "INT_EXP",                     // 利息支出
	"P8":  "HANDLING_CHRG_COMM_EXP",      // 手续费及佣金支出
	"P9":  "PREPAY_SURR",                 // 退保金
	"P10": "NET_CLAIM_EXP",               // 赔付支出净额
	"P11": "NET_INSUR_CONT_RSRV",         // 提取保险合同准备金净额
	"P12": "DVD_EXP_INSURED",             // 保单红利支出
	"P13": "REINSURANCE_EXP",             // 分保费用
	"P14": "TAXES_SURCHARGES_OPS",        // 营业税金及附加
	"P15": "SELLING_DIST_EXP",            // 销售费用
	"P16": "GERL_ADMIN_EXP",              // 管理费用
	"P17": "FIN_EXP_IS",                  // 财务费用
	"P18": "IMPAIR_LOSS_ASSETS",          // 资产减值损失
	"P19": "NET_GAIN_CHG_FV",             // 公允价值变动净收益
	"P20": "NET_INVEST_INC",              // 投资净收益
	"P21": "INC_INVEST_ASSOC_JV_ENTP",    // 对联营企业和合营企业的投资收益
	"P22": "NET_GAIN_FX_TRANS",           // 汇兑净收益
	"P23": "OPPROFIT",                    // 营业利润
	"P24": "NON_OPER_REV",                // 营业外收入
	"P25": "NON_OPER_EXP",                // 营业外支出
	"P26": "NET_LOSS_DISP_NONCUR_ASSET",  // 非流动资产处置净损失
	"P27": "TOT_PROFIT",                  // 利润总额
	"P28": "TAX",                         // 所得税
	"P29": "NET_PROFIT_IS",               // 净利润
	"P30": "MINORITY_INT_INC",            // 少数股东损益
	"P31": "NP_BELONGTO_PARCOMSH",        // 归属于母公司所有者的净利润
	"P32": "OTHER_COMPREH_INC",           // 其他综合收益
	"P33": "TOT_COMPREH_INC",             // 综合收益总额
	"P34": "TOT_COMPREH_INC_MIN_SHRHLDR", // 归属于少数股东的综合收益总额
	"P35": "TOT_COMPREH_INC_PARENT_COMP", // 归属于母公司普通股东综合收益总额
	/*****************************************************************************
	 * 现金流量表--经营活动
	 *****************************************************************************/
	"CO0":  "CASH_RECP_SG_AND_RS",           // 销售商品、提供劳务收到的现金
	"CO1":  "RECP_TAX_RENDS",                // 收到的税费返还
	"CO2":  "OTHER_CASH_RECP_RAL_OPER_ACT",  // 收到其他与经营活动有关的现金
	"CO3":  "NET_INCR_INSURED_DEP",          // 保户储金净增加额
	"CO4":  "NET_INCR_DEP_COB",              // 客户存款和同业存放款项净增加额
	"CO5":  "NET_INCR_LOANS_CENTRAL_BANK",   // 向中央银行借款净增加额
	"CO6":  "NET_INCR_FUND_BORR_OFI",        // 向其他金融机构拆入资金净增加额
	"CO7":  "NET_INCR_INT_HANDLING_CHRG",    // 收取利息和手续费净增加额
	"CO8":  "CASH_RECP_PREM_ORIG_INCO",      // 收到的原保险合同保费取得的现金
	"CO9":  "NET_CASH_RECEIVED_REINSU_BUS",  // 收到的再保业务现金净额
	"CO10": "NET_INCR_DISP_TFA",             // 处置交易性金融资产净增加额
	"CO11": "NET_INCR_LOANS_OTHER_BANK",     // 拆入资金净增加额
	"CO12": "NET_INCR_REPURCH_BUS_FUND",     // 回购业务资金净增加额
	"CO13": "STOT_CASH_INFLOWS_OPER_ACT",    // 经营活动现金流入小计
	"CO14": "CASH_PAY_GOODS_PURCH_SERV_REC", // 购买商品、接受劳务支付的现金
	"CO15": "CASH_PAY_BEH_EMPL",             // 支付给职工以及为职工支付的现金
	"CO16": "PAY_ALL_TYP_TAX",               // 支付的各项税费
	"CO17": "OTHER_CASH_PAY_RAL_OPER_ACT",   // 支付其他与经营活动有关的现金
	"CO18": "NET_INCR_CLIENTS_LOAN_ADV",     // 客户贷款及垫款净增加额
	"CO19": "NET_INCR_DEP_CBOB",             // 存放央行和同业款项净增加额
	"CO20": "CASH_PAY_CLAIMS_ORIG_INCO",     // 支付原保险合同赔付款项的现金
	"CO21": "HANDLING_CHRG_PAID",            // 支付手续费的现金
	"CO22": "COMM_INSUR_PLCY_PAID",          // 支付保单红利的现金
	"CO23": "STOT_CASH_OUTFLOWS_OPER_ACT",   // 经营活动现金流出小计
	"CO24": "NET_CASH_FLOWS_OPER_ACT",       // 经营活动产生的现金流量净额
	/*****************************************************************************
	 * 现金流量表--投资活动
	 *****************************************************************************/
	"CI0":  "CASH_RECP_DISP_WITHDRWL_INVEST", // 收回投资收到的现金
	"CI1":  "CASH_RECP_RETURN_INVEST",        // 取得投资收益收到的现金
	"CI2":  "NET_CASH_RECP_DISP_FIOLTA",      // 处置固定资产、无形资产和其他长期资产收回的现金净额
	"CI3":  "NET_CASH_RECP_DISP_SOBU",        // 处置子公司及其他营业单位收到的现金净额
	"CI4":  "OTHER_CASH_RECP_RAL_INV_ACT",    // 收到其他与投资活动有关的现金
	"CI5":  "STOT_CASH_INFLOWS_INV_ACT",      // 投资活动现金流入小计
	"CI6":  "CASH_PAY_ACQ_CONST_FIOLTA",      // 购建固定资产、无形资产和其他长期资产支付的现金
	"CI7":  "CASH_PAID_INVEST",               // 投资支付的现金
	"CI8":  "NET_CASH_PAY_AQUIS_SOBU",        // 取得子公司及其他营业单位支付的现金净额
	"CI9":  "OTHER_CASH_PAY_RAL_INV_ACT",     // 支付其他与投资活动有关的现金
	"CI10": "STOT_CASH_OUTFLOWS_INV_ACT",     // 投资活动现金流出小计
	"CI11": "NET_CASH_FLOWS_INV_ACT",         // 投资活动产生的现金流量净额
	/*****************************************************************************
	 * 现金流量表--筹资活动
	 *****************************************************************************/
	"CF0":  "CASH_RECP_CAP_CONTRIB",       // 吸收投资收到的现金
	"CF1":  "CASH_REC_SAIMS",              // 子公司吸收少数股东投资收到的现金
	"CF2":  "CASH_RECP_BORROW",            // 取得借款收到的现金
	"CF3":  "OTHER_CASH_RECP_RAL_FNC_ACT", // 收到其他与筹资活动有关的现金
	"CF4":  "PROC_ISSUE_BONDS",            // 发行债券收到的现金
	"CF5":  "STOT_CASH_INFLOWS_FNC_ACT",   // 筹资活动现金流入小计
	"CF6":  "CASH_PREPAY_AMT_BORR",        // 偿还债务支付的现金
	"CF7":  "CASH_PAY_DIST_DPCP_INT_EXP",  // 分配股利、利润或偿付利息支付的现金
	"CF8":  "DVD_PROFIT_PAID_SC_MS",       // 子公司支付给少数股东的股利、利润
	"CF9":  "OTHER_CASH_PAY_RAL_FNC_ACT",  // 支付其他与筹资活动有关的现金
	"CF10": "STOT_CASH_OUTFLOWS_FNC_ACT",  // 筹资活动现金流出小计
	"CF11": "NET_CASH_FLOWS_FNC_ACT",      // 筹资活动产生的现金流量净额
	/*****************************************************************************
	 * 现金流量表--现金变动
	 *****************************************************************************/
	"CC0": "EFF_FX_FLU_CASH",           // 汇率变动对现金的影响
	"CC1": "NET_INCR_CASH_CASH_EQU_DM", // 现金及现金等价物净增加额
	"CC2": "CASH_CASH_EQU_BEG_PERIOD",  // 期初现金及现金等价物余额
	"CC3": "CASH_CASH_EQU_END_PERIOD",  // 期末现金及现金等价物余额
	/*****************************************************************************
	 * 现金流量表--补充资料
	 *****************************************************************************/
	"CS0":  "NET_PROFIT_CS",                 // 净利润
	"CS1":  "PROV_DEPR_ASSETS",              // 资产减值准备
	"CS2":  "DEPR_FA_COGA_DPBA",             // 固定资产折旧、油气资产折耗、生产性生物资产折旧
	"CS3":  "AMORT_INTANG_ASSETS",           // 无形资产摊销
	"CS4":  "AMORT_LT_DEFERRED_EXP",         // 长期待摊费用摊销
	"CS5":  "DECR_DEFERRED_EXP",             // 待摊费用减少
	"CS6":  "INCR_ACC_EXP",                  // 预提费用增加
	"CS7":  "LOSS_DISP_FIOLTA",              // 处置固定资产、无形资产和其他长期资产的损失
	"CS8":  "LOSS_SCR_FA",                   // 固定资产报废损失
	"CS9":  "LOSS_FV_CHG",                   // 公允价值变动损失
	"CS10": "FIN_EXP_CS",                    // 财务费用
	"CS11": "INVEST_LOSS",                   // 投资损失
	"CS12": "DECR_DEFERRED_INC_TAX_ASSETS",  // 递延所得税资产减少
	"CS13": "INCR_DEFERRED_INC_TAX_LIAB",    // 递延所得税负债增加
	"CS14": "DECR_INVENTORIES",              // 存货的减少
	"CS15": "DECR_OPER_PAYABLE",             // 经营性应收项目的减少
	"CS16": "INCR_OPER_PAYABLE",             // 经营性应付项目的增加
	"CS17": "UNCONFIRMED_INVEST_LOSS_CS",    // 未确认的投资损失
	"CS18": "OTHERS",                        // 其他
	"CS19": "IM_NET_CASH_FLOWS_OPER_ACT",    // 经营活动产生的现金流量净额
	"CS20": "CONV_DEBT_INTO_CAP",            // 债务转为资本
	"CS21": "CONV_CORP_BONDS_DUE_WITHIN_1Y", // 一年内到期的可转换公司债券
	"CS22": "FA_FNC_LEASES",                 // 融资租入固定资产
	"CS23": "END_BAL_CASH",                  // 现金的期末余额
	"CS24": "BEG_BAL_CASH",                  // 现金的期初余额
	"CS25": "END_BAL_CASH_EQU",              // 现金等价物的期末余额
	"CS26": "BEG_BAL_CASH_EQU",              // 现金等价物的期初余额
	"CS27": "NET_INCR_CASH_CASH_EQU_IM",     // 现金及现金等价物净增加额
}
