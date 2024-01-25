import { Button } from './ui/button';
import { Form, FormField, FormItem, FormMessage } from './ui/form';
import { Input } from './ui/input';
import { useForm } from 'react-hook-form';
import { IconScissors } from '@tabler/icons-react';
import { zodResolver } from '@hookform/resolvers/zod';
import { homeFormSchema } from '../services/zod/homeFormSchema';

const DEFAULT_VALUES = {
	link: '',
};

interface HomeFormProps {
	addUrl: (url: string) => void;
}

const ShortUrlForm = ({ addUrl }: HomeFormProps) => {
	const form = useForm({
		defaultValues: DEFAULT_VALUES,
		resolver: zodResolver(homeFormSchema),
	});

	const onSubmit = () => {
		const linkValue = form.getValues('link');

		return new Promise((resolve) => {
			addUrl(linkValue);

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
