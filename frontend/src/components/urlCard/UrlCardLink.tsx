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
				<TooltipTrigger asChild>
					<div aria-label="url-link" className="flex items-center gap-2 w-[170px]">
						<a
							href={originalUrl}
							onClick={onUrlClick}
							target="blank"
							className="hover:underline truncate text-xs">
							{shortUrl}
						</a>
					</div>
				</TooltipTrigger>
				<TooltipContent className="z-50">{originalUrl}</TooltipContent>
			</Tooltip>
		</TooltipProvider>
	);
};

export default UrlCardLink;
