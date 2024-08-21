import ButtonLogin from './ButtonLogin';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, it } from 'vitest';

describe('ButtonLogin', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<ButtonLogin />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByTestId('login-btn');
		component.getByText('Login');
	});
});
