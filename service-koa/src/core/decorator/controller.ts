import { controllerPathKeyResolve } from '../constants';
import { Metadata } from '../../common/metadata';

export const controller = (path = '/'): ClassDecorator => {
  return (target: any): void => {
    console.log(target, 1);
    Metadata.set(controllerPathKeyResolve(target), path, target);
  };
};