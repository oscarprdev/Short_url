import { z } from 'zod';

export const homeFormSchema = z.object({
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
