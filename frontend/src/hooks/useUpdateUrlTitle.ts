import { useMutation } from '@tanstack/react-query';
import { useGlobalStore } from '../store/globalState';
import { updateUrlTitle } from '../services/api/updateUrlTitle';

export const useUpdateUrlTitle = () => {
	const { user, urls, setUrls, setError } = useGlobalStore();

	return useMutation({
		mutationFn: async ({ urlId, urlTitle }: { urlId: string; urlTitle: string }) => {
			return await updateUrlTitle({ urlId, userId: user?.id, urlTitle });
		},
		onSuccess: ({ url }) => {
			if (urls) {
				const updatedUrls = urls.map((currentUrl) => (currentUrl.id === url.id ? url : currentUrl));
				setUrls(updatedUrls);
			} else {
				setError('No urls availables');
			}
		},
		onError: (error) => {
			setError(error.message);
		},
	});
};
