import { useToast } from '../components/ui/use-toast';
import { useGlobalStore } from '../store/globalState';
import { useEffect } from 'react';

export const useErrorToast = (error: string | null) => {
	const { setError } = useGlobalStore();
	const { toast } = useToast();

	useEffect(() => {
		if (error) {
			toast({
				title: 'Error',
				description: error,
			});
			setError(null);
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [error, toast]);
};
