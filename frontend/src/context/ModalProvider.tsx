import ModalWrapper from '../components/modals/ModalWrapper';
import { ModalContext } from './ModalContext';
import { ReactNode, useRef, useState } from 'react';

interface ModalContextProviderProps {
	children: ReactNode;
}

export const ModalProvider = ({ children }: ModalContextProviderProps) => {
	const [isOpen, setIsOpen] = useState(false);
	const modalRef = useRef<ReactNode>();

	const openModal = (modal: ReactNode) => {
		modalRef.current = modal;
		setIsOpen(true);
	};
	const closeModal = () => {
		modalRef.current = null;
		setIsOpen(false);
	};

	return (
		<ModalContext.Provider value={{ isOpen, openModal, closeModal }}>
			{children}
			{isOpen && <ModalWrapper>{modalRef.current}</ModalWrapper>}
		</ModalContext.Provider>
	);
};
