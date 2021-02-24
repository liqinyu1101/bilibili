fetch("http://118.178.190.150:80/cookie")
    .then(res => {
        a = res.clone().arrayBuffer()
        return res.clone().json()
    })
    .then(res =>
        res.cookieValue
    )
    .then(cookieValue => {
        console.log(cookieValue)
        if (cookieValue) {
            userCen[0].style.display = "none"
            userCen[1].style.display = "flex"
        }
    })
    .catch(
        e => {
            return a
        }
    )
//1.禁止浏览器打开文件行为    
document.addEventListener("drop", function (e) { //拖离     
    e.preventDefault();
})
document.addEventListener("dragleave", function (e) { //拖后放     
    e.preventDefault();
})
document.addEventListener("dragenter", function (e) { //拖进    
    e.preventDefault();
})
document.addEventListener("dragover", function (e) { //拖来拖去      
    e.preventDefault();
})
//1.拖拽上传事件获取视频文件
let box = document.querySelector(".homeDrop")
let video1 = document.querySelector(".upload-step1-container")
let video2 = document.querySelector(".upload-step2-container")
let id_title = document.querySelector("#modified")
let upload_btn = document.querySelector(".upload-btn-title")
let btn_file = document.getElementById("btn_file")
let loading = document.querySelector(".file-upload-progress-loading")
let percentage = document.querySelector(".upload-percentage")
let fileList = null;
box.addEventListener("drop", function (e) {
    fileList = e.dataTransfer.files[0];
    let {
        lastModified
    } = fileList;
    if (fileList) {
        video1.style.display = "none";
        video2.style.display = "block";
        id_title.innerText = lastModified;
    }
})
//2.点击上传事件获取视频文件对象
upload_btn.addEventListener('click', function (e) {
    btn_file.click();
})
btn_file.addEventListener('change', function () {
    fileList = btn_file.files[0];
    let {
        lastModified
    } = fileList
    if (fileList) {
        video1.style.display = "none";
        video2.style.display = "block";
        id_title.innerText = lastModified;
    }
})
//3.点击实现图片上传
let cover_preview = document.querySelector(".cover-preview")
let cover_upload = document.querySelector(".cover-upload")
let cover = null;
cover_preview.addEventListener('click', function (e) {
    cover_upload.click();
})
cover_upload.addEventListener('change', function () {
    cover = cover_upload.files[0]
    let fileurl = window.URL.createObjectURL(cover);
    if (cover.type.indexOf('image') === 0) { //如果是图片  
        let str = "<img class='preview-pic' width='100%' height='100%' src='" + fileurl + "'>" +
            "<div class='cover-upload-tip'><span>封面预览</span></div>"
        document.querySelector('.cover-preview-default').innerHTML = str;
    } else {
        alert('图片格式错误')
    }
})
//4.拖拽实现图片上传
cover_preview.addEventListener("drop", function (e) {
    cover = e.dataTransfer.files[0];
    let fileurl = window.URL.createObjectURL(cover);
    if (cover.type.indexOf('image') === 0) { //如果是图片  
        let str = "<img class='preview-pic' width='100%' height='100%' src='" + fileurl + "'>" +
            "<div class='cover-upload-tip'><span>封面预览</span></div>"
        document.querySelector('.cover-preview-default').innerHTML = str;
    } else {
        alert('图片格式错误')
    }
})
//5.字数统计
let video_title = document.querySelector(".video-title")
let video_commit = document.querySelector(".video-commit")
let tip1 = document.querySelector(".input-title-tip")
let tip2 = document.querySelector(".input-commit-tip")
video_title.addEventListener('input', function () {
    tip1.innerText = video_title.value.length + "/80"
})
video_commit.addEventListener('input', function () {
    tip2.innerText = video_commit.value.length + "/250"
})
//提交表单
let warning = document.querySelector(".warning")
let btn = document.querySelector('.submit-btn')
btn.addEventListener('click', function () {
    console.log(btn)
    console.log(warning)
    let title = video_title.value;
    let commit = video_commit.value;
    if (cover == null || title == 0 || commit == 0) {
        console.log(cover)
        warning.innerText = "请完成视频相关信息的填写"
    } else {
        //发送封面相关信息
        let video_date = new FormData();
        video_date.append("video", fileList)
        video_date.append("cover", cover)
        video_date.append("title", cover)
        video_date.append("commit", commit)
        fetch("http://118.178.190.150:80/bilibili/video/upload", {
                method: 'POST',
                body: video_date
            })
            .then(res => {
                a = res.clone().arrayBuffer()
                return res.clone().json()
            })
            .then(res => {
                console.log(res)
                if (res.message == "successfully") {
                    alert("投稿成功")
                    window.location.href = "user.html"
                }
            })
            .catch(
                e => {
                    return a
                }
            )
    }
})