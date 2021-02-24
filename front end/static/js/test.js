let xhr = new XMLHttpRequest()
xhr.open('POST', 'http://118.178.190.150:80/bilibili/register');
xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
xhr.send(params);
xhr.onload = () => {
    const res = xhr.responseText;
    const {
        res
    } = message;
}