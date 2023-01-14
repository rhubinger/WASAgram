function now() {
    var pad = (x, len) => String(x).padStart(len, '0')
    var currentdate = new Date(); 
    var year = pad(currentdate.getFullYear(), 4);
    var month = pad((currentdate.getMonth()+1), 2);
    var date = pad(currentdate.getDate(), 2);
    var hour = pad(currentdate.getHours(), 2);
    var minute = pad(currentdate.getMinutes(), 2);
    var second = pad(currentdate.getSeconds(), 2);
    return `${year}-${month}-${date} ${hour}:${minute}:${second}`;
}

export default {
    now,
}