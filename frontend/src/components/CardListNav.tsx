import { IconLayoutGrid, IconLayoutList } from '@tabler/icons-react';
import { useGlobalStore } from '../store/globalState';

interface CardListNavProps {
	isRowsLayout: boolean;
	onSetRowLayout(): void;
	onSetGridLayout(): void;
}

const CardListNav = ({ isRowsLayout, onSetRowLayout, onSetGridLayout }: CardListNavProps) => {
	const { urls } = useGlobalStore();
	return (
		<nav className='relative w-[90vw] md:w-[1100px] flex justify-center pb-2'>
			<p
				aria-label='link-counter'
				className='absolute left-2 md:left-20 text-sm text-stone-400'>
				Total {urls?.length}
			</p>
			<div className='flex items-center gap-2'>
				<IconLayoutList
					onClick={onSetRowLayout}
					className={`${isRowsLayout ? 'text-stone-200' : 'text-stone-600 hover:text-stone-200'} cursor-pointer duration-300`}
				/>
				<IconLayoutGrid
					onClick={onSetGridLayout}
					className={`${!isRowsLayout ? 'text-stone-200' : 'text-stone-600 hover:text-stone-200'} cursor-pointer duration-300`}
				/>
			</div>
		</nav>
	);
};

export default CardListNav;
