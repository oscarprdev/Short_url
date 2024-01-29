import { ReactNode } from 'react';

interface AuthInputProps {
	children: ReactNode;
	contrast?: boolean;
	href: string;
}

const AuthInput = ({ children, contrast, href }: AuthInputProps) => {
	return (
		<a
			href={href}
			className={`${
				contrast ? 'bg-[var(--contrast)] hover:bg-[var(--contrast-hover)]' : 'bg-primary hover:bg-primary/90'
			} cursor-pointer flex-items gap-2 text-primary-foreground font-geistLight border-2 shadow h-9 px-4 py-2 inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50`}>
			{children}
		</a>
	);
};

export default AuthInput;
