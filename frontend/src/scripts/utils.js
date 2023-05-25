import {ref, reactive, onMounted, onUnmounted} from 'vue';

export function clearObject(obj) {
	if (Array.isArray(obj))
		for (let i = 0; i < obj.length; i++) {
			obj.pop();
		}
	else
		for (var [key, value] in Object.entries(obj)) {
			delete obj[key];
		}
	return obj;
}

export function setObjectValue({obj, key, resObj = {}, resKey = null, includeEmpty = true, includeUndefined = false, type = null, func = null}) {
	// func=(x)=>x}

	let k = resKey ? resKey : key;
	if (obj && obj[key] && obj[key] != '') resObj[k] = obj[key];
	else if (obj && obj[key] && includeEmpty) resObj[k] = obj[key];
	else if (includeUndefined) resObj[k] = '';
	if (resObj[k] && func) resObj[k] = func(resObj[k]);
	return resObj;
}

export function gameidToDateTimeStr(str = '20221210-114856') {
	//2023-03-25 19:42:21
	let arr = str.split('');
	arr.splice(4, 0, '-');
	arr.splice(7, 0, '-');
	arr.splice(10, 1, ' ');
	arr.splice(13, 0, ':');
	arr.splice(16, 0, ':');
	str = arr.join('');
	return str;
}

export function normalizeLength(text, paddingLength = 2, padding = '.') {
	if (text.toString().length == Math.abs(paddingLength)) return text;
	else if (paddingLength < 0) return this.normalizeLength((padding + text).substr(paddingLength), paddingLength, padding);
	else if (paddingLength > 0) return this.normalizeLength((text + padding).substr(0, paddingLength), paddingLength, padding);
}

let twoNumbers = (number) => ('0' + number).substr(-2);
export function currentDateTime(date_time_datetime = 'datetime', separators = ['-', ' ', ':']) {
	let dateTime = new Date(Date.now());
	let [year, month, day] = [dateTime.getFullYear(), twoNumbers(dateTime.getMonth() + 1), twoNumbers(dateTime.getDate())];
	let [hours, minutes, seconds, ms] = [twoNumbers(dateTime.getHours()), twoNumbers(dateTime.getMinutes()), twoNumbers(dateTime.getSeconds()), dateTime.getMilliseconds()];

	let date_str = year + separators[0] + month + separators[0] + day;
	let time_str = hours + separators[2] + minutes + separators[2] + seconds;
	let dateTime_str = date_str + separators[1] + time_str;

	if (date_time_datetime == 'date') return date_str;
	else if (date_time_datetime == 'time') return time_str;
	else if (date_time_datetime == 'yyyyMM') return year + separators[0] + month;
	else return dateTime_str;
}
export function getObjectValue(obj = {}, path = []) {
	if (!obj) return;
	let k = path.shift();
	if (!path || path.length > 0) return getObjectValue(obj[k], path);
	return obj[k];
}
export function capitalize(key) {
	return key.charAt(0).toUpperCase() + key.slice(1);
}
export function type(val, minYear = 1980) {
	if (Array.isArray(val)) return 'array';
	if (typeof val == 'object') return 'object';
	if (!isNaN(Date.parse(val)) && new Date(val).getFullYear() > minYear) return 'date';
	if (typeof val == 'boolean') return 'bool';
	if (typeof val == 'string' && (val.toLowerCase() == 'true' || val.toLowerCase() == 'false')) return 'bool';
	if (isNaN(val)) return 'string';
	val = Number(val);
	if (isFloat(val)) return 'float'; //parseFloat(val)
	if (isInt(val)) return 'int'; //parseInt(val)
}
export function isInt(n) {
	return Number(n) === n && n % 1 === 0;
}
export function isFloat(n) {
	return Number(n) === n && n % 1 !== 0;
}
export function sortObject(obj) {
	return Object.entries(obj)
		.sort((a, b) => b[1] - a[1])
		.reduce((acc, c) => ({...acc, [c[0]]: c[1]}), {});
}
