import { ReactNode, createContext } from 'react';

interface ModalContext {
	isOpen: boolean;
	openModal: (modal: ReactNode) => void;
	closeModal: () => void;
}

const modalContextState: ModalContext = {
	isOpen: false,
	openModal() {},
	closeModal() {},
};

export const ModalContext = createContext<ModalContext>(modalContextState);
