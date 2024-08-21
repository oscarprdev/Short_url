import { ShortUrlInput } from '../services/api/shortUrl';
import { urlFormSchema } from '../services/zod/urlFormSchema';
import { Button } from './ui/button';
import { Form, FormField, FormItem, FormMessage } from './ui/form';
import { Input } from './ui/input';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';

const DEFAULT_VALUES = {
	link: '',
};

interface ShortUrlFormProps {
	addUrl: ({ originalUrl, userId }: ShortUrlInput) => void;
	userId?: string | undefined;
}

const ShortUrlForm = ({ addUrl, userId }: ShortUrlFormProps) => {
	const form = useForm({
		defaultValues: DEFAULT_VALUES,
		resolver: zodResolver(urlFormSchema),
	});

	const onSubmit = () => {
		const linkValue = form.getValues('link');

		return new Promise(resolve => {
			addUrl({ originalUrl: linkValue, userId });

			form.reset(DEFAULT_VALUES);
			resolve(linkValue);
		});
	};

	return (
		<>
			<Form {...form}>
				<form data-testid="url-form" onSubmit={form.handleSubmit(onSubmit)} className="flex items-center gap-2">
					<FormField
						control={form.control}
						name="link"
						render={({ field }) => (
							<FormItem>
								<Input
									className="text-md sm:w-[99vw] md:w-[400px] border-stone-600"
									placeholder="Enter your link here..."
									{...field}
								/>
								{form.formState.isDirty && <FormMessage className="absolute text-stone-400" />}
							</FormItem>
						)}
					/>
					<Button data-testid="submit-btn" type="submit" className="ml-auto font-bold">
						Short link
					</Button>
				</form>
			</Form>
		</>
	);
};

export default ShortUrlForm;
