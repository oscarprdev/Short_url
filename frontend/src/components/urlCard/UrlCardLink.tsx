import { IconLink } from '@tabler/icons-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '../ui/tooltip';

interface UrlCardLinkProps {
	originalUrl: string;
	shortUrl: string;
	onUrlClick(): void;
}
const UrlCardLink = ({ originalUrl, shortUrl, onUrlClick }: UrlCardLinkProps) => {
	return (
		<TooltipProvider>
			<Tooltip>
				<TooltipTrigger>
					<div className='flex items-center gap-2'>
						<IconLink className='w-4 text-stone-400' />
						<a
							href={originalUrl}
							onClick={onUrlClick}
							target='blank'
							className='hover:underline truncate'>
							{shortUrl}
						</a>
					</div>
				</TooltipTrigger>
				<TooltipContent>Shorten link of: {originalUrl}</TooltipContent>
			</Tooltip>
		</TooltipProvider>
	);
};

export default UrlCardLink;
