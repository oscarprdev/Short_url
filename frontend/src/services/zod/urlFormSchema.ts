import { z } from 'zod';

export const urlFormSchema = z.object({
	link: z.string().refine(
		(url) => {
			try {
				new URL(url);
				return true;
			} catch (error) {
				return false;
			}
		},
		{
			message: 'Invalid URL format',
		}
	),
});
