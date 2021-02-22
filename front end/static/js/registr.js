//获取html中的元素节点
let fm_username = document.getElementById("fm-username");
let fm_password = document.getElementById("fm-password");
let return_message = document.getElementById("return-message");
let judge_password = document.getElementsByClassName("judge-password")[0];
let checkbox = document.getElementsByClassName("fm-checkbox")[0];
let frontBox = document.getElementsByClassName("front-box")[0];
let register_submit = document.getElementById("register-submit");
let phoneNumber = document.getElementsByClassName("register-phone")[0];
let judge_phone = document.getElementsByClassName("judge-phone")[0];
//复选盒子选中状态
let i = 0;
frontBox.addEventListener('click', function () {
    if (i % 2 == 0) {
        checkbox.style.backgroundImage = "url(../../../static/image/checkedbox.jpg)";
        register_submit.style.cursor = 'pointer';
        register_submit.style.color = '#fff';
        register_submit.style.backgroundColor = '#33b4de';
        register_submit.style.borderColor = '#33b4de';
        i = 1;
    } else {
        console.log(i);
        checkbox.style.backgroundImage = "url(../../../static/image/uncheckedbox.png)";
        register_submit.style.cursor = 'not-allowed';
        register_submit.style.backgroundColor = '#f5f5f5';
        register_submit.style.borderColor = '#d9d9d9';
        register_submit.style.color = 'rgba(0,0,0,.25)';
        i = 0;
    }
}, true)
//检查用户写的昵称是否符合规范
fm_username.oninput = function () {
    let username = this.value;
    let regUser = /^[a-zA-Z0-9\u4e00-\u9fa5-_]{2,10}$/;
    if (username.length == 0) {
        return_message.innerHTML = "请告诉我你的昵称吧"
    } else if (username.length < 2) {
        return_message.innerHTML = "用户昵称过短";
    } else if (username.length > 10) {
        return_message.innerHTML = "用户昵称过长";
    } else if (!regUser.test(username)) {
        return_message.innerHTML = "昵称不可包含除-和_以外的特殊字符";
    } else return_message.innerHTML = "";
}
//检查用户写的密码是否符合规范
fm_password.onblur = function () {
    let password = this.value;
    let regPass = /^[\n]{6,16}$/;
    if (password.length < 6) {
        judge_password.innerHTML = "密码不能小于6个字符";
    } else if (password.length > 16) {
        judge_password.innerHTML = "密码不能大于16个字符";
    } else judge_password.innerHTML = "";
}
//检查用户写的手机号是否符合规范
phoneNumber.oninput = function () {
    let phone = this.value;
    let phoneReg = /^1[0-9]{10}$/;
    if (!phoneReg.test(phone)) {
        judge_phone.innerHTML = "请输入正确的手机号码";
    } else {
        judge_phone.innerHTML = "";
    }
}
//检验用户名是否重复
//url处填写后端网址
//点击确定发送数据
register_submit.addEventListener('click', function () {
    let username = fm_username.value
    let password = fm_password.value
    let phone = phoneNumber.value
    if (username && password && phone && i == 1) {
        fetch("http: //118.178.190.150:80/bilibili/register", {
                method: 'POST',
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                body: JSON.stringify({
                    "username": username,
                    "password": password,
                    "phone": phone,
                })
            })
            .then(res => res.json())
            .then(res => {
                console.log(res)
            })
    }
})