export const resize = (chart) => {
	let timer = 0;
	window.addEventListener("resize", () => {
		clearTimeout(timer);

		timer = setTimeout(() => {
			chart.resize();
		}, 300);
	});
};


export const FormatDate = (timestamp) => {
	// 补全为13位
	timestamp = new Date(timestamp);
	let minute = 1000 * 60;
	let hour = minute * 60;
	let day = hour * 24;
	let month = day * 30;
	let now = new Date().getTime();
	let diffValue = now - timestamp;
	// 如果本地时间反而小于变量时间
	if (diffValue < 0) {
		return '不久前';
	}
	// 计算差异时间的量级
	let monthC = diffValue / month;
	let weekC = diffValue / (7 * day);
	let dayC = diffValue / day;
	let hourC = diffValue / hour;
	let minC = diffValue / minute;
	// 数值补0方法
	let zero = function (value) {
		if (value < 10) {
			return '0' + value;
		}
		return value;
	};
	// 使用
	if (monthC > 4) {
		// 超过1年，直接显示年月日
		return (function () {
			let date = new Date(timestamp);
			return date.getFullYear() + '年' + zero(date.getMonth() + 1) + '月' + zero(date.getDate()) + '日';
		})();
	} else if (monthC >= 1) {
		return parseInt(monthC) + '月前';
	} else if (weekC >= 1) {
		return parseInt(weekC) + '周前';
	} else if (dayC >= 1) {
		return parseInt(dayC) + '天前';
	} else if (hourC >= 1) {
		return parseInt(hourC) + '小时前';
	} else if (minC >= 1) {
		return parseInt(minC) + '分钟前';
	}
	return '刚刚';
}

/**
* 模块名: GetNowTime
* 代码描述:
* 作者:lizibin
* 创建时间:2023/01/02 18:11:47
*/
export const GetNowTime = () => {
	let now = new Date();
	let year = now.getFullYear(); //获取完整的年份(4位,1970-????)
	let month = now.getMonth() + 1; //获取当前月份(0-11,0代表1月)
	let today = now.getDate(); //获取当前日(1-31)
	let hour = now.getHours(); //获取当前小时数(0-23)
	let minute = now.getMinutes(); //获取当前分钟数(0-59)
	let second = now.getSeconds(); //获取当前秒数(0-59)
	let nowTime = ''
	nowTime = year + '-' + fillZero(month) + '-' + fillZero(today) + ' ' + fillZero(hour) + ':' +
		fillZero(minute) + ':' + fillZero(second)
	return nowTime
};

function fillZero(str) {
	var realNum;
	if (str < 10) {
		realNum = '0' + str;
	} else {
		realNum = str;
	}
	return realNum;
}