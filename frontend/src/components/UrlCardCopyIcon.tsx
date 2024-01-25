import { IconCopy, IconCheck } from '@tabler/icons-react';

import { useState } from 'react';

interface UrlCardIconCopyProps {
	url: string;
}

const UrlCardCopyIcon = ({ url }: UrlCardIconCopyProps) => {
	const [isCopied, setIsCopied] = useState(false);

	const handleCopyClick = async () => {
		const shortUrl = url;
		try {
			await navigator.clipboard.writeText(shortUrl);
			setIsCopied(true);

			setTimeout(() => setIsCopied(false), 5000);
		} catch (err) {
			console.error('Unable to copy to clipboard:', err);
		}
	};

	return (
		<>
			{isCopied ? (
				<IconCheck className='animate-fade-in w-4 text-stone-300 absolute top-2 right-2' />
			) : (
				<IconCopy
					onClick={handleCopyClick}
					className='animate-fade-in w-4 text-stone-300 absolute top-2 right-2'
				/>
			)}
		</>
	);
};

export default UrlCardCopyIcon;
