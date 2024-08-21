import { Button } from '../ui/button';
import { cn } from '../ui/lib/utils';
import { ReactNode } from 'react';

type WrapperActionProps = {
	type?: 'url' | 'button' | 'card';
	size?: 'md' | 'sm';
	url?: string;
	id: string;
	onClick?: () => void;
	classname?: string;
	color: 'contrast' | 'contrast-pink' | 'contrast-blue' | 'default';
	children: ReactNode;
};

const WrapperAction = ({
	id,
	type = 'button',
	size = 'sm',
	url,
	onClick,
	classname,
	color,
	children,
}: WrapperActionProps) => {
	return (
		<>
			{type === 'url' && url ? (
				<a
					href={url}
					data-testid={`url-link-${id}`}
					target="blank"
					className={cn(classname, 'group relative mt-5 self-center')}>
					<div
						className={cn(
							size !== 'sm' ? 'px-14 py-5' : 'px-7 py-5',
							'flex items-center rounded-md border text-zinc-300 hover:text-zinc-100 border-zinc-500 bg-black duration-200 ease-out group-hover:translate-x-1 group-hover:translate-y-1 group-hover:rounded-lg'
						)}>
						{children}
					</div>
					<div
						aria-hidden
						className={cn(
							color !== 'default' ? `bg-[var(--${color})]` : 'bg-white/50',
							'absolute left-1 top-1 z-[-1] h-full w-full rounded-lg'
						)}></div>
				</a>
			) : type === 'button' ? (
				<Button
					data-testid={`${id}-btn`}
					type="submit"
					onClick={onClick}
					variant={'action'}
					className="group relative w-24 h-10 ml-2">
					<div className="absolute -top-1 -left-1 z-1 group-hover:bg-black group-hover:text-white rounded-lg ml-auto font-bold bg-[var(--background)] text-zinc-100 border border-zinc-500 w-full h-full grid place-items-center duration-200 ease-out group-hover:translate-x-1 group-hover:translate-y-1">
						{children}
					</div>
					<div
						aria-hidden
						className={cn(
							color !== 'default' ? `bg-[var(--${color})]` : 'bg-white/50',
							'absolute left-0 top-0 z-[-1] h-full w-full rounded-lg'
						)}></div>
				</Button>
			) : (
				<article data-testid={`${id}`} className="group relative w-[200px] h-[110px] p-2 pl-0">
					<div className="absolute flex flex-col gap-1 p-5 w-[200px] -top-1 -left-1 z-1 bg-black group-hover:text-white rounded-lg ml-auto font-bold bg-[var(--background)] text-zinc-100 border border-zinc-500 h-full duration-200 ease-out group-hover:translate-x-1 group-hover:translate-y-1">
						{children}
					</div>
					<div
						aria-hidden
						className={cn(
							color !== 'default' ? `bg-[var(--${color})]` : 'bg-white/50',
							'absolute left-0 top-0 z-[-1] h-full w-full rounded-lg'
						)}></div>
				</article>
			)}
		</>
	);
};

export default WrapperAction;
