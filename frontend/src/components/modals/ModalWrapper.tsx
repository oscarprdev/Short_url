import { useModal } from '../../hooks/useModal';
import { IconX } from '@tabler/icons-react';
import { ReactNode } from 'react';

interface ModalProps {
	children: ReactNode;
}

const ModalWrapper = ({ children }: ModalProps) => {
	const { closeModal } = useModal();

	return (
		<div className="z-20 animate-fade-up absolute grid place-items-center top-0 w-screen h-screen bg-[var(--backdrop)] overflow-hidden">
			<div className="relative p-5 pt-10 w-fit h-fit min-w-[300px] z-[-2] bg-black border border-stone-800 shadow-lg rounded-2xl">
				<div
					onClick={closeModal}
					className="absolute grid place-items-center w-10 h-10 text-white rounded-full p-2 top-2 left-2 cursor-pointer hover:bg-zinc-900 duration-300">
					<IconX />
				</div>
				{children}
			</div>
		</div>
	);
};

export default ModalWrapper;
