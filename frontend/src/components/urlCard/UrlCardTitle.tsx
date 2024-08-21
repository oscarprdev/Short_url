import { useUpdateUrlTitle } from '../../hooks/useUpdateUrlTitle';
import { formatCapitalize } from '../../utils/formatCapitalized';
import { IconCheck, IconLoader2, IconX } from '@tabler/icons-react';
import { ChangeEvent, FocusEvent, KeyboardEvent, useRef, useState } from 'react';

interface UrlCardTitleProps {
	title: string;
	urlId: string;
}

const UrlCardTitle = ({ title, urlId }: UrlCardTitleProps) => {
	const [titleInput, setTitleInput] = useState(title);
	const titleRef = useRef<HTMLInputElement>(null);
	const { mutate: updateTitle, isPending, isSuccess, isError } = useUpdateUrlTitle();

	const onInputTitleChange = (e: ChangeEvent) => {
		if (e.target instanceof HTMLInputElement && titleRef.current) {
			setTitleInput(e.target.value.substring(0, 20));
		}
	};

	const onInputKeyDown = (e: KeyboardEvent<HTMLInputElement>): void => {
		if (e.key === 'Enter') {
			titleRef.current?.blur();
		}
	};

	const onInputBlur = (e: FocusEvent<HTMLInputElement, Element>): void => {
		e.preventDefault();

		updateTitle({ urlId, urlTitle: formatCapitalize(titleInput) });
	};

	return (
		<div className="flex items-center justify-start w-full gap-2">
			<input
				maxLength={40}
				ref={titleRef}
				value={titleInput}
				onChange={onInputTitleChange}
				onKeyDown={onInputKeyDown}
				onBlur={onInputBlur}
				className="bg-transparent outline-none w-fit max-w-[80%] border-b border-stone-600"
			/>
			{isPending ? (
				<IconLoader2 className="w-4 animate-spin text-stone-300" />
			) : isSuccess ? (
				<IconCheck className="animate-fade-in w-4 text-stone-300" />
			) : (
				isError && <IconX className="animate-fade-in w-4 text-stone-300" />
			)}
		</div>
	);
};

export default UrlCardTitle;
