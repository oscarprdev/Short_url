import { afterEach, beforeEach, describe, it } from 'vitest';
import { RenderResult, render } from '@testing-library/react';
import ButtonLogin from './ButtonLogin';

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
