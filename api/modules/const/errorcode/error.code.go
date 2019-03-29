package errorcode

const HTTPErrorActivityProductStatusError = 601 //活动产品状态不正确，请刷新后重试
const HTTPErrorActivityProductCountError = 602  //活动产品数量不足，请刷新后重试
const HTTPErrorActivityTypeError = 610          //活动类型不正确

const HTTPErrorActivityStatusError = 603       //该订单所在活动已过期，请退出
const HTTPErrorProductCustomerQuota = 604      //用户购买已达上限
const HTTPErrorActivityReceiveTypeError = 605  //用户领取类型错误
const HTTPErrorLimitAccountQuotaError = 606    //用户资格已经兑换
const HTTPErrorActivityAlrealdyJoinError = 607 //用户已参与了该活动
const HTTPErrorFailedValidateRuleError = 608   //用户为满足验证规则

const HTTPErrorFailedPrecheckError = 609         //用户为满足验证规则
const HTTPErrorFailedAuthParamsVerifyError = 611 //用户授权参数错误

const HTTPErrorFailedSMSCheckedError = 612     //短信验证码失败
const HTTPErrorFailedCodeCheckedError = 613    //短信验证码失败
const HTTPErrorFailedCodeErrorCountError = 614 //短信验证码失败

const HTTPErrorPhoneExistError = 615  //该限制手机号已存在
const HTTPErrorParamsError = 616      //请求参数有误
const HTTPErrorInviteExistError = 616 //该限制邀请码已存在

const HTTPErrorRecvCountOverLimitError = 617        //领取次数超过限制
const HTTPErrorNoneDataRequiredError = 618          //无数据需要处理
const HTTPErrorCodeError = 619                      //该错误码已经存在
const HTTPErrorParamsNoneProductError = 620         //请求参数有误
const HTTPErrorParamsUniquePhoneNotMatchError = 621 //唯一码和手机号不匹配
const HTTPErrorUniqueCodeError = 623                //活动链接错误

const HTTPErrorRequestLimitTooFastError = 701 //请求太频繁了

const HTTPErrorSalemanPhoneExistError = 622 //当前业务员处于启用状态

const HTTPErrorSceneTypeExistError = 624  //当前场景类型已经存在，请修正
const HTTPErrorActivityIDExistError = 625 //当前活动已经添加过微信回复了，请修正

const HTTPErrorWechatCodeError = 702 //微信验证码错误

const HTTPErrorUserAccountError = 406  //用户账号不存在
const HTTPErrorUserPasswordError = 900 //用户密码错误
const HTTPErrorBranchNameError = 901   //门店名称已经存在

const HTTPErrorSysInfoError = 703 //系统信息不存在

const HTTPErrorBankCardPWDError = 704 //两次密码输入不一致
const HTTPErrorOldPWDError = 705      //原密码错误
