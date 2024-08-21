import { deleteUrl } from '../services/api/deleteUrl';
import { useGlobalStore } from '../store/globalState';
import { useMutation } from '@tanstack/react-query';

export const useDeleteUrl = () => {
	const { user, urls, setUrls, setError } = useGlobalStore();

	return useMutation({
		mutationFn: async ({ urlId }: { urlId: string }) => {
			await deleteUrl({ urlId, userId: user?.id });

			return urlId;
		},
		onSuccess: (id: string) => {
			if (Array.isArray(urls)) {
				const updatedUrls = urls.filter(currentUrl => currentUrl.id !== id);
				setUrls(updatedUrls);
			}
		},
		onError: error => {
			setError(error.message);
		},
	});
};
