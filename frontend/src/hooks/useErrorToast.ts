import { useEffect } from 'react';
import { useToast } from '../components/ui/use-toast';
import { useGlobalStore } from '../store/globalState';

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
	}, [error, toast]);
};
