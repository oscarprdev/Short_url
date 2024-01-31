import { IconCopy, IconCheck } from '@tabler/icons-react';

import { useState } from 'react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '../ui/tooltip';

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
				<div className='absolute top-2 right-2'>
					<TooltipProvider>
						<Tooltip>
							<TooltipTrigger>
								<IconCheck className='animate-fade-in w-4 text-stone-300 ' />
							</TooltipTrigger>
							<TooltipContent>Copied!</TooltipContent>
						</Tooltip>
					</TooltipProvider>
				</div>
			) : (
				<span className='absolute top-2 right-2'>
					<TooltipProvider>
						<Tooltip>
							<TooltipTrigger>
								<IconCopy
									data-testid='icon-copy'
									onClick={handleCopyClick}
									className='animate-fade-in w-4 text-stone-300 '
								/>
							</TooltipTrigger>
							<TooltipContent>Copy link</TooltipContent>
						</Tooltip>
					</TooltipProvider>
				</span>
			)}
		</>
	);
};

export default UrlCardCopyIcon;
