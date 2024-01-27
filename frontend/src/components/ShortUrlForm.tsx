import { Button } from './ui/button';
import { Form, FormField, FormItem, FormMessage } from './ui/form';
import { Input } from './ui/input';
import { useForm } from 'react-hook-form';
import { IconScissors } from '@tabler/icons-react';
import { zodResolver } from '@hookform/resolvers/zod';
import { urlFormSchema } from '../services/zod/urlFormSchema';
import { ShortUrlInput } from '../services/api/shortUrl';

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

		return new Promise((resolve) => {
			addUrl({ originalUrl: linkValue, userId });

			form.reset(DEFAULT_VALUES);
			resolve(linkValue);
		});
	};

	return (
		<>
			<Form {...form}>
				<form
					onSubmit={form.handleSubmit(onSubmit)}
					className='flex items-center gap-2'>
					<FormField
						control={form.control}
						name='link'
						render={({ field }) => (
							<FormItem>
								<Input
									className='text-md w-[400px] border-stone-600'
									placeholder='Enter your link here...'
									{...field}
								/>
								{form.formState.isDirty && <FormMessage className='absolute text-stone-400' />}
							</FormItem>
						)}
					/>
					<Button
						type='submit'
						className='ml-auto'>
						<IconScissors className={`${form.formState.isSubmitting && 'animate-spin duration-500'}`} />
					</Button>
				</form>
			</Form>
		</>
	);
};

export default ShortUrlForm;
