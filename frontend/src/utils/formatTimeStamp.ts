export const formatTimeStamp = (timestamp: string) => {
	const date = new Date(timestamp);
	const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric' };

	return date.toLocaleString('en-US', options);
};
