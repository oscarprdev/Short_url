import { useEffect } from 'react';
import { useToast } from '../components/ui/use-toast';

export const useErrorToast = (error: string | null) => {
	const { toast } = useToast();

	useEffect(() => {
		if (error) {
			toast({
				title: 'Error',
				description: error,
			});
		}
	}, [error, toast]);
};
