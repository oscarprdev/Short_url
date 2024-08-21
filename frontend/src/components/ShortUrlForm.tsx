import { ShortUrlInput } from '../services/api/shortUrl';
import { urlFormSchema } from '../services/zod/urlFormSchema';
import WrapperAction from './containers/WrapperAction';
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
				<form
					data-testid="url-form"
					onSubmit={form.handleSubmit(onSubmit)}
					className="flex items-center justify-center gap-2 w-full">
					<FormField
						control={form.control}
						name="link"
						render={({ field }) => (
							<FormItem>
								<Input
									className="text-md w-full md:w-[450px] border-stone-600"
									placeholder="Enter your link here..."
									{...field}
								/>
								{form.formState.isDirty && <FormMessage className="absolute text-red-400" />}
							</FormItem>
						)}
					/>
					<WrapperAction id="submit" color="contrast-pink">
						Short link
					</WrapperAction>
				</form>
			</Form>
		</>
	);
};

export default ShortUrlForm;
