function formatDate(timestamp) {
    let timestampLength = String(timestamp).length;
    if (timestampLength<13){
        timestamp = timestamp * 10 ** (13-timestampLength);
    }else if (timestampLength>13){
        timestamp = String(timestamp).slice(0,13)
    }
    timestamp = new Date(Number(timestamp));
    let year = timestamp.getFullYear(); //取得4位数的年份
    let month = timestamp.getMonth() + 1; //取得日期中的月份，其中0表示1月，11表示12月
    let date = timestamp.getDate(); //返回日期月份中的天数（1到31）
    let hour = timestamp.getHours(); //返回日期中的小时数（0到23）
    let minute = timestamp.getMinutes(); //返回日期中的分钟数（0到59）
    let second = timestamp.getSeconds(); //返回日期中的秒数（0到59）
    return year + "-" + (month < 10 ? "0" + month : month) + "-" + (date < 10 ? "0" + date : date) + " " + hour + ":" + minute + ":" + second;
}

// 3601s => 1h 0m 1s
function RunTimeFormater(timestamp) {
    timestamp = Number(timestamp)
    let hours = Number.parseInt(timestamp / (60 * 60))
    let remainder = Number.parseInt(timestamp % (60 * 60))
    let mins = Number.parseInt(remainder / 60)
    let seconds = Number.parseInt(remainder % 60)
    return String(hours) + "h " + String(mins) + "m " + String(seconds) + "s"
}

export {
    formatDate,
    RunTimeFormater
};