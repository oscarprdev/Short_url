export const useQueryParams = () => {
	const query = window.location.search;
	const urlSearchParams = new URLSearchParams(query);

	const queryParams = Object.fromEntries(urlSearchParams);

	return {
		queryParams,
	};
};
